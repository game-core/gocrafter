// Package masterLoginBonus ログインボーナス
package masterLoginBonus

type MasterLoginBonuses []*MasterLoginBonus

type MasterLoginBonus struct {
	Id                      int64
	MasterLoginBonusEventId int64
	Name                    string
}

func NewMasterLoginBonus() *MasterLoginBonus {
	return &MasterLoginBonus{}
}

func NewMasterLoginBonuses() MasterLoginBonuses {
	return MasterLoginBonuses{}
}

func SetMasterLoginBonus(id int64, masterLoginBonusEventId int64, name string) *MasterLoginBonus {
	return &MasterLoginBonus{
		Id:                      id,
		MasterLoginBonusEventId: masterLoginBonusEventId,
		Name:                    name,
	}
}

func (t *MasterLoginBonus) TableName() string {
	return "master_login_bonus"
}
