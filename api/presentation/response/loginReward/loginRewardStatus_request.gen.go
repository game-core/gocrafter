package loginReward

import (
	"time"
)

type LoginRewardStatuses []LoginRewardStatus

type LoginRewardStatus struct {
	ID int64 `json:"id"`

	LoginRewardModel LoginRewardModel `json:"login_reward_model"`

	Item Item `json:"item"`

	LastReceivedAt time.Time `json:"last_received_at"`
}
