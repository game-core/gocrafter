package account

import (
	"context"

	accountServer "github.com/game-core/gocrafter/api/game/presentation/server/account"
	accountService "github.com/game-core/gocrafter/pkg/domain/model/account"
)

type AccountUsecase interface {
	Create(ctx context.Context, req *accountServer.AccountCreateRequest) (*accountServer.AccountCreateResponse, error)
}

type accountUsecase struct {
	accountService     accountService.AccountService
	transactionService service.TransactionService
}
