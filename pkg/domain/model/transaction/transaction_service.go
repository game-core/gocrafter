//go:generate mockgen -source=./transaction_service.go -destination=./transaction_service_mock.gen.go -package=transaction
package transaction

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/commonTransaction"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/masterTransaction"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/userTransaction"
)

type TransactionService interface {
	CommonMysqlBegin(ctx context.Context) (*gorm.DB, error)
	CommonMysqlEnd(ctx context.Context, tx *gorm.DB, err error)
	MasterMysqlBegin(ctx context.Context) (*gorm.DB, error)
	MasterMysqlEnd(ctx context.Context, tx *gorm.DB, err error)
	UserMysqlBegin(ctx context.Context, userId string) (*gorm.DB, error)
	UserMysqlEnd(ctx context.Context, tx *gorm.DB, err error)
	MultiUserMysqlBegin(ctx context.Context, userIds []string) (map[string]*gorm.DB, error)
	MultiUserMysqlEnd(ctx context.Context, txs map[string]*gorm.DB, err error)
	UserRedisBegin() redis.Pipeliner
	UserRedisEnd(ctx context.Context, tx redis.Pipeliner, err error)
}

type transactionService struct {
	commonTransactionMysqlRepository commonTransaction.CommonTransactionMysqlRepository
	masterTransactionMysqlRepository masterTransaction.MasterTransactionMysqlRepository
	userTransactionMysqlRepository   userTransaction.UserTransactionMysqlRepository
	userTransactionRedisRepository   userTransaction.UserTransactionRedisRepository
}

func NewTransactionService(
	commonTransactionMysqlRepository commonTransaction.CommonTransactionMysqlRepository,
	masterTransactionMysqlRepository masterTransaction.MasterTransactionMysqlRepository,
	userTransactionMysqlRepository userTransaction.UserTransactionMysqlRepository,
	userTransactionRedisRepository userTransaction.UserTransactionRedisRepository,
) TransactionService {
	return &transactionService{
		commonTransactionMysqlRepository: commonTransactionMysqlRepository,
		masterTransactionMysqlRepository: masterTransactionMysqlRepository,
		userTransactionMysqlRepository:   userTransactionMysqlRepository,
		userTransactionRedisRepository:   userTransactionRedisRepository,
	}
}

// CommonMysqlBegin トランザクションを開始する
func (s *transactionService) CommonMysqlBegin(ctx context.Context) (*gorm.DB, error) {
	tx, err := s.commonTransactionMysqlRepository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// CommonMysqlEnd トランザクションを終了する
func (s *transactionService) CommonMysqlEnd(ctx context.Context, tx *gorm.DB, err error) {
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

// MasterMysqlBegin トランザクションを開始する
func (s *transactionService) MasterMysqlBegin(ctx context.Context) (*gorm.DB, error) {
	tx, err := s.masterTransactionMysqlRepository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// MasterMysqlEnd トランザクションを終了する
func (s *transactionService) MasterMysqlEnd(ctx context.Context, tx *gorm.DB, err error) {
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

// UserMysqlBegin トランザクションを開始する
func (s *transactionService) UserMysqlBegin(ctx context.Context, userId string) (*gorm.DB, error) {
	tx, err := s.userTransactionMysqlRepository.Begin(ctx, keys.GetShardKeyByUserId(userId))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// UserMysqlEnd トランザクションを終了する
func (s *transactionService) UserMysqlEnd(ctx context.Context, tx *gorm.DB, err error) {
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

// MultiUserMysqlBegin マルチトランザクションを開始する
func (s *transactionService) MultiUserMysqlBegin(ctx context.Context, userIds []string) (map[string]*gorm.DB, error) {
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

// MultiUserMysqlEnd マルチトランザクションを終了する
func (s *transactionService) MultiUserMysqlEnd(ctx context.Context, txs map[string]*gorm.DB, err error) {
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

// UserRedisBegin トランザクションを開始する
func (s *transactionService) UserRedisBegin() redis.Pipeliner {
	return s.userTransactionRedisRepository.Begin()
}

// UserRedisEnd トランザクションを終了する
func (s *transactionService) UserRedisEnd(ctx context.Context, tx redis.Pipeliner, err error) {
	if err != nil {
		s.userTransactionRedisRepository.Rollback(tx)
		log.Panicln(err)
	} else {
		if err := s.userTransactionRedisRepository.Commit(ctx, tx); err != nil {
			log.Panicln(err)
		}
	}
}
