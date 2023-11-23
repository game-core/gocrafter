package loginReward

import (
	"time"
)

// HasReceived 報酬を受け取っているか
func (e *LoginRewardStatus) HasReceived(now time.Time, resetHour int) bool {
	return e.LastReceivedAt.After(time.Date(now.Year(), now.Month(), now.Day(), resetHour, 0, 0, 0, now.Location()))
}
