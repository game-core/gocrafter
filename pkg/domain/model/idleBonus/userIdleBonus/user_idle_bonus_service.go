package userIdleBonus

import (
	"time"
)

// GetReceivedAt 最終受け取り日時を取得を取得（nilの場合は現在日時を返す）
func (s *UserIdleBonus) GetReceivedAt(now time.Time) time.Time {
	var receivedAt time.Time
	if s != nil {
		receivedAt = s.ReceivedAt
	} else {
		receivedAt = now
	}

	return receivedAt
}
