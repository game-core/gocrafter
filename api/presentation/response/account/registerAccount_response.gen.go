package account

type RegisterAccounts []RegisterAccount

type RegisterAccount struct {
	Status int64 `json:"status"`

	Item Account `json:"item"`
}

func ToRegisterAccount(Status int64, Item Account) *RegisterAccount {
	return &RegisterAccount{
		Status: Status,
		Item:   Item,
	}
}
