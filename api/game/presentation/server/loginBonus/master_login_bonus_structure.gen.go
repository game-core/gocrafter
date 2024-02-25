// Package loginBonus ログインボーナス
package loginBonus

func SetMasterLoginBonus(id int64, masterLoginBonusEventId int64, name string) *MasterLoginBonus {
	return &MasterLoginBonus{
		Id:                      id,
		MasterLoginBonusEventId: masterLoginBonusEventId,
		Name:                    name,
	}
}
