package loginReward

import (
	"time"
)

type Events []Event

type Event struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	ResetHour int `json:"reset_hour"`

	RepeatSetting bool `json:"repeat_setting"`

	RepeatStartAt *time.Time `json:"repeat_start_at"`

	StartAt *time.Time `json:"start_at"`

	EndAt *time.Time `json:"end_at"`
}

func ToEvent(ID int64, Name string, ResetHour int, RepeatSetting bool, RepeatStartAt *time.Time, StartAt *time.Time, EndAt *time.Time) *Event {
	return &Event{
		ID:            ID,
		Name:          Name,
		ResetHour:     ResetHour,
		RepeatSetting: RepeatSetting,
		RepeatStartAt: RepeatStartAt,
		StartAt:       StartAt,
		EndAt:         EndAt,
	}
}
