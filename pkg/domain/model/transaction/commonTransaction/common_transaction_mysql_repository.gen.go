//go:generate mockgen -source=./common_transaction_mysql_repository.gen.go -destination=./common_transaction_mysql_repository_mock.gen.go -package=commonTransaction
package commonTransaction

import (
	"context"

	"gorm.io/gorm"
)

type CommonTransactionMysqlRepository interface {
	Begin(ctx context.Context) (*gorm.DB, error)
	Commit(ctx context.Context, tx *gorm.DB) error
	Rollback(ctx context.Context, tx *gorm.DB) error
}
