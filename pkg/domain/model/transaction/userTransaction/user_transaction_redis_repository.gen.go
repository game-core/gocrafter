//go:generate mockgen -source=./user_transaction_redis_repository.gen.go -destination=./user_transaction_redis_repository_mock.gen.go -package=userTransaction
package userTransaction

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type UserTransactionRedisRepository interface {
	Begin() redis.Pipeliner
	Commit(ctx context.Context, tx redis.Pipeliner) error
	Rollback(tx redis.Pipeliner)
}
