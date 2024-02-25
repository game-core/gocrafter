// Package account アカウントログインレスポンス
package account

func SetAccountLoginResponse(token string, userAccount *UserAccount) *AccountLoginResponse {
	return &AccountLoginResponse{
		Token:       token,
		UserAccount: userAccount,
	}
}
