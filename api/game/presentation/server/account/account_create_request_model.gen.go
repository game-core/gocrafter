// Package account アカウント作成リクエスト
package account

func SetAccountCreateRequest(userId string, name string) *AccountCreateRequest {
	return &AccountCreateRequest{
		UserId: userId,
		Name:   name,
	}
}
