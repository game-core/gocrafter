// Package account アカウント作成リクエスト
package account

type AccountCreateRequests []*AccountCreateRequest

type AccountCreateRequest struct {
	UserId   string
	Name     string
	Password string
}

func NewAccountCreateRequest() *AccountCreateRequest {
	return &AccountCreateRequest{}
}

func NewAccountCreateRequests() AccountCreateRequests {
	return AccountCreateRequests{}
}

func SetAccountCreateRequest(userId string, name string, password string) *AccountCreateRequest {
	return &AccountCreateRequest{
		UserId:   userId,
		Name:     name,
		Password: password,
	}
}
