package masterTransaction

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction/userTransaction"
)

type userTransactionDao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func NewUserTransactionDao(conn *database.RedisHandler) userTransaction.UserTransactionRedisRepository {
	return &userTransactionDao{
		ReadRedisConn:  conn.User.ReadRedisConn,
		WriteRedisConn: conn.User.WriteRedisConn,
	}
}

func (d *userTransactionDao) Begin() redis.Pipeliner {
	return d.WriteRedisConn.TxPipeline()
}

func (d *userTransactionDao) Commit(ctx context.Context, tx redis.Pipeliner) error {
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (d *userTransactionDao) Rollback(tx redis.Pipeliner) {
	tx.Discard()
}
