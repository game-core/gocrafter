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
	tx, err := s.transactionService.UserBegin(ctx, userId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserBegin", err)
	}
	defer func() {
		s.transactionService.UserEnd(ctx, tx, err)
	}()

	userAccount, err := s.accountService.Create(ctx, tx, accountService.SetAccountCreateRequest(userId, req.Name))
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.Create", err)
	}

	return accountServer.SetAccountCreateResponse(
		accountServer.SetUserAccount(
			userAccount.UserAccount.UserId,
			userAccount.UserAccount.Name,
			userAccount.UserAccount.Password,
			"",
			times.TimeToPb(&userAccount.UserAccount.LoginAt),
			times.TimeToPb(&userAccount.UserAccount.LogoutAt),
		),
	), nil
}

// Login アカウントをログインする
func (s *accountUsecase) Login(ctx context.Context, req *accountServer.AccountLoginRequest) (*accountServer.AccountLoginResponse, error) {
	// transaction
	tx, err := s.transactionService.UserBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserBegin", err)
	}
	defer func() {
		s.transactionService.UserEnd(ctx, tx, err)
	}()

	return nil, nil
}

// Check アカウントを確認する
func (s *accountUsecase) Check(ctx context.Context, req *accountServer.AccountCheckRequest) (*accountServer.AccountCheckResponse, error) {
	return nil, nil
}
