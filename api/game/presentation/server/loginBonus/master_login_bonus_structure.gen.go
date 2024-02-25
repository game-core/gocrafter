// Package loginBonus ログインボーナス
package loginBonus

func SetMasterLoginBonus(id int64, masterEventId int64, name string) *MasterLoginBonus {
	return &MasterLoginBonus{
		Id:            id,
		MasterEventId: masterEventId,
		Name:          name,
	}
}
