package account

type LoginAccounts []LoginAccount

type LoginAccount struct {
	Email string `json:"email"`

	Password string `json:"password"`
}
