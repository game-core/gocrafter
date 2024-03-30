package masterRankingEvent

import (
	"time"
)

// CheckEventPeriod イベント期間を確認する
func (s *MasterRankingEvent) CheckEventPeriod(now time.Time) bool {
	if s.StartAt.After(now) {
		return false
	}

	if s.EndAt != nil && s.EndAt.Before(now) {
		return false
	}

	return true
}

// GetLastEventAt イベントの最終更新日時を取得する
func (s *MasterRankingEvent) GetLastEventAt(now time.Time) time.Time {
	lastUpdateDate := s.StartAt
	numIntervals := int(now.Sub(s.StartAt).Hours() / float64(s.IntervalHour))
	lastUpdateDate = lastUpdateDate.Add(time.Duration(numIntervals * int(time.Hour) * int(s.IntervalHour)))

	if lastUpdateDate.After(now) {
		lastUpdateDate = s.StartAt
	}

	return lastUpdateDate
}
