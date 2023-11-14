package account

type RegisterAccount struct {
	ID inbt64 `json:"id"`

	UUID string `json:"uuid"`

	Name string `json:"name"`

	Password string `json:"password"`
}

func RegisterAccountResponse(iD inbt64, uUID string, name string, password string) *RegisterAccount {
	return &RegisterAccount{ID: iD, UUID: uUID, Name: name, Password: password}
}
