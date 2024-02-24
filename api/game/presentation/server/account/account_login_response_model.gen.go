// Package account アカウントログインレスポンス
package account

func SetAccountLoginResponse(userAccount *UserAccount) *AccountLoginResponse {
	return &AccountLoginResponse{
		UserAccount: userAccount,
	}
}
