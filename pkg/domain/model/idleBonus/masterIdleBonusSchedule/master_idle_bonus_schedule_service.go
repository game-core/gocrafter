package masterIdleBonusSchedule

import (
	"time"

	"github.com/game-core/gocrafter/internal/errors"
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
func (s *MasterIdleBonusSchedules) GetStep(intervalHour int32, receivedAt, now time.Time) (int32, error) {
	maxStep := int32(len(*s) - 1)
	intervalStep := int32(now.Sub(receivedAt).Hours() / float64(intervalHour))

	if intervalStep <= 0 {
		return 0, errors.NewError("already received")
	}

	if intervalStep > maxStep {
		return maxStep, nil
	}

	return intervalStep - 1, nil
}
