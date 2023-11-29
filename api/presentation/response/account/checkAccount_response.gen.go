package account

type CheckAccounts []CheckAccount

type CheckAccount struct {
	Status int64 `json:"status"`

	Item Account `json:"item"`
}

func ToCheckAccount(Status int64, Item Account) *CheckAccount {
	return &CheckAccount{
		Status: Status,
		Item:   Item,
	}
}
