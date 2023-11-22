package loginReward

import (
	"time"
)

type LoginRewardStatuses []LoginRewardStatus

type LoginRewardStatus struct {
	ID int64 `json:"id"`

	LoginRewardModel LoginRewardModel `json:"login_reward_model"`

	LoginRewardReceiveSteps LoginRewardReceiveSteps `json:"login_reward_receive_steps"`

	Item Item `json:"item"`

	LastReceivedAt time.Time `json:"created_at"`
}