// Package userAction ユーザーアカウント
package userAction

import (
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

func (t *UserAccount) TableName() string {
	return "user_account"
}
