// Package account アカウント確認レスポンス
package account

import (
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
)

type AccountCheckResponses []*AccountCheckResponse

type AccountCheckResponse struct {
	UserAccount *userAccount.UserAccount
}

func NewAccountCheckResponse() *AccountCheckResponse {
	return &AccountCheckResponse{}
}

func NewAccountCheckResponses() AccountCheckResponses {
	return AccountCheckResponses{}
}

func SetAccountCheckResponse(userAccount *userAccount.UserAccount) *AccountCheckResponse {
	return &AccountCheckResponse{
		UserAccount: userAccount,
	}
}
