package account

type Accounts []Account

type Account struct {
	ID int64 `json:"id"`

	UUID string `json:"uuid"`

	Name string `json:"name"`

	Password string `json:"password"`

	Token string `json:"token"`
}
