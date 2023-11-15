package account

type CheckAccount struct {
	Status int64 `json:"status"`

	Item Account `json:"item"`
}
