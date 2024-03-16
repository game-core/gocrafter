// Package idleBonus 放置ボーナスマスター取得リクエスト
package idleBonus

func SetIdleBonusGetMasterRequest(userId string, masterIdleBonusId int64) *IdleBonusGetMasterRequest {
	return &IdleBonusGetMasterRequest{
		UserId:            userId,
		MasterIdleBonusId: masterIdleBonusId,
	}
}
