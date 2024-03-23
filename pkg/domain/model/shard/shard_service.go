//go:generate mockgen -source=./shard_service.go -destination=./shard_service_mock.gen.go -package=shard
package shard

import (
	"context"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/shard/commonShard"
)

type ShardService interface {
	GetShardKey(ctx context.Context) (string, error)
}

type shardService struct {
	commonShardMysqlRepository commonShard.CommonShardMysqlRepository
}

func NewShardService(
	commonShardMysqlRepository commonShard.CommonShardMysqlRepository,
) ShardService {
	return &shardService{
		commonShardMysqlRepository: commonShardMysqlRepository,
	}
}

// GetShardKey シャードキーを取得して更新する
func (s *shardService) GetShardKey(ctx context.Context) (string, error) {
	commonShards := commonShard.NewCommonShards()
	shardKey, err := commonShards.GetShardKey(ctx, s.commonShardMysqlRepository)
	if err != nil {
		return "", errors.NewMethodError("shards.GetShardKey", err)
	}

	return shardKey, nil
}
