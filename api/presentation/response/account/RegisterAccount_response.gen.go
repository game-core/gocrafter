package account

type RegisterAccount struct {
	Status int64 `json:"status"`

	ID int64 `json:"id"`

	UUID string `json:"uuid"`

	Name string `json:"name"`

	Password string `json:"password"`
}
