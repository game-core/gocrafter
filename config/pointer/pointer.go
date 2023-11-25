package pointer

import (
	"time"
)

func PointerToString(str *string) string {
	return *str
}

func StringToPointer(str string) *string {
	return &str
}

func TimeToPointer(t time.Time) *time.Time {
	return &t
}

func PointerToTime(pt *time.Time) time.Time {
	return *pt
}
