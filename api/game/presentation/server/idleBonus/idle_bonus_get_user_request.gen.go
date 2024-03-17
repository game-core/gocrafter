// Package idleBonus 放置ボーナスユーザー取得リクエスト
package idleBonus

func SetIdleBonusGetUserRequest(userId string) *IdleBonusGetUserRequest {
	return &IdleBonusGetUserRequest{
		UserId: userId,
	}
}
