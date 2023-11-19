package account

type RegisterAccounts []RegisterAccount

type RegisterAccount struct {
	Status int64 `json:"status"`

	Item Account `json:"item"`
}
