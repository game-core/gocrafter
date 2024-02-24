//go:generate mockgen -source=./common_transaction_repository.go -destination=./common_transaction_repository_mock.go -package=commonTransaction
package commonTransaction

import (
	"context"

	"gorm.io/gorm"
)

type CommonTransactionRepository interface {
	Begin(ctx context.Context) (*gorm.DB, error)
	Commit(ctx context.Context, tx *gorm.DB) error
	Rollback(ctx context.Context, tx *gorm.DB) error
}
