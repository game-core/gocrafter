package master

import (
	"github.com/jinzhu/gorm"

	"github.com/game-core/gocrafter/config/database"
	repository "github.com/game-core/gocrafter/domain/repository/master"
)

type transactionDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewTransactionDao(conn *database.SqlHandler) repository.TransactionRepository {
	return &transactionDao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
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
