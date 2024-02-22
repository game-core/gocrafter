// Package account アカウント作成レスポンス
package account

import (
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
)

type AccountCreateResponses []*AccountCreateResponse

type AccountCreateResponse struct {
	UserAccount *userAccount.UserAccount
}

func NewAccountCreateResponse() *AccountCreateResponse {
	return &AccountCreateResponse{}
}

func NewAccountCreateResponses() AccountCreateResponses {
	return AccountCreateResponses{}
}

func SetAccountCreateResponse(userAccount *userAccount.UserAccount) *AccountCreateResponse {
	return &AccountCreateResponse{
		UserAccount: userAccount,
	}
}
