// Package userAccount ユーザーアカウント
//
//go:generate mockgen -source=./user_account_redis_repository.gen.go -destination=./user_account_redis_repository_mock.gen.go -package=userAccount
package userAccount

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type UserAccountRedisRepository interface {
	Find(ctx context.Context, userId string) (*UserAccount, error)
	FindOrNil(ctx context.Context, userId string) (*UserAccount, error)
	Set(ctx context.Context, tx redis.Pipeliner, m *UserAccount) (*UserAccount, error)
	Delete(ctx context.Context, tx redis.Pipeliner, m *UserAccount) error
}
