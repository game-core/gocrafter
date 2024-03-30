// Package commonRankingRoom ルームランキング
//
//go:generate mockgen -source=./common_ranking_room_redis_repository.gen.go -destination=./common_ranking_room_redis_repository_mock.gen.go -package=commonRankingRoom
package commonRankingRoom

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CommonRankingRoomRedisRepository interface {
	Find(ctx context.Context, masterRankingId int64, roomId string, userId string) (*CommonRankingRoom, error)
	FindOrNil(ctx context.Context, masterRankingId int64, roomId string, userId string) (*CommonRankingRoom, error)
	Set(ctx context.Context, tx redis.Pipeliner, m *CommonRankingRoom) (*CommonRankingRoom, error)
	Delete(ctx context.Context, tx redis.Pipeliner, m *CommonRankingRoom) error
}
