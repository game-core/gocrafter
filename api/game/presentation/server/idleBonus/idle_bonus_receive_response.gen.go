// Package idleBonus 放置ボーナス受け取りレスポンス
package idleBonus

func SetIdleBonusReceiveResponse(userIdleBonus *UserIdleBonus, masterIdleBonus *MasterIdleBonus, masterIdleBonusEvent *MasterIdleBonusEvent, masterIdleBonusItems []*MasterIdleBonusItem, masterIdleBonusSchedules []*MasterIdleBonusSchedule) *IdleBonusReceiveResponse {
	return &IdleBonusReceiveResponse{
		UserIdleBonus:            userIdleBonus,
		MasterIdleBonus:          masterIdleBonus,
		MasterIdleBonusEvent:     masterIdleBonusEvent,
		MasterIdleBonusItems:     masterIdleBonusItems,
		MasterIdleBonusSchedules: masterIdleBonusSchedules,
	}
}
