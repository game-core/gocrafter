// Package account アカウント作成リクエスト
package account

func SetAccountCreateRequest(name string, password string) *AccountCreateRequest {
	return &AccountCreateRequest{
		Name:     name,
		Password: password,
	}
}
