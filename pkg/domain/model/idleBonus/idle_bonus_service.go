//go:generate mockgen -source=./idle_bonus_service.go -destination=./idle_bonus_service_mock.gen.go -package=idleBonus
package idleBonus

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusSchedule"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/userIdleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/item"
)

type IdleBonusService interface {
	Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *IdleBonusReceiveRequest) (*IdleBonusReceiveResponse, error)
}

type idleBonusService struct {
	itemService                       item.ItemService
	userIdleBonusRepository           userIdleBonus.UserIdleBonusRepository
	masterIdleBonusRepository         masterIdleBonus.MasterIdleBonusRepository
	masterIdleBonusEventRepository    masterIdleBonusEvent.MasterIdleBonusEventRepository
	masterIdleBonusItemRepository     masterIdleBonusItem.MasterIdleBonusItemRepository
	masterIdleBonusScheduleRepository masterIdleBonusSchedule.MasterIdleBonusScheduleRepository
}

func NewIdleBonusService(
	itemService item.ItemService,
	userIdleBonusRepository userIdleBonus.UserIdleBonusRepository,
	masterIdleBonusRepository masterIdleBonus.MasterIdleBonusRepository,
	masterIdleBonusEventRepository masterIdleBonusEvent.MasterIdleBonusEventRepository,
	masterIdleBonusItemRepository masterIdleBonusItem.MasterIdleBonusItemRepository,
	masterIdleBonusScheduleRepository masterIdleBonusSchedule.MasterIdleBonusScheduleRepository,
) IdleBonusService {
	return &idleBonusService{
		itemService:                       itemService,
		userIdleBonusRepository:           userIdleBonusRepository,
		masterIdleBonusRepository:         masterIdleBonusRepository,
		masterIdleBonusEventRepository:    masterIdleBonusEventRepository,
		masterIdleBonusItemRepository:     masterIdleBonusItemRepository,
		masterIdleBonusScheduleRepository: masterIdleBonusScheduleRepository,
	}
}

// Receive 放置ボーナスを受け取る
func (s *idleBonusService) Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *IdleBonusReceiveRequest) (*IdleBonusReceiveResponse, error) {
	masterIdleBonusModel, err := s.masterIdleBonusRepository.Find(ctx, req.MasterIdleBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterIdleBonusRepository.Find", err)
	}

	masterIdleBonusEventModel, err := s.getEvent(ctx, now, masterIdleBonusModel.MasterIdleBonusEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.getEvent", err)
	}

	userIdleBonusModel, err := s.userIdleBonusRepository.FindOrNil(ctx, req.UserId, req.MasterIdleBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.userIdleBonusRepository.FindOrNil", err)
	}

	masterIdleBonusScheduleModel, err := s.getSchedules(ctx, now, req.MasterIdleBonusId, masterIdleBonusEventModel.IntervalHour, userIdleBonusModel.ReceivedAt)
	if err != nil {
		return nil, errors.NewMethodError("s.getSchedules", err)
	}

	masterIdleBonusItemModels, err := s.getItems(ctx, masterIdleBonusScheduleModel)
	if err != nil {
		return nil, errors.NewMethodError("s.getItems", err)
	}

	if err := s.receive(ctx, tx, req.UserId, masterIdleBonusItemModels); err != nil {
		return nil, errors.NewMethodError("s.receive", err)
	}

	result, err := s.update(ctx, tx, now, req.UserId, req.MasterIdleBonusId, userIdleBonusModel)
	if err != nil {
		return nil, errors.NewMethodError("s.update", err)
	}

	return SetIdleBonusReceiveResponse(
		result,
		masterIdleBonusModel,
		masterIdleBonusEventModel,
		masterIdleBonusItemModels,
		masterIdleBonusScheduleModel,
	), nil
}

// getEvent イベントを取得する
func (s *idleBonusService) getEvent(ctx context.Context, now time.Time, masterIdleBonusEventId int64) (*masterIdleBonusEvent.MasterIdleBonusEvent, error) {
	masterIdleBonusEvent, err := s.masterIdleBonusEventRepository.Find(ctx, masterIdleBonusEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterIdleBonusEventRepository.FindByMasterIdleBonusId", err)
	}

	// イベント期間外の場合
	if !masterIdleBonusEvent.CheckEventPeriod(now) {
		return nil, errors.NewError("outside the event period")
	}

	return masterIdleBonusEvent, nil
}

// getSchedules スケジュール一覧を取得する
func (s *idleBonusService) getSchedules(ctx context.Context, now time.Time, masterIdleBonusId int64, intervalHour int32, receivedAt time.Time) (masterIdleBonusSchedule.MasterIdleBonusSchedules, error) {
	masterIdleBonusSchedules, err := s.masterIdleBonusScheduleRepository.FindListByMasterIdleBonusId(ctx, masterIdleBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterIdleBonusScheduleRepository.FindListByMasterIdleBonusId", err)
	}

	return masterIdleBonusSchedules.GetSchedulesByStep(masterIdleBonusSchedules.GetStep(intervalHour, receivedAt, now)), nil
}

// getItems アイテム一覧を取得する
func (s *idleBonusService) getItems(ctx context.Context, masterIdleBonusScheduleModels masterIdleBonusSchedule.MasterIdleBonusSchedules) (masterIdleBonusItem.MasterIdleBonusItems, error) {
	result := masterIdleBonusItem.NewMasterIdleBonusItems()

	for _, masterIdleBonusScheduleModel := range masterIdleBonusScheduleModels {
		masterIdleBonusItemModels, err := s.masterIdleBonusItemRepository.FindListByMasterIdleBonusScheduleId(ctx, masterIdleBonusScheduleModel.Id)
		if err != nil {
			return nil, errors.NewMethodError("s.masterIdleBonusItemRepository.FindListByMasterIdleBonusScheduleId", err)
		}

		result = append(result, masterIdleBonusItemModels...)
	}

	return result, nil
}

// receive 受け取り
func (s *idleBonusService) receive(ctx context.Context, tx *gorm.DB, userId string, masterIdleBonusItemModels masterIdleBonusItem.MasterIdleBonusItems) error {
	var items item.Items

	for _, masterIdleBonusItem := range masterIdleBonusItemModels {
		items = append(items, item.SetItem(masterIdleBonusItem.MasterItemId, masterIdleBonusItem.Count))
	}

	if _, err := s.itemService.Receive(ctx, tx, item.SetItemReceiveRequest(userId, items)); err != nil {
		return errors.NewMethodError("s.itemService.Receive", err)
	}

	return nil
}

// update ユーザーログインボーナスを更新
func (s *idleBonusService) update(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterIdleBonusId int64, userIdleBonusModel *userIdleBonus.UserIdleBonus) (*userIdleBonus.UserIdleBonus, error) {
	if userIdleBonusModel != nil {
		userIdleBonusModel.ReceivedAt = now
		result, err := s.userIdleBonusRepository.Update(ctx, tx, userIdleBonusModel)
		if err != nil {
			return nil, errors.NewMethodError("s.userIdleBonusRepository.Update", err)
		}

		return result, nil
	}

	result, err := s.userIdleBonusRepository.Create(ctx, tx, userIdleBonus.SetUserIdleBonus(userId, masterIdleBonusId, now))
	if err != nil {
		return nil, errors.NewMethodError("s.userIdleBonusRepository.Create", err)
	}

	return result, nil
}
