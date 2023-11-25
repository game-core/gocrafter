package loginReward

import (
	"time"
)

type Events []Event

type Event struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	ResetHour int `json:"repeat_hour"`

	RepeatSetting bool `json:"repeat_setting"`

	RepeatStartAt *time.Time `json:"repeat_start_at"`

	StartAt *time.Time `json:"start_at"`

	EndAt *time.Time `json:"end_at"`
}
