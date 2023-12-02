package account

type RegisterAccounts []RegisterAccount

type RegisterAccount struct {
	Status int64 `json:"status"`

	Account Account `json:"account"`
}

func ToRegisterAccount(Status int64, Account Account) *RegisterAccount {
	return &RegisterAccount{
		Status:  Status,
		Account: Account,
	}
}
