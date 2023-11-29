package loginReward

type LoginRewardRewards []LoginRewardReward

type LoginRewardReward struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	StepNumber int `json:"step_number"`

	Items Items `json:"items"`
}

func ToLoginRewardReward(ID int64, Name string, StepNumber int, Items Items) *LoginRewardReward {
	return &LoginRewardReward{
		ID:         ID,
		Name:       Name,
		StepNumber: StepNumber,
		Items:      Items,
	}
}
