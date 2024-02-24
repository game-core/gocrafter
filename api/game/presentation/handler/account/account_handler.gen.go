package account

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/account"
	accountUsecase "github.com/game-core/gocrafter/api/game/usecase/account"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/tokens"
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

// Login アカウントをログインする
func (s *accountHandler) Login(ctx context.Context, req *account.AccountLoginRequest) (*account.AccountLoginResponse, error) {
	res, err := s.accountUsecase.Login(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.accountUsecase.Login", err)
	}

	return res, nil
}

// Check アカウントを確認する
func (s *accountHandler) Check(ctx context.Context, req *account.AccountCheckRequest) (*account.AccountCheckResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.accountUsecase.Check(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.accountUsecase.Check", err)
	}

	return res, nil
}
