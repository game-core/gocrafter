package loginReward

type GetLoginRewardModels []GetLoginRewardModel

type GetLoginRewardModel struct {
	Status int64 `json:"status"`

	Item LoginRewardModel `json:"login_reward_model"`
}
