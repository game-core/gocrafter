//go:generate mockgen -source=./loginReward_service.go -destination=./loginReward_service_mock.gen.go -package=loginReward
package loginReward

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"

	itemRequest "github.com/game-core/gocrafter/api/presentation/request/item"
	request "github.com/game-core/gocrafter/api/presentation/request/loginReward"
	response "github.com/game-core/gocrafter/api/presentation/response/loginReward"
	"github.com/game-core/gocrafter/config/times"
	masterEventEntity "github.com/game-core/gocrafter/domain/entity/master/event"
	masterItemEntity "github.com/game-core/gocrafter/domain/entity/master/item"
	masterLoginRewardEntity "github.com/game-core/gocrafter/domain/entity/master/loginReward"
	userLoginRewardEntity "github.com/game-core/gocrafter/domain/entity/user/loginReward"
	masterLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/master/loginReward"
	userRepository "github.com/game-core/gocrafter/domain/repository/user"
	userLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/user/loginReward"
	eventService "github.com/game-core/gocrafter/domain/service/event"
	itemService "github.com/game-core/gocrafter/domain/service/item"
)

type LoginRewardService interface {
	GetLoginRewardModel(req *request.GetLoginRewardModel, now time.Time) (*response.GetLoginRewardModel, error)
	ReceiveLoginReward(req *request.ReceiveLoginReward, now time.Time) (*response.ReceiveLoginReward, error)
}

type loginRewardService struct {
	transactionRepository       userRepository.TransactionRepository
	loginRewardStatusRepository userLoginRewardRepository.LoginRewardStatusRepository
	loginRewardModelRepository  masterLoginRewardRepository.LoginRewardModelRepository
	loginRewardRewardRepository masterLoginRewardRepository.LoginRewardRewardRepository
	eventService                eventService.EventService
	itemService                 itemService.ItemService
}

func NewLoginRewardService(
	transactionRepository userRepository.TransactionRepository,
	loginRewardStatusRepository userLoginRewardRepository.LoginRewardStatusRepository,
	loginRewardModelRepository masterLoginRewardRepository.LoginRewardModelRepository,
	loginRewardRewardRepository masterLoginRewardRepository.LoginRewardRewardRepository,
	eventService eventService.EventService,
	itemService itemService.ItemService,
) LoginRewardService {
	return &loginRewardService{
		transactionRepository:       transactionRepository,
		loginRewardStatusRepository: loginRewardStatusRepository,
		loginRewardModelRepository:  loginRewardModelRepository,
		loginRewardRewardRepository: loginRewardRewardRepository,
		eventService:                eventService,
		itemService:                 itemService,
	}
}

// GetLoginRewardModel ログイン報酬モデルを取得する
func (s *loginRewardService) GetLoginRewardModel(req *request.GetLoginRewardModel, now time.Time) (*response.GetLoginRewardModel, error) {
	lrm, lrrs, e, err := s.getLoginRewardModelAndRewardsAndEvent(req.LoginRewardModelName, now)
	if err != nil {
		return nil, err
	}

	rewards := make(response.LoginRewardRewards, len(*lrrs))
	for i, lrr := range *lrrs {
		reward := &response.LoginRewardReward{
			ID:         lrr.ID,
			ItemName:   lrr.ItemName,
			Name:       lrr.Name,
			StepNumber: lrr.StepNumber,
		}
		rewards[i] = *reward
	}

	return &response.GetLoginRewardModel{
		Status: 200,
		Item: response.LoginRewardModel{
			ID:   lrm.ID,
			Name: lrm.Name,
			Event: response.Event{
				ID:            e.ID,
				Name:          e.Name,
				RepeatSetting: e.RepeatSetting,
				ResetHour:     e.ResetHour,
				RepeatStartAt: e.RepeatStartAt,
				StartAt:       e.StartAt,
				EndAt:         e.EndAt,
			},
			LoginRewardRewards: rewards,
		},
	}, nil
}

// ReceiveLoginReward 受け取る
func (s *loginRewardService) ReceiveLoginReward(req *request.ReceiveLoginReward, now time.Time) (*response.ReceiveLoginReward, error) {
	// transaction
	tx, err := s.transactionRepository.Begin(req.ShardKey)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := s.transactionRepository.Rollback(tx); err != nil {
				log.Panicln(err)
			}
		} else {
			if err := s.transactionRepository.Commit(tx); err != nil {
				log.Panicln(err)
			}
		}
	}()

	// ログイン報酬モデルを取得
	lrm, lrrs, e, err := s.getLoginRewardModelAndRewardsAndEvent(req.LoginRewardModelName, now)
	if err != nil {
		return nil, err
	}

	// ログイン報酬ステータスを取得
	lrs, err := s.loginRewardStatusRepository.FindOrNilByLoginRewardModelName(lrm.Name, req.ShardKey)
	if err != nil {
		return nil, err
	}

	// 既に今日の報酬を受け取っている場合
	if lrs != nil && s.hasReceived(lrs.LastReceivedAt, now, *e.ResetHour) {
		return nil, errors.New("already received")
	}

	// アイテム取得
	item, err := s.getItem(e, lrrs, now)
	if err != nil {
		return nil, err
	}

	// アイテム受け取り
	if _, err := s.itemService.ReceiveItemInBox(
		&itemRequest.ReceiveItemInBox{
			ShardKey:  req.ShardKey,
			AccountID: req.AccountID,
			ItemName:  item.Name,
		},
	); err != nil {
		return nil, err
	}

	// 受け取りステータスを更新
	newLrs, err := s.updateLoginRewardStatus(lrs, now, lrm.Name, req.AccountID, req.ShardKey, tx)
	if err != nil {
		return nil, err
	}

	return &response.ReceiveLoginReward{
		Status: 200,
		Item: response.LoginRewardStatus{
			LoginRewardModel: response.LoginRewardModel{
				ID:   lrm.ID,
				Name: lrm.Name,
				Event: response.Event{
					ID:            e.ID,
					Name:          e.Name,
					RepeatSetting: e.RepeatSetting,
					StartAt:       e.StartAt,
					EndAt:         e.EndAt,
				},
			},
			Item: response.Item{
				ID:     item.ID,
				Name:   item.Name,
				Detail: item.Name,
			},
			LastReceivedAt: newLrs.LastReceivedAt,
		},
	}, nil
}

// ログイン報酬モデル、報酬一覧、イベントを取得
func (s *loginRewardService) getLoginRewardModelAndRewardsAndEvent(loginRewardModelName string, now time.Time) (*masterLoginRewardEntity.LoginRewardModel, *masterLoginRewardEntity.LoginRewardRewards, *masterEventEntity.Event, error) {
	lrm, err := s.loginRewardModelRepository.FindByName(loginRewardModelName)
	if err != nil {
		return nil, nil, nil, err
	}

	lrrs, err := s.loginRewardRewardRepository.ListByLoginRewardModelName(lrm.Name)
	if err != nil {
		return nil, nil, nil, err
	}

	e, err := s.eventService.GetEventToEntity(lrm.EventName, now)
	if err != nil {
		return nil, nil, nil, err
	}

	return lrm, lrrs, e, nil
}

// updateLoginRewardStatus 受け取りステータスを更新
func (s *loginRewardService) updateLoginRewardStatus(lrs *userLoginRewardEntity.LoginRewardStatus, now time.Time, loginRewardModelName string, accountID int64, shardKey int, tx *gorm.DB) (*userLoginRewardEntity.LoginRewardStatus, error) {
	if lrs != nil {
		lrs.LastReceivedAt = now
		lrs, err := s.loginRewardStatusRepository.Update(lrs, shardKey, tx)
		if err != nil {
			return nil, err
		}

		return lrs, nil
	}

	lrs, err := s.loginRewardStatusRepository.Create(
		&userLoginRewardEntity.LoginRewardStatus{
			ID:                   1,
			ShardKey:             shardKey,
			AccountID:            accountID,
			LoginRewardModelName: loginRewardModelName,
			LastReceivedAt:       now,
		}, shardKey, tx,
	)
	if err != nil {
		return nil, err
	}

	return lrs, nil
}

// getItem アイテムを取得
func (s *loginRewardService) getItem(e *masterEventEntity.Event, lrrs *masterLoginRewardEntity.LoginRewardRewards, now time.Time) (*masterItemEntity.Item, error) {
	maxStepNumber := s.getMaxStepNumber(lrrs)
	dayCount := s.getDayCount(e, now)
	if maxStepNumber < dayCount && maxStepNumber > 0 {
		dayCount %= maxStepNumber
	}

	var itemName string
	for _, rewards := range *lrrs {
		if dayCount == rewards.StepNumber {
			itemName = rewards.ItemName
		}
	}

	// アイテム取得
	i, err := s.itemService.GetItemToEntity(itemName)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// getDayCount 経過日数を取得
func (s *loginRewardService) getDayCount(e *masterEventEntity.Event, now time.Time) int {
	if e.RepeatSetting {
		return times.GetDayCount(*e.RepeatStartAt, now)
	}

	return times.GetDayCount(*e.StartAt, now)
}

// getMaxStepNumber ステップナンバーの最大値を取得
func (s *loginRewardService) getMaxStepNumber(lrrs *masterLoginRewardEntity.LoginRewardRewards) int {
	var maxStepNumber int
	for _, rewards := range *lrrs {
		if rewards.StepNumber > maxStepNumber {
			maxStepNumber = rewards.StepNumber
		}
	}

	return maxStepNumber
}

// hasReceived 最終受け取り日時(lastReceiveAt)が今日のリセット時間(RepeatHour)より後か判定
func (s *loginRewardService) hasReceived(lastReceiveAt, now time.Time, resetHour int) bool {
	return lastReceiveAt.After(time.Date(now.Year(), now.Month(), now.Day(), resetHour, 0, 0, 0, now.Location()))
}
