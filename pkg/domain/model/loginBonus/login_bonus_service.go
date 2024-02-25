package loginBonus

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/userLoginBonus"
)

type LoginBonusService interface {
	Receive(ctx context.Context, tx *gorm.DB, req *LoginBonusReceiveRequest) (*LoginBonusReceiveResponse, error)
}

type loginBonusService struct {
	userLoginBonusRepository userLoginBonus.UserLoginBonusRepository
}

func NewLoginBonusService(
	userLoginBonusRepository userLoginBonus.UserLoginBonusRepository,
) LoginBonusService {
	return &loginBonusService{
		userLoginBonusRepository: userLoginBonusRepository,
	}
}

// Receive ログインボーナスを受け取る
func (s *loginBonusService) Receive(ctx context.Context, tx *gorm.DB, req *LoginBonusReceiveRequest) (*LoginBonusReceiveResponse, error) {
	return nil, nil
}
