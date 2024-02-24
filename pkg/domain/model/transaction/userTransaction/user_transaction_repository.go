//go:generate mockgen -source=./user_transaction_repository.go -destination=./user_transaction_repository_mock.go -package=userTransaction
package userTransaction

import (
	"context"

	"gorm.io/gorm"
)

type UserTransactionRepository interface {
	Begin(ctx context.Context, shardKey string) (*gorm.DB, error)
	Commit(ctx context.Context, tx *gorm.DB) error
	Rollback(ctx context.Context, tx *gorm.DB) error
}
