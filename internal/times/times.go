package times

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Now 現在時刻
func Now() time.Time {
	return time.Now()
}

// TimeToPb timeをtimestamppbに変換
func TimeToPb(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(*t)
}
