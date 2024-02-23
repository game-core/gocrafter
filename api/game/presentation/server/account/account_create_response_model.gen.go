// Package account アカウント作成レスポンス
package account

func SetAccountCreateResponse(userAccount *UserAccount) *AccountCreateResponse {
	return &AccountCreateResponse{
		UserAccount: userAccount,
	}
}
