package userAction

import (
	"time"
)

// CheckExpiration 有効期限を確認する
func (s *UserAction) CheckExpiration(now time.Time, expiration *int32) bool {
	if expiration == nil {
		return true
	}

	if float64(*expiration) < now.Sub(s.StartedAt).Hours() {
		return false
	}

	return true
}
