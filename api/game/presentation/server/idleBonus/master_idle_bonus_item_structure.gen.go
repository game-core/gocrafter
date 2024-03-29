// Package idleBonus 放置ボーナスアイテム
package idleBonus

func SetMasterIdleBonusItem(id int64, masterIdleBonusScheduleId int64, masterItemId int64, name string, count int32) *MasterIdleBonusItem {
	return &MasterIdleBonusItem{
		Id:                        id,
		MasterIdleBonusScheduleId: masterIdleBonusScheduleId,
		MasterItemId:              masterItemId,
		Name:                      name,
		Count:                     count,
	}
}
