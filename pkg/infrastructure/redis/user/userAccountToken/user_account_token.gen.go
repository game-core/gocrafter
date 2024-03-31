// Package userAccountToken ユーザーアカウントトークン
package userAccountToken

import (
	"encoding/json"

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

func (t *UserAccountToken) TableToJson() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (t *UserAccountToken) JsonToTable(data string) error {
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		return err
	}

	return nil
}

func (t *UserAccountToken) TableName() string {
	return "user_account_token"
}
