package account

import (
	"context"
	"fmt"

	accountServer "github.com/game-core/gocrafter/api/game/presentation/server/account"
	"github.com/game-core/gocrafter/internal"
	accountService "github.com/game-core/gocrafter/pkg/domain/model/account"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

type AccountUsecase interface {
	Create(ctx context.Context, req *accountServer.AccountCreateRequest) (*accountServer.AccountCreateResponse, error)
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
	// transaction
	tx, err := s.transactionService.UserBegin(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to s.transactionService.UserBegin: %s", err)
	}
	defer func() {
		s.transactionService.UserEnd(ctx, tx, err)
	}()

	userAccount, err := s.accountService.Create(ctx, tx, accountService.SetAccountCreateRequest(req.UserId, req.Name))
	if err != nil {
		return nil, fmt.Errorf("failed to s.accountService.Create: %s", err)
	}

	return accountServer.SetAccountCreateResponse(
		accountServer.SetUserAccount(
			userAccount.UserAccount.UserId,
			userAccount.UserAccount.Name,
			userAccount.UserAccount.Password,
			internal.TimeToPb(&userAccount.UserAccount.LoginAt),
			internal.TimeToPb(&userAccount.UserAccount.LogoutAt),
		),
	), nil
}
