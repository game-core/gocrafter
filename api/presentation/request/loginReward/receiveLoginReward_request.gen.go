package loginReward

type ReceiveLoginRewards []ReceiveLoginReward

type ReceiveLoginReward struct {
	ShardKey int `json:"shard_key"`

	UUID string `json:"uuid"`
}
