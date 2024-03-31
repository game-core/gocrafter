// Package account アカウントトークン確認レスポンス
package account

import (
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccountToken"
)

type AccountCheckTokenResponses []*AccountCheckTokenResponse

type AccountCheckTokenResponse struct {
	UserAccountToken *userAccountToken.UserAccountToken
}

func NewAccountCheckTokenResponse() *AccountCheckTokenResponse {
	return &AccountCheckTokenResponse{}
}

func NewAccountCheckTokenResponses() AccountCheckTokenResponses {
	return AccountCheckTokenResponses{}
}

func SetAccountCheckTokenResponse(userAccountToken *userAccountToken.UserAccountToken) *AccountCheckTokenResponse {
	return &AccountCheckTokenResponse{
		UserAccountToken: userAccountToken,
	}
}
