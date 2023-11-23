package loginReward

import (
	"time"
)

type LoginRewardRewards []LoginRewardReward

type LoginRewardReward struct {
	ID int64 `json:"id"`

	LoginRewardModelName string `json:"login_reward_model_name"`

	ItemName string `json:"item_name"`

	Name string `json:"name"`

	StepNumber int `json:"step_number"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
