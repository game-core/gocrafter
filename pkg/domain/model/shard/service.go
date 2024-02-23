package shard

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/model/shard/commonShard"
)

type ShardService interface {
	GetShardKeyAndUpdate(ctx context.Context, tx *gorm.DB) (string, error)
}

type shardService struct {
	commonShardRepository commonShard.CommonShardRepository
}

func NewShardService(
	commonShardRepository commonShard.CommonShardRepository,
) ShardService {
	return &shardService{
		commonShardRepository: commonShardRepository,
	}
}

// GetShardKeyAndUpdate シャードキーを取得して更新する
func (s *shardService) GetShardKeyAndUpdate(ctx context.Context, tx *gorm.DB) (string, error) {
	commonShard, err := s.commonShardRepository.FindList(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to s.commonShardRepository.FindList: %s", err)
	}
	if len(commonShard) == 0 {
		return "", fmt.Errorf("common_shard does not exist")
	}

	minShard := (commonShard)[0]
	for _, s := range commonShard {
		if s.Count < minShard.Count {
			minShard = s
		}
	}
	minShard.Count++

	if _, err := s.commonShardRepository.Update(ctx, tx, minShard); err != nil {
		return "", fmt.Errorf("failed to s.commonShardRepository.Update: %s", err)
	}

	return minShard.ShardKey, nil
}
