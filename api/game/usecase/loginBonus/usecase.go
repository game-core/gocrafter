package loginBonus

import (
	"context"

	loginBonusServer "github.com/game-core/gocrafter/api/game/presentation/server/loginBonus"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/times"
	loginBonusService "github.com/game-core/gocrafter/pkg/domain/model/loginBonus"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

type LoginBonusUsecase interface {
	Receive(ctx context.Context, req *loginBonusServer.LoginBonusReceiveRequest) (*loginBonusServer.LoginBonusReceiveResponse, error)
}

type loginBonusUsecase struct {
	loginBonusService  loginBonusService.LoginBonusService
	transactionService transactionService.TransactionService
}

func NewLoginBonusUsecase(
	loginBonusService loginBonusService.LoginBonusService,
	transactionService transactionService.TransactionService,
) LoginBonusUsecase {
	return &loginBonusUsecase{
		loginBonusService:  loginBonusService,
		transactionService: transactionService,
	}
}

// Receive ログインボーナス受け取り
func (s *loginBonusUsecase) Receive(ctx context.Context, req *loginBonusServer.LoginBonusReceiveRequest) (*loginBonusServer.LoginBonusReceiveResponse, error) {
	// transaction
	tx, err := s.transactionService.UserBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserBegin", err)
	}
	defer func() {
		s.transactionService.UserEnd(ctx, tx, err)
	}()

	result, err := s.loginBonusService.Receive(ctx, tx, times.Now(), loginBonusService.SetLoginBonusReceiveRequest(req.UserId, req.MasterLoginBonusId))
	if err != nil {
		return nil, errors.NewMethodError("s.loginBonusService.Receive", err)
	}

	return loginBonusServer.SetLoginBonusReceiveResponse(
		loginBonusServer.SetUserLoginBonus(
			result.UserLoginBonus.UserId,
			result.UserLoginBonus.MasterLoginBonusId,
			times.TimeToPb(&result.UserLoginBonus.ReceivedAt),
		),
		loginBonusServer.SetMasterLoginBonus(
			result.MasterLoginBonus.Id,
			result.MasterLoginBonus.MasterLoginBonusEventId,
			result.MasterLoginBonus.Name,
		),
		loginBonusServer.SetMasterLoginBonusEvent(
			result.MasterLoginBonusEvent.Id,
			result.MasterLoginBonusEvent.Name,
			result.MasterLoginBonusEvent.ResetHour,
			result.MasterLoginBonusEvent.IntervalHour,
			result.MasterLoginBonusEvent.RepeatSetting,
			times.TimeToPb(&result.MasterLoginBonusEvent.StartAt),
			times.TimeToPb(result.MasterLoginBonusEvent.EndAt),
		),
		loginBonusServer.SetMasterLoginBonusItems(result.MasterLoginBonusItems),
		loginBonusServer.SetMasterLoginBonusSchedule(
			result.MasterLoginBonusSchedule.Id,
			result.MasterLoginBonusSchedule.MasterLoginBonusId,
			result.MasterLoginBonusSchedule.Step,
			result.MasterLoginBonusSchedule.Name,
		),
	), nil
}
