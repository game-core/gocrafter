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
	GetUser(ctx context.Context, req *loginBonusServer.LoginBonusGetUserRequest) (*loginBonusServer.LoginBonusGetUserResponse, error)
	GetMaster(ctx context.Context, req *loginBonusServer.LoginBonusGetMasterRequest) (*loginBonusServer.LoginBonusGetMasterResponse, error)
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

// GetUser ユーザーデータを取得する
func (s *loginBonusUsecase) GetUser(ctx context.Context, req *loginBonusServer.LoginBonusGetUserRequest) (*loginBonusServer.LoginBonusGetUserResponse, error) {
	result, err := s.loginBonusService.GetUser(ctx, loginBonusService.SetLoginBonusGetUserRequest(req.UserId))
	if err != nil {
		return nil, errors.NewMethodError("s.loginBonusService.GetUser", err)
	}

	return loginBonusServer.SetLoginBonusGetUserResponse(loginBonusServer.SetUserLoginBonuses(result.UserLoginBonuses)), nil
}

// GetMaster マスターデータを取得する
func (s *loginBonusUsecase) GetMaster(ctx context.Context, req *loginBonusServer.LoginBonusGetMasterRequest) (*loginBonusServer.LoginBonusGetMasterResponse, error) {
	result, err := s.loginBonusService.GetMaster(ctx, loginBonusService.SetLoginBonusGetMasterRequest(req.MasterLoginBonusId))
	if err != nil {
		return nil, errors.NewMethodError("s.loginBonusService.GetMaster", err)
	}

	return loginBonusServer.SetLoginBonusGetMasterResponse(
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
		loginBonusServer.SetMasterLoginBonusSchedules(result.MasterLoginBonusSchedules),
	), nil
}

// Receive ログインボーナス受け取り
func (s *loginBonusUsecase) Receive(ctx context.Context, req *loginBonusServer.LoginBonusReceiveRequest) (*loginBonusServer.LoginBonusReceiveResponse, error) {
	// transaction
	tx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
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
