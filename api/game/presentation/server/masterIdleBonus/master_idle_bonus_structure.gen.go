// Package masterIdleBonus 放置ボーナス
package masterIdleBonus

func SetMasterIdleBonus(id int64, masterIdleBonusEventId int64, name string) *MasterIdleBonus {
	return &MasterIdleBonus{
		Id:                     id,
		MasterIdleBonusEventId: masterIdleBonusEventId,
		Name:                   name,
	}
}
