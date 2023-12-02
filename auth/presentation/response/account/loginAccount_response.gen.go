package account

type LoginAccounts []LoginAccount

type LoginAccount struct {
	Status int64 `json:"status"`

	Account Account `json:"account"`
}

func ToLoginAccount(Status int64, Account Account) *LoginAccount {
	return &LoginAccount{
		Status:  Status,
		Account: Account,
	}
}
