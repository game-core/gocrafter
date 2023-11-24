package loginReward

import (
	"time"
)

type LoginRewardStatuses []LoginRewardStatus

type LoginRewardStatus struct {
	ID int64 `json:"id"`

	ShardKey string `json:"shard_key"`

	AccountID int64 `json:"account_id"`

	LoginRewardModelName string `json:"login_reward_model_Name"`

	LastReceivedAt time.Time `json:"last_received_at"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *LoginRewardStatus) TableName() string {
	return "login_reward_status"
}
