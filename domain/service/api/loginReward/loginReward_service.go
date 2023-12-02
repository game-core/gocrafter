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
	masterLoginRewardEntity "github.com/game-core/gocrafter/domain/entity/master/loginReward"
	userLoginRewardEntity "github.com/game-core/gocrafter/domain/entity/user/loginReward"
	masterLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/master/loginReward"
	userRepository "github.com/game-core/gocrafter/domain/repository/user"
	userLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/user/loginReward"
	eventService "github.com/game-core/gocrafter/domain/service/api/event"
	itemService "github.com/game-core/gocrafter/domain/service/api/item"
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

	rewards, err := response.ToRewards(lrrs)
	if err != nil {
		return nil, err
	}

	return response.ToGetLoginRewardModel(
		200,
		*response.ToLoginRewardModel(
			lrm.ID,
			lrm.Name,
			*response.ToEvent(
				e.ID,
				e.Name,
				e.ResetHour,
				e.RepeatSetting,
				e.RepeatStartAt,
				e.StartAt,
				e.EndAt,
			),
			rewards,
		),
	), nil
}

// ReceiveLoginReward 受け取る
func (s *loginRewardService) ReceiveLoginReward(req *request.ReceiveLoginReward, now time.Time) (*response.ReceiveLoginReward, error) {
	// transaction
	tx, err := s.transactionRepository.Begin(req.ShardKey)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := s.transactionRepository.CommitOrRollback(tx, err); err != nil {
			log.Panicln(err)
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

	lrs, err = s.receive(lrs, lrrs, e, req, now, tx)
	if err != nil {
		return nil, err
	}

	rewards, err := response.ToRewards(lrrs)
	if err != nil {
		return nil, err
	}

	items, err := response.ToItems(lrrs.GetItems(e.GetDayCount(now)))
	if err != nil {
		return nil, err
	}

	return response.ToReceiveLoginReward(
		200,
		*response.ToLoginRewardStatus(
			lrs.ID,
			*response.ToLoginRewardModel(
				lrm.ID,
				lrm.Name,
				*response.ToEvent(
					e.ID,
					e.Name,
					e.ResetHour,
					e.RepeatSetting,
					e.RepeatStartAt,
					e.StartAt,
					e.EndAt,
				),
				rewards,
			),
			items,
			lrs.LastReceivedAt,
		),
	), nil
}

// getLoginRewardModelAndRewardsAndEvent ログイン報酬モデル、報酬一覧、イベントを取得
func (s *loginRewardService) getLoginRewardModelAndRewardsAndEvent(loginRewardModelName string, now time.Time) (*masterLoginRewardEntity.LoginRewardModel, *masterLoginRewardEntity.LoginRewardRewards, *masterEventEntity.Event, error) {
	lrm, err := s.loginRewardModelRepository.FindByName(loginRewardModelName)
	if err != nil {
		return nil, nil, nil, err
	}

	lrrs, err := s.loginRewardRewardRepository.ListByLoginRewardModelName(lrm.Name)
	if err != nil {
		return nil, nil, nil, err
	}

	e, err := s.eventService.GetEventToEntity(lrm.EventName)
	if err != nil {
		return nil, nil, nil, err
	}

	if !e.IsEventPeriod(now) {
		return nil, nil, nil, errors.New("outside the event period")
	}

	return lrm, lrrs, e, nil
}

// receive 受け取り
func (s *loginRewardService) receive(lrs *userLoginRewardEntity.LoginRewardStatus, lrrs *masterLoginRewardEntity.LoginRewardRewards, e *masterEventEntity.Event, req *request.ReceiveLoginReward, now time.Time, tx *gorm.DB) (*userLoginRewardEntity.LoginRewardStatus, error) {
	if lrs != nil && !lrs.HasReceived(now, e.ResetHour) {
		return nil, errors.New("already received")
	}

	if err := s.receiveItem(lrrs, e, now, req.AccountID, req.ShardKey); err != nil {
		return nil, err
	}

	res, err := s.updateLoginRewardStatus(lrs, now, req.LoginRewardModelName, req.AccountID, req.ShardKey, tx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// receiveItem アイテムを受け取り
func (s *loginRewardService) receiveItem(lrrs *masterLoginRewardEntity.LoginRewardRewards, e *masterEventEntity.Event, now time.Time, accountID int64, shardKey string) error {
	rewardItems := &masterLoginRewardEntity.LoginRewardItems{}
	if err := rewardItems.ToEntities(lrrs.GetItems(e.GetDayCount(now))); err != nil {
		return err
	}

	var items itemRequest.Items
	for _, ri := range *rewardItems {
		item := itemRequest.Item{
			Name:  ri.Name,
			Count: ri.Count,
		}
		items = append(items, item)
	}

	if _, err := s.itemService.ReceiveItemInBox(
		&itemRequest.ReceiveItemInBox{
			ShardKey:  shardKey,
			AccountID: accountID,
			Items:     items,
		},
	); err != nil {
		return err
	}

	return nil
}

// updateLoginRewardStatus 受け取りステータスを更新
func (s *loginRewardService) updateLoginRewardStatus(lrs *userLoginRewardEntity.LoginRewardStatus, now time.Time, loginRewardModelName string, accountID int64, shardKey string, tx *gorm.DB) (*userLoginRewardEntity.LoginRewardStatus, error) {
	if lrs == nil {
		lrs = &userLoginRewardEntity.LoginRewardStatus{
			ShardKey:             shardKey,
			AccountID:            accountID,
			LoginRewardModelName: loginRewardModelName,
			LastReceivedAt:       now,
		}
	} else {
		lrs.LoginRewardModelName = loginRewardModelName
		lrs.LastReceivedAt = now
	}

	res, err := s.loginRewardStatusRepository.Save(lrs, shardKey, tx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
