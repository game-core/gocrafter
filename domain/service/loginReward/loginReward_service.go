//go:generate mockgen -source=./loginReward_service.go -destination=./loginReward_service_mock.gen.go -package=loginReward
package loginReward

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	itemRequest "github.com/game-core/gocrafter/api/presentation/request/item"
	request "github.com/game-core/gocrafter/api/presentation/request/loginReward"
	response "github.com/game-core/gocrafter/api/presentation/response/loginReward"
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

	lrm, lrrs, e, err := s.getLoginRewardModelAndRewardsAndEvent(req.LoginRewardModelName, now)
	if err != nil {
		return nil, err
	}

	lrs, err := s.loginRewardStatusRepository.FindOrNilByLoginRewardModelName(lrm.Name, req.ShardKey)
	if err != nil {
		return nil, err
	}

	newLrs, item, err := s.receive(lrs, lrrs, e, now, req, tx)
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

// getLoginRewardModelAndRewardsAndEvent ログイン報酬モデル、報酬一覧、イベントを取得
func (s *loginRewardService) getLoginRewardModelAndRewardsAndEvent(
	loginRewardModelName string,
	now time.Time,
) (*masterLoginRewardEntity.LoginRewardModel, *masterLoginRewardEntity.LoginRewardRewards, *masterEventEntity.Event, error) {
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

// receive 受け取り
func (s *loginRewardService) receive(
	lrs *userLoginRewardEntity.LoginRewardStatus,
	lrrs *masterLoginRewardEntity.LoginRewardRewards,
	e *masterEventEntity.Event,
	now time.Time,
	req *request.ReceiveLoginReward,
	tx *gorm.DB,
) (*userLoginRewardEntity.LoginRewardStatus, *masterItemEntity.Item, error) {
	if lrs != nil && !lrs.HasReceived(now, *e.ResetHour) {
		return nil, nil, errors.New("already received")
	}

	item, err := s.receiveItem(lrrs, e, now, req.AccountID, req.ShardKey)
	if err != nil {
		return nil, nil, err
	}

	newLrs, err := s.updateLoginRewardStatus(lrs, now, req.LoginRewardModelName, req.AccountID, req.ShardKey, tx)
	if err != nil {
		return nil, nil, err
	}

	return newLrs, item, nil
}

// receiveItem アイテムを受け取り
func (s *loginRewardService) receiveItem(
	lrrs *masterLoginRewardEntity.LoginRewardRewards,
	e *masterEventEntity.Event,
	now time.Time,
	accountID int64,
	shardKey int,
) (*masterItemEntity.Item, error) {
	item, err := s.itemService.GetItemToEntity(lrrs.GetItemName(e.GetDayCount(now)))
	if err != nil {
		return nil, err
	}

	if _, err := s.itemService.ReceiveItemInBox(
		&itemRequest.ReceiveItemInBox{
			ShardKey:  shardKey,
			AccountID: accountID,
			ItemName:  item.Name,
		},
	); err != nil {
		return nil, err
	}

	return item, nil
}

// updateLoginRewardStatus 受け取りステータスを更新
func (s *loginRewardService) updateLoginRewardStatus(
	lrs *userLoginRewardEntity.LoginRewardStatus,
	now time.Time,
	loginRewardModelName string,
	accountID int64,
	shardKey int,
	tx *gorm.DB,
) (*userLoginRewardEntity.LoginRewardStatus, error) {
	if lrs != nil {
		lrs.LastReceivedAt = now
		lrs.LoginRewardModelName = loginRewardModelName
		lrs, err := s.loginRewardStatusRepository.Save(lrs, shardKey, tx)
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
