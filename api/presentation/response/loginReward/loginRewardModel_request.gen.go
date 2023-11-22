package loginReward

type LoginRewardModels []LoginRewardModel

type LoginRewardModel struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Event Event `json:"event"`
}
