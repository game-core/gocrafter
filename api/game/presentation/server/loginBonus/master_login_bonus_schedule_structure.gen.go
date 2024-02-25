// Package loginBonus ログインボーナススケジュール
package loginBonus

func SetMasterLoginBonusSchedule(id int64, masterLoginBonusId int64, step int32, name string) *MasterLoginBonusSchedule {
	return &MasterLoginBonusSchedule{
		Id:                 id,
		MasterLoginBonusId: masterLoginBonusId,
		Step:               step,
		Name:               name,
	}
}
