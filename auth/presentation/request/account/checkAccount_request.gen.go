package account

type CheckAccounts []CheckAccount

type CheckAccount struct {
	Email string `json:"email"`
}
