package event

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

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *Event) TableName() string {
	return "event"
}
