package account

import (
	"context"

	accountServer "github.com/game-core/gocrafter/api/game/presentation/server/account"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/times"
	accountService "github.com/game-core/gocrafter/pkg/domain/model/account"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

type AccountUsecase interface {
	Create(ctx context.Context, req *accountServer.AccountCreateRequest) (*accountServer.AccountCreateResponse, error)
	Login(ctx context.Context, req *accountServer.AccountLoginRequest) (*accountServer.AccountLoginResponse, error)
	Check(ctx context.Context, req *accountServer.AccountCheckRequest) (*accountServer.AccountCheckResponse, error)
}

type accountUsecase struct {
	accountService     accountService.AccountService
	transactionService transactionService.TransactionService
}

func NewAccountUsecase(
	accountService accountService.AccountService,
	transactionService transactionService.TransactionService,
) AccountUsecase {
	return &accountUsecase{
		accountService:     accountService,
		transactionService: transactionService,
	}
}

// Create アカウントを作成する
func (s *accountUsecase) Create(ctx context.Context, req *accountServer.AccountCreateRequest) (*accountServer.AccountCreateResponse, error) {
	userId, err := s.accountService.GenerateUserID(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.GenerateUserID", err)
	}

	// transaction
	tx, err := s.transactionService.UserMysqlBegin(ctx, userId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
	}()

	result, err := s.accountService.Create(ctx, tx, accountService.SetAccountCreateRequest(userId, req.Name))
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.Create", err)
	}

	return accountServer.SetAccountCreateResponse(
		accountServer.SetUserAccount(
			result.UserAccount.UserId,
			result.UserAccount.Name,
			result.UserAccount.Password,
			times.TimeToPb(&result.UserAccount.LoginAt),
			times.TimeToPb(&result.UserAccount.LogoutAt),
		),
	), nil
}

// Login アカウントをログインする
func (s *accountUsecase) Login(ctx context.Context, req *accountServer.AccountLoginRequest) (*accountServer.AccountLoginResponse, error) {
	// transaction
	tx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
	}()

	result, err := s.accountService.Login(ctx, tx, accountService.SetAccountLoginRequest(req.UserId, req.Name, req.Password))
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.Login", err)
	}

	return accountServer.SetAccountLoginResponse(
		result.Token,
		accountServer.SetUserAccount(
			result.UserAccount.UserId,
			result.UserAccount.Name,
			result.UserAccount.Password,
			times.TimeToPb(&result.UserAccount.LoginAt),
			times.TimeToPb(&result.UserAccount.LogoutAt),
		),
	), nil
}

// Check アカウントを確認する
func (s *accountUsecase) Check(ctx context.Context, req *accountServer.AccountCheckRequest) (*accountServer.AccountCheckResponse, error) {
	result, err := s.accountService.Check(ctx, accountService.SetAccountCheckRequest(req.UserId))
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.Check", err)
	}

	return accountServer.SetAccountCheckResponse(
		accountServer.SetUserAccount(
			result.UserAccount.UserId,
			result.UserAccount.Name,
			result.UserAccount.Password,
			times.TimeToPb(&result.UserAccount.LoginAt),
			times.TimeToPb(&result.UserAccount.LogoutAt),
		),
	), nil
}
