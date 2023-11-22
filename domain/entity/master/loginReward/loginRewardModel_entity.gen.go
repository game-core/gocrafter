package loginReward

import (
	"time"
)

type LoginRewardModels []LoginRewardModel

type LoginRewardModel struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	EventID int64 `json:"event_id"`

	Rewards string `json:"rewards"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
