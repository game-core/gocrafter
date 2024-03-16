// Package loginBonus ログインボーナスユーザー取得レスポンス
package loginBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/userLoginBonus"
)

type LoginBonusGetUserResponses []*LoginBonusGetUserResponse

type LoginBonusGetUserResponse struct {
	UserLoginBonuses userLoginBonus.UserLoginBonuses
}

func NewLoginBonusGetUserResponse() *LoginBonusGetUserResponse {
	return &LoginBonusGetUserResponse{}
}

func NewLoginBonusGetUserResponses() LoginBonusGetUserResponses {
	return LoginBonusGetUserResponses{}
}

func SetLoginBonusGetUserResponse(userLoginBonuses userLoginBonus.UserLoginBonuses) *LoginBonusGetUserResponse {
	return &LoginBonusGetUserResponse{
		UserLoginBonuses: userLoginBonuses,
	}
}
