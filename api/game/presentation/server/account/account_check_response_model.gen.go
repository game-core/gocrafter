// Package account アカウント確認レスポンス
package account

func SetAccountCheckResponse(userAccount *UserAccount) *AccountCheckResponse {
	return &AccountCheckResponse{
		UserAccount: userAccount,
	}
}
