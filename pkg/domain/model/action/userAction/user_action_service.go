package userAction

import (
	"time"
)

// CheckExpiration 有効期限を確認する
func (s *UserAction) CheckExpiration(expiration *int32, now time.Time) bool {
	if expiration == nil {
		return true
	}

	if float64(*expiration) < now.Sub(s.StartedAt).Hours() {
		return false
	}

	return true
}
