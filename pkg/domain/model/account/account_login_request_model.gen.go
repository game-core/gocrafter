// Package account アカウントログインリクエスト
package account

type AccountLoginRequests []*AccountLoginRequest

type AccountLoginRequest struct {
	UserId   string
	Name     string
	Password string
}

func NewAccountLoginRequest() *AccountLoginRequest {
	return &AccountLoginRequest{}
}

func NewAccountLoginRequests() AccountLoginRequests {
	return AccountLoginRequests{}
}

func SetAccountLoginRequest(userId string, name string, password string) *AccountLoginRequest {
	return &AccountLoginRequest{
		UserId:   userId,
		Name:     name,
		Password: password,
	}
}
