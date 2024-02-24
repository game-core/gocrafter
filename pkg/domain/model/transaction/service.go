//go:generate mockgen -source=./service.go -destination=./service_mock.gen.go -package=service
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
	commonTransactionRepository commonTransaction.CommonTransactionRepository
	masterTransactionRepository masterTransaction.MasterTransactionRepository
	userTransactionRepository   userTransaction.UserTransactionRepository
}

func NewTransactionService(
	commonTransactionRepository commonTransaction.CommonTransactionRepository,
	masterTransactionRepository masterTransaction.MasterTransactionRepository,
	userTransactionRepository userTransaction.UserTransactionRepository,
) TransactionService {
	return &transactionService{
		commonTransactionRepository: commonTransactionRepository,
		masterTransactionRepository: masterTransactionRepository,
		userTransactionRepository:   userTransactionRepository,
	}
}

// CommonBegin トランザクションを開始する
func (s *transactionService) CommonBegin(ctx context.Context) (*gorm.DB, error) {
	tx, err := s.commonTransactionRepository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// CommonEnd トランザクションを終了する
func (s *transactionService) CommonEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.commonTransactionRepository.Rollback(ctx, tx); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := s.commonTransactionRepository.Commit(ctx, tx); err != nil {
			log.Panicln(err)
		}
	}
}

// MasterBegin トランザクションを開始する
func (s *transactionService) MasterBegin(ctx context.Context) (*gorm.DB, error) {
	tx, err := s.masterTransactionRepository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// MasterEnd トランザクションを終了する
func (s *transactionService) MasterEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.masterTransactionRepository.Rollback(ctx, tx); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := s.masterTransactionRepository.Commit(ctx, tx); err != nil {
			log.Panicln(err)
		}
	}
}

// UserBegin トランザクションを開始する
func (s *transactionService) UserBegin(ctx context.Context, userId string) (*gorm.DB, error) {
	tx, err := s.userTransactionRepository.Begin(ctx, keys.GetShardKeyByUserId(userId))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// UserEnd トランザクションを終了する
func (s *transactionService) UserEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.userTransactionRepository.Rollback(ctx, tx); err != nil {
			log.Panicln(err)
		}
	} else {
		if err := s.userTransactionRepository.Commit(ctx, tx); err != nil {
			log.Panicln(err)
		}
	}
}

// MultiUserBegin マルチトランザクションを開始する
func (s *transactionService) MultiUserBegin(ctx context.Context, userIds []string) (map[string]*gorm.DB, error) {
	txs := make(map[string]*gorm.DB)
	for _, userId := range userIds {
		tx, err := s.userTransactionRepository.Begin(ctx, keys.GetShardKeyByUserId(userId))
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
			if rollbackErr := s.userTransactionRepository.Rollback(ctx, tx); rollbackErr != nil {
				log.Panicln(rollbackErr)
			}
		}
		return
	}

	for _, tx := range txs {
		if commitErr := s.userTransactionRepository.Commit(ctx, tx); commitErr != nil {
			log.Panicln(commitErr)
		}
	}
}
