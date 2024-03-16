package idleBonus

import (
	"context"

	idleBonusServer "github.com/game-core/gocrafter/api/game/presentation/server/idleBonus"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/times"
	idleBonusService "github.com/game-core/gocrafter/pkg/domain/model/idleBonus"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

type IdleBonusUsecase interface {
	GetMaster(ctx context.Context, req *idleBonusServer.IdleBonusGetMasterRequest) (*idleBonusServer.IdleBonusGetMasterResponse, error)
	Receive(ctx context.Context, req *idleBonusServer.IdleBonusReceiveRequest) (*idleBonusServer.IdleBonusReceiveResponse, error)
}

type idleBonusUsecase struct {
	idleBonusService   idleBonusService.IdleBonusService
	transactionService transactionService.TransactionService
}

func NewIdleBonusUsecase(
	idleBonusService idleBonusService.IdleBonusService,
	transactionService transactionService.TransactionService,
) IdleBonusUsecase {
	return &idleBonusUsecase{
		idleBonusService:   idleBonusService,
		transactionService: transactionService,
	}
}

// GetMaster マスターデータを取得する
func (s *idleBonusUsecase) GetMaster(ctx context.Context, req *idleBonusServer.IdleBonusGetMasterRequest) (*idleBonusServer.IdleBonusGetMasterResponse, error) {
	result, err := s.idleBonusService.GetMaster(ctx, idleBonusService.SetIdleBonusGetMasterRequest(req.MasterIdleBonusId))
	if err != nil {
		return nil, errors.NewMethodError("s.idleBonusService.GetMaster", err)
	}

	return idleBonusServer.SetIdleBonusGetMasterResponse(
		idleBonusServer.SetMasterIdleBonus(
			result.MasterIdleBonus.Id,
			result.MasterIdleBonus.MasterIdleBonusEventId,
			result.MasterIdleBonus.Name,
		),
		idleBonusServer.SetMasterIdleBonusEvent(
			result.MasterIdleBonusEvent.Id,
			result.MasterIdleBonusEvent.Name,
			result.MasterIdleBonusEvent.ResetHour,
			result.MasterIdleBonusEvent.IntervalHour,
			result.MasterIdleBonusEvent.RepeatSetting,
			times.TimeToPb(&result.MasterIdleBonusEvent.StartAt),
			times.TimeToPb(result.MasterIdleBonusEvent.EndAt),
		),
		idleBonusServer.SetMasterIdleBonusItems(result.MasterIdleBonusItems),
		idleBonusServer.SetMasterIdleBonusSchedules(result.MasterIdleBonusSchedules),
	), nil
}

// Receive 放置ボーナス受け取り
func (s *idleBonusUsecase) Receive(ctx context.Context, req *idleBonusServer.IdleBonusReceiveRequest) (*idleBonusServer.IdleBonusReceiveResponse, error) {
	// transaction
	tx, err := s.transactionService.UserBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserBegin", err)
	}
	defer func() {
		s.transactionService.UserEnd(ctx, tx, err)
	}()

	result, err := s.idleBonusService.Receive(ctx, tx, times.Now(), idleBonusService.SetIdleBonusReceiveRequest(req.UserId, req.MasterIdleBonusId))
	if err != nil {
		return nil, errors.NewMethodError("s.idleBonusService.Receive", err)
	}

	return idleBonusServer.SetIdleBonusReceiveResponse(
		idleBonusServer.SetUserIdleBonus(
			result.UserIdleBonus.UserId,
			result.UserIdleBonus.MasterIdleBonusId,
			times.TimeToPb(&result.UserIdleBonus.ReceivedAt),
		),
		idleBonusServer.SetMasterIdleBonus(
			result.MasterIdleBonus.Id,
			result.MasterIdleBonus.MasterIdleBonusEventId,
			result.MasterIdleBonus.Name,
		),
		idleBonusServer.SetMasterIdleBonusEvent(
			result.MasterIdleBonusEvent.Id,
			result.MasterIdleBonusEvent.Name,
			result.MasterIdleBonusEvent.ResetHour,
			result.MasterIdleBonusEvent.IntervalHour,
			result.MasterIdleBonusEvent.RepeatSetting,
			times.TimeToPb(&result.MasterIdleBonusEvent.StartAt),
			times.TimeToPb(result.MasterIdleBonusEvent.EndAt),
		),
		idleBonusServer.SetMasterIdleBonusItems(result.MasterIdleBonusItems),
		idleBonusServer.SetMasterIdleBonusSchedules(result.MasterIdleBonusSchedules),
	), nil
}
