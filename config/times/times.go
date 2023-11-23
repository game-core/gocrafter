package times

import "time"

// GetDayCount 経過日数を取得
func GetDayCount(from time.Time, to time.Time) int {
	if from.After(to) {
		return 0
	}

	return int(to.Sub(from).Hours() / 24)
}
