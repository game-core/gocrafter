package user

import (
	"github.com/jinzhu/gorm"
)

type TransactionRepository interface {
	Begin() (tx *gorm.DB, err error)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)
}
