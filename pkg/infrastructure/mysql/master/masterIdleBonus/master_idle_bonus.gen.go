// Package masterIdleBonus 放置ボーナス
package masterIdleBonus

type MasterIdleBonuses []*MasterIdleBonus

type MasterIdleBonus struct {
	Id                     int64
	MasterIdleBonusEventId int64
	Name                   string
}

func NewMasterIdleBonus() *MasterIdleBonus {
	return &MasterIdleBonus{}
}

func NewMasterIdleBonuses() MasterIdleBonuses {
	return MasterIdleBonuses{}
}

func SetMasterIdleBonus(id int64, masterIdleBonusEventId int64, name string) *MasterIdleBonus {
	return &MasterIdleBonus{
		Id:                     id,
		MasterIdleBonusEventId: masterIdleBonusEventId,
		Name:                   name,
	}
}

func (t *MasterIdleBonus) TableName() string {
	return "master_idle_bonus"
}
