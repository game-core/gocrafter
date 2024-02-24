// Package account アカウント確認リクエスト
package account

func SetAccountCheckRequest(userId string) *AccountCheckRequest {
	return &AccountCheckRequest{
		UserId: userId,
	}
}
