// Package userAccount ユーザーアカウント
//
//go:generate mockgen -source=./user_account_repository.gen.go -destination=./user_account_repository_mock.gen.go -package=userAccount
package userAccount

import (
	"context"

	"gorm.io/gorm"
)

type UserAccountMysqlRepository interface {
	Find(ctx context.Context, userId string) (*UserAccount, error)
	FindOrNil(ctx context.Context, userId string) (*UserAccount, error)
	FindList(ctx context.Context, userId string) (UserAccounts, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserAccount) (*UserAccount, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserAccounts) (UserAccounts, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserAccount) (*UserAccount, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserAccount) error
}
