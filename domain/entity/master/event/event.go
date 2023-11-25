package event

import (
	"github.com/game-core/gocrafter/config/times"
	"time"
)

// GetDayCount イベントの経過日数を取得
func (e *Event) GetDayCount(now time.Time) int {
	if e.RepeatSetting {
		return times.GetDayCount(*e.RepeatStartAt, now)
	}

	return times.GetDayCount(*e.StartAt, now)
}

// IsEventPeriod イベント期間中か
func (e *Event) IsEventPeriod(now time.Time) bool {
	if e.StartAt != nil && e.StartAt.After(now) {
		return false
	}

	// イベント終了後の場合
	if e.EndAt != nil && e.EndAt.Before(now) {
		return false
	}

	// 定常イベント開始前の場合
	if e.RepeatSetting && e.RepeatStartAt.After(now) {
		return false
	}

	return true
}
