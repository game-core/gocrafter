// Package userAccountToken ユーザーアカウントトークン
//
//go:generate mockgen -source=./user_account_token_mysql_repository.gen.go -destination=./user_account_token_mysql_repository_mock.gen.go -package=userAccountToken
package userAccountToken

import (
	"context"

	"gorm.io/gorm"
)

type UserAccountTokenMysqlRepository interface {
	Find(ctx context.Context, userId string) (*UserAccountToken, error)
	FindOrNil(ctx context.Context, userId string) (*UserAccountToken, error)
	FindList(ctx context.Context, userId string) (UserAccountTokens, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserAccountToken) (*UserAccountToken, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserAccountTokens) (UserAccountTokens, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserAccountToken) (*UserAccountToken, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserAccountToken) error
}
