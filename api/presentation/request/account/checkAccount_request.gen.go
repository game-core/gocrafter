package account

type CheckAccounts []CheckAccount

type CheckAccount struct {
	ShardKey string `json:"shard_key"`

	UUID string `json:"uuid"`
}
