package loginReward

import (
	"time"
)

type LoginRewardReceiveSteps []LoginRewardReceiveStep

type LoginRewardReceiveStep struct {
	ID int64 `json:"id"`

	ShardKey int `json:"shard_key"`

	AccountID int64 `json:"user_id"`

	LoginRewardStatusID int64 `json:"login_reward_model_id"`

	StepNumber int `json:"step_number"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
