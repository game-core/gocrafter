package commonShard

import (
	"context"

	"github.com/game-core/gocrafter/internal/errors"
)

var CommonShardInstances CommonShards

// GetShardKey シャードキーを取得する
func (s *CommonShards) GetShardKey(ctx context.Context, commonShardRepository CommonShardRepository) (string, error) {
	if len(CommonShardInstances) <= 0 {
		commonShards, err := commonShardRepository.FindList(ctx)
		if err != nil {
			return "", errors.NewMethodError("s.commonShardRepository.FindList", err)
		}
		if len(commonShards) <= 0 {
			return "", errors.NewError("common_shard does not exist")
		}

		CommonShardInstances = commonShards
	}

	minShard := CommonShardInstances[0]
	minIndex := 0
	for i, s := range CommonShardInstances {
		if s.Count < minShard.Count {
			minShard = s
			minIndex = i
			break
		}
	}

	minShard.Count++
	CommonShardInstances[minIndex] = minShard

	return minShard.ShardKey, nil
}
