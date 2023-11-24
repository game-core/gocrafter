package loginReward

type ReceiveLoginRewards []ReceiveLoginReward

type ReceiveLoginReward struct {
	ShardKey string `json:"shard_key"`

	AccountID int64 `json:"account_id"`

	UUID string `json:"uuid"`

	LoginRewardModelName string `json:"login_reward_model_name"`
}
