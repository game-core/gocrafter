package times

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimeToPb(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(*t)
}
