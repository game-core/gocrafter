// Package idleBonus 放置ボーナスユーザー取得レスポンス
package idleBonus

func SetIdleBonusGetUserResponse(userIdleBonuses []*UserIdleBonus) *IdleBonusGetUserResponse {
	return &IdleBonusGetUserResponse{
		UserIdleBonuses: userIdleBonuses,
	}
}
