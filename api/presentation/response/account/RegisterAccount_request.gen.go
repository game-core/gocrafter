package account

type RegisterAccount struct {
	Status int64 `json:"status"`

	ID int64 `json:"id"`

	UUID string `json:"uuid"`

	Name string `json:"name"`

	Password string `json:"password"`
}

func RegisterAccountResponse(uUID string, name string, password string, status int64, iD int64) *RegisterAccount {
	return &RegisterAccount{UUID: uUID, Name: name, Password: password, Status: status, ID: iD}
}
