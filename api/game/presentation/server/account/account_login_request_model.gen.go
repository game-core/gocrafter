// Package account アカウントログインリクエスト
package account

func SetAccountLoginRequest(userId string, name string, password string) *AccountLoginRequest {
	return &AccountLoginRequest{
		UserId:   userId,
		Name:     name,
		Password: password,
	}
}
