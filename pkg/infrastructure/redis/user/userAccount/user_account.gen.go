// Package userAccount ユーザーアカウント
package userAccount

import (
	"encoding/json"

	"time"
)

type UserAccounts []*UserAccount

type UserAccount struct {
	UserId    string
	Name      string
	Password  string
	LoginAt   time.Time
	LogoutAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserAccount() *UserAccount {
	return &UserAccount{}
}

func NewUserAccounts() UserAccounts {
	return UserAccounts{}
}

func SetUserAccount(userId string, name string, password string, loginAt time.Time, logoutAt time.Time, createdAt time.Time, updatedAt time.Time) *UserAccount {
	return &UserAccount{
		UserId:    userId,
		Name:      name,
		Password:  password,
		LoginAt:   loginAt,
		LogoutAt:  logoutAt,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (t *UserAccount) TableToJson() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (t *UserAccount) JsonToTable(data string) error {
	m := NewUserAccount()
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return err
	}

	return nil
}

func (t *UserAccount) TableName() string {
	return "user_account"
}
