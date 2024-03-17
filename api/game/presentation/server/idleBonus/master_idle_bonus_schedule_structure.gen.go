// Package idleBonus 放置ボーナススケジュール
package idleBonus

func SetMasterIdleBonusSchedule(id int64, masterIdleBonusId int64, step int32, name string) *MasterIdleBonusSchedule {
	return &MasterIdleBonusSchedule{
		Id:                id,
		MasterIdleBonusId: masterIdleBonusId,
		Step:              step,
		Name:              name,
	}
}
