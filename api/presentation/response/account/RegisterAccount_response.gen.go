package account

type RegisterAccount struct {
	Status int64 `json:"status"`

	ID int64 `json:"id"`

	UUID string `json:"uuid"`

	Name string `json:"name"`

	Password string `json:"password"`
}

func RegisterAccountResponse(status int64, iD int64, uUID string, name string, password string) *RegisterAccount {
	return &RegisterAccount{Status: status, ID: iD, UUID: uUID, Name: name, Password: password}
}
