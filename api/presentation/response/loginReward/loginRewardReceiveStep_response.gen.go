package loginReward

type LoginRewardReceiveSteps []LoginRewardReceiveStep

type LoginRewardReceiveStep struct {
	ID int64 `json:"id"`

	StepNumber int `json:"step_number"`
}

func ToLoginRewardReceiveStep(ID int64, StepNumber int) *LoginRewardReceiveStep {
	return &LoginRewardReceiveStep{
		ID:         ID,
		StepNumber: StepNumber,
	}
}
