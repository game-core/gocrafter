package loginReward

type ReceiveLoginRewards []ReceiveLoginReward

type ReceiveLoginReward struct {
	ShardKey int `json:"shard_key"`

	AccountID int64 `json:"account_id"`

	UUID string `json:"uuid"`

	LoginRewardModelID int64 `json:"login_reward_model_id"`
}
