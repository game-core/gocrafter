//go:generate mockgen -source=./transaction.go -destination=./transaction_mock.go -package=transaction
package transaction

import (
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Begin(shardKey string) (*gorm.DB, error)
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
	CommitOrRollback(tx *gorm.DB, err error) error
}
