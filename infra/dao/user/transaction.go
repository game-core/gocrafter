package user

import (
	"github.com/jinzhu/gorm"

	"github.com/game-core/gocrafter/config/database"
	repository "github.com/game-core/gocrafter/domain/repository/user"
)

type transactionDao struct {
	ShardConn *database.ShardConn
}

func NewTransactionDao(conn *database.SqlHandler) repository.TransactionRepository {
	return &transactionDao{
		ShardConn: conn.User,
	}
}

func (d *transactionDao) Begin(shardKey int) (*gorm.DB, error) {
	tx := d.ShardConn.Shards[shardKey].WriteConn.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *transactionDao) Commit(tx *gorm.DB) error {
	tx.Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (d *transactionDao) Rollback(tx *gorm.DB) error {
	tx.Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
