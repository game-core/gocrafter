// Package userAccountToken ユーザーアカウントトークン
package userAccountToken

import (
	"time"
)

type UserAccountTokens []*UserAccountToken

type UserAccountToken struct {
	UserId    string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserAccountToken() *UserAccountToken {
	return &UserAccountToken{}
}

func NewUserAccountTokens() UserAccountTokens {
	return UserAccountTokens{}
}

func SetUserAccountToken(userId string, token string, createdAt time.Time, updatedAt time.Time) *UserAccountToken {
	return &UserAccountToken{
		UserId:    userId,
		Token:     token,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (t *UserAccountToken) TableName() string {
	return "user_account_token"
}
