// Package idleBonus 放置ボーナスマスター取得レスポンス
package idleBonus

func SetIdleBonusGetMasterResponse(masterIdleBonus *MasterIdleBonus, masterIdleBonusEvent *MasterIdleBonusEvent, masterIdleBonusItems []*MasterIdleBonusItem, masterIdleBonusSchedules []*MasterIdleBonusSchedule) *IdleBonusGetMasterResponse {
	return &IdleBonusGetMasterResponse{
		MasterIdleBonus:          masterIdleBonus,
		MasterIdleBonusEvent:     masterIdleBonusEvent,
		MasterIdleBonusItems:     masterIdleBonusItems,
		MasterIdleBonusSchedules: masterIdleBonusSchedules,
	}
}
