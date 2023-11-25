package loginReward

type LoginRewardRewards []LoginRewardReward

type LoginRewardReward struct {
	ID int64 `json:"id"`

	ItemName string `json:"item_name"`

	StepNumber int `json:"step_number"`

	Name string `json:"name"`

	Count int `json:"count"`
}
