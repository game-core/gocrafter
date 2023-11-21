package account

import (
	"time"
)

type CheckAccounts []CheckAccount

type CheckAccount struct {
	UUID string `json:"uuid"`

	Time time.Time `json:"time"`
}
