// Package masterIdleBonusSchedule 放置ボーナススケジュール
package masterIdleBonusSchedule

type MasterIdleBonusSchedules []*MasterIdleBonusSchedule

type MasterIdleBonusSchedule struct {
	Id                int64
	MasterIdleBonusId int64
	Step              int32
	Name              string
}

func NewMasterIdleBonusSchedule() *MasterIdleBonusSchedule {
	return &MasterIdleBonusSchedule{}
}

func NewMasterIdleBonusSchedules() MasterIdleBonusSchedules {
	return MasterIdleBonusSchedules{}
}

func SetMasterIdleBonusSchedule(id int64, masterIdleBonusId int64, step int32, name string) *MasterIdleBonusSchedule {
	return &MasterIdleBonusSchedule{
		Id:                id,
		MasterIdleBonusId: masterIdleBonusId,
		Step:              step,
		Name:              name,
	}
}
