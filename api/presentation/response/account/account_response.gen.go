package account

type Accounts []Account

type Account struct {
	ID int64 `json:"id"`

	ShardKey string `json:"shard_key"`

	UUID string `json:"uuid"`

	Name string `json:"name"`

	Password string `json:"password"`

	Token string `json:"token"`
}

func ToAccount(ID int64, ShardKey string, UUID string, Name string, Password string, Token string) *Account {
	return &Account{
		ID:       ID,
		ShardKey: ShardKey,
		UUID:     UUID,
		Name:     Name,
		Password: Password,
		Token:    Token,
	}
}
