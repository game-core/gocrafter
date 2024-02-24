// Package account アカウントログインレスポンス
package account

import (
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
)

type AccountLoginResponses []*AccountLoginResponse

type AccountLoginResponse struct {
	Token       string
	UserAccount *userAccount.UserAccount
}

func NewAccountLoginResponse() *AccountLoginResponse {
	return &AccountLoginResponse{}
}

func NewAccountLoginResponses() AccountLoginResponses {
	return AccountLoginResponses{}
}

func SetAccountLoginResponse(token string, userAccount *userAccount.UserAccount) *AccountLoginResponse {
	return &AccountLoginResponse{
		Token:       token,
		UserAccount: userAccount,
	}
}
