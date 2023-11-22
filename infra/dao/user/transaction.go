package user

import (
	"github.com/jinzhu/gorm"

	"github.com/game-core/gocrafter/config/database"
	repository "github.com/game-core/gocrafter/domain/repository/user"
)

type transactionDao struct {
	ShardedConn *database.ShardedConn
}

func NewTransactionDao(conn *database.SqlHandler) repository.TransactionRepository {
	return &transactionDao{
		ShardedConn: conn.User,
	}
}

func (d *transactionDao) Begin(accountID int64) (tx *gorm.DB, err error) {
	tx = d.ShardedConn.Shards[shardKey(accountID, len(d.ShardedConn.Shards))].WriteConn.Begin()
	if err := tx.Error; err != nil {
		return tx, err
	}

	return tx, err
}

func (d *transactionDao) Commit(tx *gorm.DB) (err error) {
	tx.Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return err
}

func (d *transactionDao) Rollback(tx *gorm.DB) (err error) {
	tx.Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return err
}

func shardKey(accountID int64, shardCount int) int {
	return int(accountID) % shardCount
}
