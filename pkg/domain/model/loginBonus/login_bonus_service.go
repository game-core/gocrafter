//go:generate mockgen -source=./login_bonus_service.go -destination=./login_bonus_service_mock.gen.go -package=loginBonus
package loginBonus

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/item"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusSchedule"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/userLoginBonus"
)

type LoginBonusService interface {
	Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *LoginBonusReceiveRequest) (*LoginBonusReceiveResponse, error)
}

type loginBonusService struct {
	itemService                        item.ItemService
	userLoginBonusRepository           userLoginBonus.UserLoginBonusRepository
	masterLoginBonusRepository         masterLoginBonus.MasterLoginBonusRepository
	masterLoginBonusEventRepository    masterLoginBonusEvent.MasterLoginBonusEventRepository
	masterLoginBonusItemRepository     masterLoginBonusItem.MasterLoginBonusItemRepository
	masterLoginBonusScheduleRepository masterLoginBonusSchedule.MasterLoginBonusScheduleRepository
}

func NewLoginBonusService(
	itemService item.ItemService,
	userLoginBonusRepository userLoginBonus.UserLoginBonusRepository,
	masterLoginBonusRepository masterLoginBonus.MasterLoginBonusRepository,
	masterLoginBonusEventRepository masterLoginBonusEvent.MasterLoginBonusEventRepository,
	masterLoginBonusItemRepository masterLoginBonusItem.MasterLoginBonusItemRepository,
	masterLoginBonusScheduleRepository masterLoginBonusSchedule.MasterLoginBonusScheduleRepository,
) LoginBonusService {
	return &loginBonusService{
		itemService:                        itemService,
		userLoginBonusRepository:           userLoginBonusRepository,
		masterLoginBonusRepository:         masterLoginBonusRepository,
		masterLoginBonusEventRepository:    masterLoginBonusEventRepository,
		masterLoginBonusItemRepository:     masterLoginBonusItemRepository,
		masterLoginBonusScheduleRepository: masterLoginBonusScheduleRepository,
	}
}

// Receive ログインボーナスを受け取る
func (s *loginBonusService) Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *LoginBonusReceiveRequest) (*LoginBonusReceiveResponse, error) {
	masterLoginBonusModel, err := s.masterLoginBonusRepository.Find(ctx, req.MasterLoginBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusRepository.Find", err)
	}

	masterLoginBonusEventModel, err := s.getEvent(ctx, now, masterLoginBonusModel.MasterLoginBonusEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.getEvent", err)
	}

	masterLoginBonusScheduleModel, err := s.getSchedule(ctx, now, req.MasterLoginBonusId, masterLoginBonusEventModel.IntervalHour, masterLoginBonusEventModel.StartAt)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusScheduleRepository.FindListByMasterLoginBonusId", err)
	}

	masterLoginBonusItemModels, err := s.masterLoginBonusItemRepository.FindListByMasterLoginBonusScheduleId(ctx, masterLoginBonusScheduleModel.Id)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusItemRepository.FindByMasterLoginBonusScheduleId", err)
	}

	userLoginBonusModel, err := s.userLoginBonusRepository.FindOrNil(ctx, req.UserId, req.MasterLoginBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.userLoginBonusRepository.Find", err)
	}
	if userLoginBonusModel != nil && userLoginBonusModel.CheckReceived(masterLoginBonusEventModel.ResetHour, now) {
		return nil, errors.NewError("already received")
	}

	if err := s.receive(ctx, tx, req.UserId, masterLoginBonusItemModels); err != nil {
		return nil, errors.NewMethodError("s.receive", err)
	}

	result, err := s.update(ctx, tx, now, req.UserId, req.MasterLoginBonusId, userLoginBonusModel)
	if err != nil {
		return nil, errors.NewMethodError("s.update", err)
	}

	return SetLoginBonusReceiveResponse(
		result,
		masterLoginBonusModel,
		masterLoginBonusEventModel,
		masterLoginBonusItemModels,
		masterLoginBonusScheduleModel,
	), nil
}

// getEvent イベントを取得する
func (s *loginBonusService) getEvent(ctx context.Context, now time.Time, masterLoginBonusEventId int64) (*masterLoginBonusEvent.MasterLoginBonusEvent, error) {
	masterLoginBonusEvent, err := s.masterLoginBonusEventRepository.Find(ctx, masterLoginBonusEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusEventRepository.FindByMasterLoginBonusId", err)
	}

	// イベント期間外の場合
	if !masterLoginBonusEvent.CheckEventPeriod(now) {
		return nil, errors.NewError("outside the event period")
	}

	return masterLoginBonusEvent, nil
}

// getSchedule スケジュールを取得する
func (s *loginBonusService) getSchedule(ctx context.Context, now time.Time, masterLoginBonusId int64, intervalHour int32, startAt time.Time) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	masterLoginBonusSchedules, err := s.masterLoginBonusScheduleRepository.FindListByMasterLoginBonusId(ctx, masterLoginBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusScheduleRepository.FindListByMasterLoginBonusId", err)
	}

	return masterLoginBonusSchedules.GetScheduleByStep(masterLoginBonusSchedules.GetStep(intervalHour, startAt, now)), nil
}

// receive 受け取り
func (s *loginBonusService) receive(ctx context.Context, tx *gorm.DB, userId string, masterLoginBonusItemModels masterLoginBonusItem.MasterLoginBonusItems) error {
	var items item.Items

	for _, masterLoginBonusItem := range masterLoginBonusItemModels {
		items = append(items, item.SetItem(masterLoginBonusItem.MasterItemId, masterLoginBonusItem.Count))
	}

	if _, err := s.itemService.Receive(ctx, tx, item.SetItemReceiveRequest(userId, items)); err != nil {
		return errors.NewMethodError("s.itemService.Receive", err)
	}

	return nil
}

// update ユーザーログインボーナスを更新
func (s *loginBonusService) update(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterLoginBonusId int64, userLoginBonusModel *userLoginBonus.UserLoginBonus) (*userLoginBonus.UserLoginBonus, error) {
	if userLoginBonusModel != nil {
		userLoginBonusModel.ReceivedAt = now
		result, err := s.userLoginBonusRepository.Update(ctx, tx, userLoginBonusModel)
		if err != nil {
			return nil, errors.NewMethodError("s.userLoginBonusRepository.Update", err)
		}

		return result, nil
	}

	result, err := s.userLoginBonusRepository.Create(ctx, tx, userLoginBonus.SetUserLoginBonus(userId, masterLoginBonusId, now))
	if err != nil {
		return nil, errors.NewMethodError("s.userLoginBonusRepository.Update", err)
	}

	return result, nil
}
