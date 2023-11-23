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
