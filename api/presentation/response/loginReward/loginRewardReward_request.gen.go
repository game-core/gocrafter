package loginReward

type LoginRewardRewards []LoginRewardReward

type LoginRewardReward struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	StepNumber int `json:"step_number"`

	Items Items `json:"items"`
}
