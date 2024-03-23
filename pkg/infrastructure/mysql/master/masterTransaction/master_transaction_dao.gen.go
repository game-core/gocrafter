package masterTransaction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/masterTransaction"
)

type masterTransactionDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewMasterTransactionDao(conn *database.MysqlHandler) masterTransaction.MasterTransactionMysqlRepository {
	return &masterTransactionDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
	}
}

func (d *masterTransactionDao) Begin(ctx context.Context) (*gorm.DB, error) {
	tx := d.WriteMysqlConn.WithContext(ctx).Begin()
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
