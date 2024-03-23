//go:generate mockgen -source=./transaction_service.go -destination=./transaction_service_mock.gen.go -package=transaction
package transaction

import (
	"context"
	"log"

	"github.com/game-core/gocrafter/internal/keys"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/model/transaction/commonTransaction"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/masterTransaction"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/userTransaction"
)

type TransactionService interface {
	CommonBegin(ctx context.Context) (*gorm.DB, error)
	CommonEnd(ctx context.Context, tx *gorm.DB, err error)
	MasterBegin(ctx context.Context) (*gorm.DB, error)
	MasterEnd(ctx context.Context, tx *gorm.DB, err error)
	UserBegin(ctx context.Context, userId string) (*gorm.DB, error)
	UserEnd(ctx context.Context, tx *gorm.DB, err error)
	MultiUserBegin(ctx context.Context, userIds []string) (map[string]*gorm.DB, error)
	MultiUserEnd(ctx context.Context, txs map[string]*gorm.DB, err error)
}

type transactionService struct {
	commonTransactionMysqlRepository commonTransaction.CommonTransactionMysqlRepository
	masterTransactionMysqlRepository masterTransaction.MasterTransactionMysqlRepository
	userTransactionMysqlRepository   userTransaction.UserTransactionMysqlRepository
}

func NewTransactionService(
	commonTransactionMysqlRepository commonTransaction.CommonTransactionMysqlRepository,
	masterTransactionMysqlRepository masterTransaction.MasterTransactionMysqlRepository,
	userTransactionMysqlRepository userTransaction.UserTransactionMysqlRepository,
) TransactionService {
	return &transactionService{
		commonTransactionMysqlRepository: commonTransactionMysqlRepository,
		masterTransactionMysqlRepository: masterTransactionMysqlRepository,
		userTransactionMysqlRepository:   userTransactionMysqlRepository,
	}
}

// CommonBegin トランザクションを開始する
func (s *transactionService) CommonBegin(ctx context.Context) (*gorm.DB, error) {
	tx, err := s.commonTransactionMysqlRepository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// CommonEnd トランザクションを終了する
func (s *transactionService) CommonEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.commonTransactionMysqlRepository.Rollback(ctx, tx); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := s.commonTransactionMysqlRepository.Commit(ctx, tx); err != nil {
			log.Panicln(err)
		}
	}
}

// MasterBegin トランザクションを開始する
func (s *transactionService) MasterBegin(ctx context.Context) (*gorm.DB, error) {
	tx, err := s.masterTransactionMysqlRepository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// MasterEnd トランザクションを終了する
func (s *transactionService) MasterEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.masterTransactionMysqlRepository.Rollback(ctx, tx); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := s.masterTransactionMysqlRepository.Commit(ctx, tx); err != nil {
			log.Panicln(err)
		}
	}
}

// UserBegin トランザクションを開始する
func (s *transactionService) UserBegin(ctx context.Context, userId string) (*gorm.DB, error) {
	tx, err := s.userTransactionMysqlRepository.Begin(ctx, keys.GetShardKeyByUserId(userId))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// UserEnd トランザクションを終了する
func (s *transactionService) UserEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.userTransactionMysqlRepository.Rollback(ctx, tx); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := s.userTransactionMysqlRepository.Commit(ctx, tx); err != nil {
			log.Panicln(err)
		}
	}
}

// MultiUserBegin マルチトランザクションを開始する
func (s *transactionService) MultiUserBegin(ctx context.Context, userIds []string) (map[string]*gorm.DB, error) {
	txs := make(map[string]*gorm.DB)
	for _, userId := range userIds {
		tx, err := s.userTransactionMysqlRepository.Begin(ctx, keys.GetShardKeyByUserId(userId))
		if err != nil {
			return nil, err
		}

		txs[userId] = tx
	}

	return txs, nil
}

// MultiUserEnd マルチトランザクションを終了する
func (s *transactionService) MultiUserEnd(ctx context.Context, txs map[string]*gorm.DB, err error) {
	if err != nil {
		for _, tx := range txs {
			if rollbackErr := s.userTransactionMysqlRepository.Rollback(ctx, tx); rollbackErr != nil {
				log.Panicln(rollbackErr)
			}
		}
		return
	}

	for _, tx := range txs {
		if commitErr := s.userTransactionMysqlRepository.Commit(ctx, tx); commitErr != nil {
			log.Panicln(commitErr)
		}
	}
}
