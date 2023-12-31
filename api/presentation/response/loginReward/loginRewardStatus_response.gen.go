package loginReward

import (
	"time"
)

type LoginRewardStatuses []LoginRewardStatus

type LoginRewardStatus struct {
	ID int64 `json:"id"`

	LoginRewardModel LoginRewardModel `json:"login_reward_model"`

	Items Items `json:"items"`

	LastReceivedAt time.Time `json:"last_received_at"`
}

func ToLoginRewardStatus(ID int64, LoginRewardModel LoginRewardModel, Items Items, LastReceivedAt time.Time) *LoginRewardStatus {
	return &LoginRewardStatus{
		ID:               ID,
		LoginRewardModel: LoginRewardModel,
		Items:            Items,
		LastReceivedAt:   LastReceivedAt,
	}
}
