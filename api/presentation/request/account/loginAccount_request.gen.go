package account

type LoginAccounts []LoginAccount

type LoginAccount struct {
	ShardKey int `json:"shard_key"`

	UUID string `json:"uuid"`

	Password string `json:"password"`
}
