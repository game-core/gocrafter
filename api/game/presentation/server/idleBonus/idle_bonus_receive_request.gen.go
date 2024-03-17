// Package idleBonus 放置ボーナス受け取りリクエスト
package idleBonus

func SetIdleBonusReceiveRequest(userId string, masterIdleBonusId int64) *IdleBonusReceiveRequest {
	return &IdleBonusReceiveRequest{
		UserId:            userId,
		MasterIdleBonusId: masterIdleBonusId,
	}
}
