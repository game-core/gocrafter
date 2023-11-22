package loginReward

import (
	"time"
)

type Events []Event

type Event struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Repeat bool `json:"repeat"`

	StartAt *time.Time `json:"start_at"`

	EndAt *time.Time `json:"end_at"`
}
