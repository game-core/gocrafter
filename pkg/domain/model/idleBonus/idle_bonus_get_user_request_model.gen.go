// Package idleBonus 放置ボーナスユーザー取得リクエスト
package idleBonus

type IdleBonusGetUserRequests []*IdleBonusGetUserRequest

type IdleBonusGetUserRequest struct {
	UserId string
}

func NewIdleBonusGetUserRequest() *IdleBonusGetUserRequest {
	return &IdleBonusGetUserRequest{}
}

func NewIdleBonusGetUserRequests() IdleBonusGetUserRequests {
	return IdleBonusGetUserRequests{}
}

func SetIdleBonusGetUserRequest(userId string) *IdleBonusGetUserRequest {
	return &IdleBonusGetUserRequest{
		UserId: userId,
	}
}
