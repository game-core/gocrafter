package commonTransaction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/commonTransaction"
)

type commonTransactionDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewCommonTransactionDao(conn *database.MysqlHandler) commonTransaction.CommonTransactionMysqlRepository {
	return &commonTransactionDao{
		ReadMysqlConn:  conn.Common.ReadMysqlConn,
		WriteMysqlConn: conn.Common.WriteMysqlConn,
	}
}

func (d *commonTransactionDao) Begin(ctx context.Context) (*gorm.DB, error) {
	tx := d.WriteMysqlConn.WithContext(ctx).Begin()
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
