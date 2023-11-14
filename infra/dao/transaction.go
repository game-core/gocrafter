package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/architecture-template/echo-ddd/config/db"
	"github.com/architecture-template/echo-ddd/domain/repository"
)

type transactionDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewTransactionDao(conn *db.SqlHandler) repository.TransactionRepository {
	return &transactionDao{
		Read:  conn.ReadConn,
		Write: conn.WriteConn,
	}
}

func (transactionDao *transactionDao) Begin() (tx *gorm.DB, err error) {
	tx = transactionDao.Write.Begin()
	if err := tx.Error; err != nil {
		return tx, err
	}

	return tx, err
}

func (transactionDao *transactionDao) Commit(tx *gorm.DB) (err error) {
	tx.Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return err
}

func (transactionDao *transactionDao) Rollback(tx *gorm.DB) (err error) {
	tx.Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return err
}
