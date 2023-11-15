package account

type RegisterAccount struct {
	Status int64 `json:"status"`

	Item Account `json:"item"`
}
