package masterIdleBonusSchedule

import (
	"time"
)

// GetSchedulesByStep ステップからスケジュール一覧を取得する
func (s *MasterIdleBonusSchedules) GetSchedulesByStep(step int32) MasterIdleBonusSchedules {
	masterIdleBonusSchedules := NewMasterIdleBonusSchedules()
	for _, mlbst := range *s {
		if step >= mlbst.Step {
			masterIdleBonusSchedules = append(masterIdleBonusSchedules, mlbst)
		}
	}

	return masterIdleBonusSchedules
}

// GetStep ステップを取得
func (s *MasterIdleBonusSchedules) GetStep(intervalHour int32, receivedAt, now time.Time) int32 {
	maxStep := int32(len(*s) - 1)
	intervalStep := int32(now.Sub(receivedAt).Hours() / float64(intervalHour))
	if intervalStep >= maxStep {
		return maxStep
	}

	return intervalStep
}
