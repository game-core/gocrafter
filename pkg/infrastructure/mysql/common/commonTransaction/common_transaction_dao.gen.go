package commonTransaction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/commonTransaction"
)

type commonTransactionDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
}

func NewCommonTransactionDao(conn *database.SqlHandler) commonTransaction.CommonTransactionRepository {
	return &commonTransactionDao{
		ReadConn:  conn.Common.ReadConn,
		WriteConn: conn.Common.WriteConn,
	}
}

func (d *commonTransactionDao) Begin(ctx context.Context) (*gorm.DB, error) {
	tx := d.WriteConn.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *commonTransactionDao) Commit(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (d *commonTransactionDao) Rollback(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
