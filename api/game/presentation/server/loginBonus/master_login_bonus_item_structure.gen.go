// Package loginBonus ログインボーナスアイテム
package loginBonus

func SetMasterLoginBonusItem(id int64, masterLoginBonusScheduleId int64, masterItemId int64, name string, count int32) *MasterLoginBonusItem {
	return &MasterLoginBonusItem{
		Id:                         id,
		MasterLoginBonusScheduleId: masterLoginBonusScheduleId,
		MasterItemId:               masterItemId,
		Name:                       name,
		Count:                      count,
	}
}
