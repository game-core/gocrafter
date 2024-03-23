//go:generate mockgen -source=./user_transaction_repository.gen.go -destination=./user_transaction_repository_mock.gen.go -package=userTransaction
package userTransaction

import (
	"context"

	"gorm.io/gorm"
)

type UserTransactionMysqlRepository interface {
	Begin(ctx context.Context, shardKey string) (*gorm.DB, error)
	Commit(ctx context.Context, tx *gorm.DB) error
	Rollback(ctx context.Context, tx *gorm.DB) error
}
