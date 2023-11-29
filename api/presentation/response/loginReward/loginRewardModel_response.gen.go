package loginReward

type LoginRewardModels []LoginRewardModel

type LoginRewardModel struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Event Event `json:"event"`

	LoginRewardRewards LoginRewardRewards `json:"login_reward_rewards"`
}

func ToLoginRewardModel(ID int64, Name string, Event Event, LoginRewardRewards LoginRewardRewards) *LoginRewardModel {
	return &LoginRewardModel{
		ID:                 ID,
		Name:               Name,
		Event:              Event,
		LoginRewardRewards: LoginRewardRewards,
	}
}
