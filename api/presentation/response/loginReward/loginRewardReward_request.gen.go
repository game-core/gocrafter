package loginReward

type LoginRewardRewards []LoginRewardReward

type LoginRewardReward struct {
	ID int64 `json:"id"`

	StepNumber int `json:"step_number"`

	Count int `json:"count"`

	Items Items `json:"items"`
}
