// Package account アカウント作成リクエスト
package account

func SetAccountCreateRequest(name string) *AccountCreateRequest {
	return &AccountCreateRequest{
		Name: name,
	}
}
