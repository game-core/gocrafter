// Package account アカウント確認リクエスト
package account

type AccountCheckRequests []*AccountCheckRequest

type AccountCheckRequest struct {
	UserId string
}

func NewAccountCheckRequest() *AccountCheckRequest {
	return &AccountCheckRequest{}
}

func NewAccountCheckRequests() AccountCheckRequests {
	return AccountCheckRequests{}
}

func SetAccountCheckRequest(userId string) *AccountCheckRequest {
	return &AccountCheckRequest{
		UserId: userId,
	}
}
