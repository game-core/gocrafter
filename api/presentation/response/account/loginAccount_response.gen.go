package account

type LoginAccounts []LoginAccount

type LoginAccount struct {
	Status int64 `json:"status"`

	Item Account `json:"item"`
}

func ToLoginAccount(Status int64, Item Account) *LoginAccount {
	return &LoginAccount{
		Status: Status,
		Item:   Item,
	}
}
