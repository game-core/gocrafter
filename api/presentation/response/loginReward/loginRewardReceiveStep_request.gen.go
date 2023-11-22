package loginReward

type LoginRewardReceiveSteps []LoginRewardReceiveStep

type LoginRewardReceiveStep struct {
	ID int64 `json:"id"`

	StepNumber int `json:"step_number"`
}
