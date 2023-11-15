package account

type LoginAccount struct {
	Status int64 `json:"status"`

	Item Account `json:"item"`
}
