// Package masterLoginBonusSchedule ログインボーナススケジュール
package masterLoginBonusSchedule

type MasterLoginBonusSchedules []*MasterLoginBonusSchedule

type MasterLoginBonusSchedule struct {
	Id                 int64
	MasterLoginBonusId int64
	Step               int32
	Name               string
}

func NewMasterLoginBonusSchedule() *MasterLoginBonusSchedule {
	return &MasterLoginBonusSchedule{}
}

func NewMasterLoginBonusSchedules() MasterLoginBonusSchedules {
	return MasterLoginBonusSchedules{}
}

func SetMasterLoginBonusSchedule(id int64, masterLoginBonusId int64, step int32, name string) *MasterLoginBonusSchedule {
	return &MasterLoginBonusSchedule{
		Id:                 id,
		MasterLoginBonusId: masterLoginBonusId,
		Step:               step,
		Name:               name,
	}
}
