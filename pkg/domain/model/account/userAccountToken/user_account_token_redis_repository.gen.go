// Package userAccountToken ユーザーアカウントトークン
//
//go:generate mockgen -source=./user_account_token_redis_repository.gen.go -destination=./user_account_token_redis_repository_mock.gen.go -package=userAccountToken
package userAccountToken

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type UserAccountTokenRedisRepository interface {
	Find(ctx context.Context, userId string) (*UserAccountToken, error)
	FindOrNil(ctx context.Context, userId string) (*UserAccountToken, error)
	Set(ctx context.Context, tx redis.Pipeliner, m *UserAccountToken) (*UserAccountToken, error)
	Delete(ctx context.Context, tx redis.Pipeliner, m *UserAccountToken) error
}
