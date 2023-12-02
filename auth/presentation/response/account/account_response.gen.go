package account

type Accounts []Account

type Account struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Email string `json:"email"`

	Password string `json:"password"`

	Token string `json:"token"`
}

func ToAccount(ID int64, Name string, Email string, Password string, Token string) *Account {
	return &Account{
		ID:       ID,
		Name:     Name,
		Email:    Email,
		Password: Password,
		Token:    Token,
	}
}
