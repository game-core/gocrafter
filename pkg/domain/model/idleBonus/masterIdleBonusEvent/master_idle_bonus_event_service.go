package masterIdleBonusEvent

import (
	"time"
)

// CheckEventPeriod イベント期間を確認する
func (s *MasterIdleBonusEvent) CheckEventPeriod(now time.Time) bool {
	if s.StartAt.After(now) {
		return false
	}

	if s.EndAt != nil && s.EndAt.Before(now) {
		return false
	}

	return true
}
