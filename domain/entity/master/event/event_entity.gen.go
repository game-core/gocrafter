package event

import (
	"time"
)

type Events []Event

type Event struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	ResetHour *int `json:"repeat_hour"`

	Repeat bool `json:"repeat"`

	RepeatStartAt *time.Time `json:"repeat_start_at"`

	StartAt *time.Time `json:"start_at"`

	EndAt *time.Time `json:"end_at"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
