// Package masterLoginBonusItem ログインボーナスアイテム
package masterLoginBonusItem

type MasterLoginBonusItems []*MasterLoginBonusItem

type MasterLoginBonusItem struct {
	Id                         int64
	MasterLoginBonusScheduleId int64
	MasterItemId               int64
	Name                       string
	Count                      int32
}

func NewMasterLoginBonusItem() *MasterLoginBonusItem {
	return &MasterLoginBonusItem{}
}

func NewMasterLoginBonusItems() MasterLoginBonusItems {
	return MasterLoginBonusItems{}
}

func SetMasterLoginBonusItem(id int64, masterLoginBonusScheduleId int64, masterItemId int64, name string, count int32) *MasterLoginBonusItem {
	return &MasterLoginBonusItem{
		Id:                         id,
		MasterLoginBonusScheduleId: masterLoginBonusScheduleId,
		MasterItemId:               masterItemId,
		Name:                       name,
		Count:                      count,
	}
}

func (t *MasterLoginBonusItem) TableName() string {
	return "master_login_bonus_item"
}
