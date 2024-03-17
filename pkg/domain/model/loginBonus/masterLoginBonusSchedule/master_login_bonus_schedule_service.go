package masterLoginBonusSchedule

import "time"

// GetScheduleByStep ステップからスケジュールを取得する
func (s *MasterLoginBonusSchedules) GetScheduleByStep(step int32) *MasterLoginBonusSchedule {
	masterLoginBonusSchedule := &MasterLoginBonusSchedule{}
	for _, mlbst := range *s {
		if mlbst.Step == step {
			masterLoginBonusSchedule = mlbst
			break
		}
	}

	return masterLoginBonusSchedule
}

// GetStep ステップを取得
func (s *MasterLoginBonusSchedules) GetStep(intervalHour int32, startAt, now time.Time) int32 {
	maxStep := int32(len(*s) - 1)
	intervalStep := int32(now.Sub(startAt).Hours() / float64(intervalHour))
	if intervalStep <= maxStep {
		return intervalStep
	}

	return intervalStep % (maxStep + 1)
}
