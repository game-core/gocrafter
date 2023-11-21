package account

type LoginAccounts []LoginAccount

type LoginAccount struct {
	UUID string `json:"uuid"`

	Password string `json:"password"`
}
