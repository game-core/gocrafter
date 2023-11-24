package loginReward

import (
	"time"
)

// HasReceived 報酬を受け取っているか
func (e *LoginRewardStatus) HasReceived(now time.Time, resetHour int) bool {
	resetTime := time.Date(now.Year(), now.Month(), now.Day(), resetHour, 0, 0, 0, now.Location())
	if now.Before(resetTime) {
		return e.LastReceivedAt.Add(24 * time.Hour).Before(now)
	}

	return e.LastReceivedAt.Before(resetTime)
}
