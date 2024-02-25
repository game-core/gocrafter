//go:generate mockgen -source=./login_bonus_service.go -destination=./login_bonus_service_mock.gen.go -package=loginBonus
package loginBonus

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
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
	userLoginBonusRepository           userLoginBonus.UserLoginBonusRepository
	masterLoginBonusRepository         masterLoginBonus.MasterLoginBonusRepository
	masterLoginBonusEventRepository    masterLoginBonusEvent.MasterLoginBonusEventRepository
	masterLoginBonusItemRepository     masterLoginBonusItem.MasterLoginBonusItemRepository
	masterLoginBonusScheduleRepository masterLoginBonusSchedule.MasterLoginBonusScheduleRepository
}

func NewLoginBonusService(
	userLoginBonusRepository userLoginBonus.UserLoginBonusRepository,
	masterLoginBonusRepository masterLoginBonus.MasterLoginBonusRepository,
	masterLoginBonusEventRepository masterLoginBonusEvent.MasterLoginBonusEventRepository,
	masterLoginBonusItemRepository masterLoginBonusItem.MasterLoginBonusItemRepository,
	masterLoginBonusScheduleRepository masterLoginBonusSchedule.MasterLoginBonusScheduleRepository,
) LoginBonusService {
	return &loginBonusService{
		userLoginBonusRepository:           userLoginBonusRepository,
		masterLoginBonusRepository:         masterLoginBonusRepository,
		masterLoginBonusEventRepository:    masterLoginBonusEventRepository,
		masterLoginBonusItemRepository:     masterLoginBonusItemRepository,
		masterLoginBonusScheduleRepository: masterLoginBonusScheduleRepository,
	}
}

// Receive ログインボーナスを受け取る
func (s *loginBonusService) Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *LoginBonusReceiveRequest) (*LoginBonusReceiveResponse, error) {
	masterLoginBonus, err := s.masterLoginBonusRepository.Find(ctx, req.MasterLoginBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusRepository.Find", err)
	}

	masterLoginBonusEvent, err := s.getEvent(ctx, now, req.MasterLoginBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.getEvent", err)
	}

	masterLoginBonusSchedule, err := s.getSchedule(ctx, now, req.MasterLoginBonusId, masterLoginBonusEvent.IntervalHour, masterLoginBonusEvent.StartAt)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusScheduleRepository.FindListByMasterLoginBonusId", err)
	}

	masterLoginBonusItems, err := s.masterLoginBonusItemRepository.FindListByMasterLoginBonusScheduleId(ctx, masterLoginBonusSchedule.Id)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusItemRepository.FindByMasterLoginBonusScheduleId", err)
	}

	userLoginBonus, err := s.userLoginBonusRepository.FindOrNil(ctx, req.UserId, req.MasterLoginBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.userLoginBonusRepository.Find", err)
	}
	if userLoginBonus != nil && userLoginBonus.CheckReceived(masterLoginBonusEvent.ResetHour, now) {
		return nil, errors.NewError("already received")
	}

	if err := s.receive(ctx, tx, req.UserId, masterLoginBonusItems); err != nil {

	}

	return nil, nil
}

// getEvent イベントを取得する
func (s *loginBonusService) getEvent(ctx context.Context, now time.Time, masterLoginBonusId int64) (*masterLoginBonusEvent.MasterLoginBonusEvent, error) {
	masterLoginBonusEvent, err := s.masterLoginBonusEventRepository.FindByMasterLoginBonusId(ctx, masterLoginBonusId)
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
func (s *loginBonusService) receive(ctx context.Context, tx *gorm.DB, userId string, items masterLoginBonusItem.MasterLoginBonusItems) error {
	return nil
}
