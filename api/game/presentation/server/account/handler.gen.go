package account

import (
	"context"
	"fmt"

	accountUsecase "github.com/game-core/gocrafter/api/game/usecase/account"
)

type AccountHandler interface {
	AccountServer
}

type accountHandler struct {
	UnimplementedAccountServer
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
func (s *accountHandler) Create(ctx context.Context, req *AccountCreateRequest) (*AccountCreateResponse, error) {
	res, err := s.accountUsecase.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("s.accountUsecase.Create: %s", err)
	}

	return res, nil
}
