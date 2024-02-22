// Package userAccount ユーザーアカウント
package userAccount

import (
	"time"
)

type UserAccounts []*UserAccount

type UserAccount struct {
	UserId   string
	Name     string
	Password string
	LoginAt  time.Time
	LogoutAt time.Time
}

func NewUserAccount() *UserAccount {
	return &UserAccount{}
}

func NewUserAccounts() UserAccounts {
	return UserAccounts{}
}

func SetUserAccount(userId string, name string, password string, loginAt time.Time, logoutAt time.Time) *UserAccount {
	return &UserAccount{
		UserId:   userId,
		Name:     name,
		Password: password,
		LoginAt:  loginAt,
		LogoutAt: logoutAt,
	}
}
