package account

type CheckAccounts []CheckAccount

type CheckAccount struct {
	ShardKey int `json:"shard_key"`

	UUID string `json:"uuid"`
}
