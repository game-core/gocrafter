package account

type CheckAccounts []CheckAccount

type CheckAccount struct {
	Status int64 `json:"status"`

	Item Account `json:"item"`
}
