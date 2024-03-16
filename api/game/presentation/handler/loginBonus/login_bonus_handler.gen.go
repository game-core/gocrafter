package loginBonus

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/loginBonus"
	loginBonusUsecase "github.com/game-core/gocrafter/api/game/usecase/loginBonus"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/tokens"
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

// GetMaster ログインボーナスマスターを取得する
func (s *loginBonusHandler) GetMaster(ctx context.Context, req *loginBonus.LoginBonusGetMasterRequest) (*loginBonus.LoginBonusGetMasterResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.loginBonusUsecase.GetMaster(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.loginBonusUsecase.GetMaster", err)
	}

	return res, nil
}

// Receive ログインボーナスを受け取る
func (s *loginBonusHandler) Receive(ctx context.Context, req *loginBonus.LoginBonusReceiveRequest) (*loginBonus.LoginBonusReceiveResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.loginBonusUsecase.Receive(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.loginBonusUsecase.Receive", err)
	}

	return res, nil
}
