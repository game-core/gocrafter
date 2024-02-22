// Package account アカウント作成リクエスト
package account

type AccountCreateRequests []*AccountCreateRequest

type AccountCreateRequest struct {
	Name string
}

func NewAccountCreateRequest() *AccountCreateRequest {
	return &AccountCreateRequest{}
}

func NewAccountCreateRequests() AccountCreateRequests {
	return AccountCreateRequests{}
}

func SetAccountCreateRequest(name string) *AccountCreateRequest {
	return &AccountCreateRequest{
		Name: name,
	}
}
