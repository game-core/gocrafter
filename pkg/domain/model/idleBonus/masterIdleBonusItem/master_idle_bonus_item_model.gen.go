// Package masterIdleBonusItem 放置ボーナスアイテム
package masterIdleBonusItem

type MasterIdleBonusItems []*MasterIdleBonusItem

type MasterIdleBonusItem struct {
	Id                        int64
	MasterIdleBonusScheduleId int64
	MasterItemId              int64
	Name                      string
	Count                     int32
}

func NewMasterIdleBonusItem() *MasterIdleBonusItem {
	return &MasterIdleBonusItem{}
}

func NewMasterIdleBonusItems() MasterIdleBonusItems {
	return MasterIdleBonusItems{}
}

func SetMasterIdleBonusItem(id int64, masterIdleBonusScheduleId int64, masterItemId int64, name string, count int32) *MasterIdleBonusItem {
	return &MasterIdleBonusItem{
		Id:                        id,
		MasterIdleBonusScheduleId: masterIdleBonusScheduleId,
		MasterItemId:              masterItemId,
		Name:                      name,
		Count:                     count,
	}
}
