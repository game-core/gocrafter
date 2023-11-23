package loginReward

import (
	"time"
)

type LoginRewardRewards []LoginRewardReward

type LoginRewardReward struct {
	ID int64 `json:"id"`

	LoginRewardID int64 `json:"login_reward_id"`

	ItemID int64 `json:"item_id"`

	Name string `json:"name"`

	StepNumber int `json:"step_number"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
