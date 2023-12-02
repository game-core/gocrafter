package account

type CheckAccounts []CheckAccount

type CheckAccount struct {
	Status int64 `json:"status"`

	Account Account `json:"account"`
}

func ToCheckAccount(Status int64, Account Account) *CheckAccount {
	return &CheckAccount{
		Status:  Status,
		Account: Account,
	}
}
