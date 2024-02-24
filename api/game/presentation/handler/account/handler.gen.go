package account

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/account"
	accountUsecase "github.com/game-core/gocrafter/api/game/usecase/account"
	"github.com/game-core/gocrafter/internal/errors"
)

type AccountHandler interface {
	account.AccountServer
}

type accountHandler struct {
	account.UnimplementedAccountServer
	accountUsecase accountUsecase.AccountUsecase
}

func NewAccountHandler(
	accountUsecase accountUsecase.AccountUsecase,
) AccountHandler {
	return &accountHandler{
		accountUsecase: accountUsecase,
	}
}

// Create アカウントを作成する
func (s *accountHandler) Create(ctx context.Context, req *account.AccountCreateRequest) (*account.AccountCreateResponse, error) {
	res, err := s.accountUsecase.Create(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.accountUsecase.Create", err)
	}

	return res, nil
}
