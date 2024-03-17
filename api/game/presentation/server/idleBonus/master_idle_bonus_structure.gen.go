// Package idleBonus 放置ボーナス
package idleBonus

func SetMasterIdleBonus(id int64, masterIdleBonusEventId int64, name string) *MasterIdleBonus {
	return &MasterIdleBonus{
		Id:                     id,
		MasterIdleBonusEventId: masterIdleBonusEventId,
		Name:                   name,
	}
}
