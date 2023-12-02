package account

type RegisterAccounts []RegisterAccount

type RegisterAccount struct {
	Name string `json:"name"`

	Email string `json:"email"`

	Password string `json:"password"`
}
