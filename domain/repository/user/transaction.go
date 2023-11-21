//go:generate mockgen -source=./transaction.go -destination=./transaction_mock.go -package=user
package user

import (
	"github.com/jinzhu/gorm"
)

type TransactionRepository interface {
	Begin() (tx *gorm.DB, err error)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)
}
