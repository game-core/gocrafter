package account

type LoginAccounts []LoginAccount

type LoginAccount struct {
	ShardKey string `json:"shard_key"`

	UUID string `json:"uuid"`

	Password string `json:"password"`
}
