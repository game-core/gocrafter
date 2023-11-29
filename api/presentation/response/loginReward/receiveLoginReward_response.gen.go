package loginReward

type ReceiveLoginRewards []ReceiveLoginReward

type ReceiveLoginReward struct {
	Status int64 `json:"status"`

	Item LoginRewardStatus `json:"item"`
}

func ToReceiveLoginReward(Status int64, Item LoginRewardStatus) *ReceiveLoginReward {
	return &ReceiveLoginReward{
		Status: Status,
		Item:   Item,
	}
}
