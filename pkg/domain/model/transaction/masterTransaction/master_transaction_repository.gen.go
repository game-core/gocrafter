//go:generate mockgen -source=./master_transaction_repository.gen.go -destination=./master_transaction_repository_mock.gen.go -package=masterTransaction
package masterTransaction

import (
	"context"

	"gorm.io/gorm"
)

type MasterTransactionRepository interface {
	Begin(ctx context.Context) (*gorm.DB, error)
	Commit(ctx context.Context, tx *gorm.DB) error
	Rollback(ctx context.Context, tx *gorm.DB) error
}
