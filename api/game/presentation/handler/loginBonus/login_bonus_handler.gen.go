package loginBonus

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/loginBonus"
	loginBonusUsecase "github.com/game-core/gocrafter/api/game/usecase/loginBonus"
	"github.com/game-core/gocrafter/internal/errors"
)

type LoginBonusHandler interface {
	loginBonus.LoginBonusServer
}

type loginBonusHandler struct {
	loginBonus.UnimplementedLoginBonusServer
	loginBonusUsecase loginBonusUsecase.LoginBonusUsecase
}

func NewLoginBonusHandler(
	loginBonusUsecase loginBonusUsecase.LoginBonusUsecase,
) LoginBonusHandler {
	return &loginBonusHandler{
		loginBonusUsecase: loginBonusUsecase,
	}
}

// Receive ログインボーナスを受け取る
func (s *loginBonusHandler) Receive(ctx context.Context, req *loginBonus.LoginBonusReceiveRequest) (*loginBonus.LoginBonusReceiveResponse, error) {
	res, err := s.loginBonusUsecase.Receive(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.loginBonusUsecase.Receive", err)
	}

	return res, nil
}
