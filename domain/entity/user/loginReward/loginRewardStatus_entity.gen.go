package loginReward

import (
	"time"
)

type LoginRewardStatuses []LoginRewardStatus

type LoginRewardStatus struct {
	ID int64 `json:"id"`

	ShardKey int `json:"shard_key"`

	AccountID int64 `json:"user_id"`

	LoginRewardModelID int64 `json:"login_reward_model_id"`

	LastReceivedAt time.Time `json:"created_at"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
