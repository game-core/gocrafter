package account

import (
	"time"
)

type Accounts []Account

type Account struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Email string `json:"email"`

	Password string `json:"password"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *Account) TableName() string {
	return "account"
}
