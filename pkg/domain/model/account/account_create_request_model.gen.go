// Package account アカウント作成リクエスト
package account

type AccountCreateRequests []*AccountCreateRequest

type AccountCreateRequest struct {
	UserId string
	Name   string
}

func NewAccountCreateRequest() *AccountCreateRequest {
	return &AccountCreateRequest{}
}

func NewAccountCreateRequests() AccountCreateRequests {
	return AccountCreateRequests{}
}

func SetAccountCreateRequest(userId string, name string) *AccountCreateRequest {
	return &AccountCreateRequest{
		UserId: userId,
		Name:   name,
	}
}
