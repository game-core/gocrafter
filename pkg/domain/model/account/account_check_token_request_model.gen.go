// Package account アカウントトークン確認リクエスト
package account

type AccountCheckTokenRequests []*AccountCheckTokenRequest

type AccountCheckTokenRequest struct {
	UserId string
}

func NewAccountCheckTokenRequest() *AccountCheckTokenRequest {
	return &AccountCheckTokenRequest{}
}

func NewAccountCheckTokenRequests() AccountCheckTokenRequests {
	return AccountCheckTokenRequests{}
}

func SetAccountCheckTokenRequest(userId string) *AccountCheckTokenRequest {
	return &AccountCheckTokenRequest{
		UserId: userId,
	}
}
