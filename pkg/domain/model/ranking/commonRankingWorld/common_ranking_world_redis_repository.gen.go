// Package commonRankingWorld ワールドランキング
//
//go:generate mockgen -source=./common_ranking_world_redis_repository.gen.go -destination=./common_ranking_world_redis_repository_mock.gen.go -package=commonRankingWorld
package commonRankingWorld

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CommonRankingWorldRedisRepository interface {
	Find(ctx context.Context, masterRankingId int64, userId string) (*CommonRankingWorld, error)
	FindOrNil(ctx context.Context, masterRankingId int64, userId string) (*CommonRankingWorld, error)
	Set(ctx context.Context, tx redis.Pipeliner, m *CommonRankingWorld) (*CommonRankingWorld, error)
	Delete(ctx context.Context, tx redis.Pipeliner, m *CommonRankingWorld) error
}
