package masterTransaction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/masterTransaction"
)

type masterTransactionDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
}

func NewMasterTransactionDao(conn *database.MysqlHandler) masterTransaction.MasterTransactionRepository {
	return &masterTransactionDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
	}
}

func (d *masterTransactionDao) Begin(ctx context.Context) (*gorm.DB, error) {
	tx := d.WriteConn.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *masterTransactionDao) Commit(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (d *masterTransactionDao) Rollback(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
