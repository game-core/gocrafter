package loginReward

type ReceiveLoginRewards []ReceiveLoginReward

type ReceiveLoginReward struct {
	Status int64 `json:"status"`

	LoginRewardStatus LoginRewardStatus `json:"login_reward_status"`
}

func ToReceiveLoginReward(Status int64, LoginRewardStatus LoginRewardStatus) *ReceiveLoginReward {
	return &ReceiveLoginReward{
		Status:            Status,
		LoginRewardStatus: LoginRewardStatus,
	}
}
