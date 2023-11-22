package loginReward

type ReceiveLoginRewards []ReceiveLoginReward

type ReceiveLoginReward struct {
	Status int64 `json:"status"`

	Item LoginRewardStatus `json:"item"`
}
