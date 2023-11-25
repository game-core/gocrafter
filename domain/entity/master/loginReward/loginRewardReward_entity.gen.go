package loginReward

import (
	"time"
)

type LoginRewardRewards []LoginRewardReward

type LoginRewardReward struct {
	ID int64 `json:"id"`

	LoginRewardModelName string `json:"login_reward_model_name"`

	Name string `json:"name"`

	StepNumber int `json:"step_number"`

	Items string `json:"items"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *LoginRewardReward) TableName() string {
	return "login_reward_reward"
}
