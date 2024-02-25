package userLoginBonus

import "time"

// CheckReceived 報酬を受け取っているか確認する
func (s *UserLoginBonus) CheckReceived(resetHour int32, now time.Time) bool {
	resetTime := time.Date(now.Year(), now.Month(), now.Day(), int(resetHour), 0, 0, 0, now.Location())
	if now.Before(resetTime) {
		return !s.ReceivedAt.Add(24 * time.Hour).Before(now)
	}

	return !s.ReceivedAt.Before(resetTime)
}
