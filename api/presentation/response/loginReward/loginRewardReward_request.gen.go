package loginReward

type LoginRewardRewards []LoginRewardReward

type LoginRewardReward struct {
	ID int64 `json:"id"`

	ItemName string `json:"item_name"`

	Name string `json:"name"`

	StepNumber int `json:"step_number"`
}
