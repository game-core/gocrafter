//go:generate mockgen -source=./shard_service.go -destination=./shard_service_mock.gen.go -package=shard
package shard

import (
	"errors"
	"log"

	response "github.com/game-core/gocrafter/api/presentation/response/shard"
	shardEntity "github.com/game-core/gocrafter/domain/entity/config/shard"
	configRepository "github.com/game-core/gocrafter/domain/repository/config"
	shardRepository "github.com/game-core/gocrafter/domain/repository/config/shard"
)

type ShardService interface {
	GetShard() (*response.GetShard, error)
}

type shardService struct {
	transactionRepository configRepository.TransactionRepository
	shardRepository       shardRepository.ShardRepository
}

func NewShardService(
	shardRepository shardRepository.ShardRepository,
	transactionRepository configRepository.TransactionRepository,
) ShardService {
	return &shardService{
		transactionRepository: transactionRepository,
		shardRepository:       shardRepository,
	}
}

// GetShard シャード設定を取得する
func (s *shardService) GetShard() (*response.GetShard, error) {
	// transaction
	tx, err := s.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := s.transactionRepository.Rollback(tx); err != nil {
				log.Panicln(err)
			}
		} else {
			if err := s.transactionRepository.Commit(tx); err != nil {
				log.Panicln(err)
			}
		}
	}()

	ss, err := s.shardRepository.List(64)
	if err != nil {
		return nil, err
	}

	if len(*ss) == 0 {
		return nil, errors.New("failed to shardRepository.List")
	}

	shards := make(response.Shards, len(*ss))
	minShard := s.getMinShard(ss, shards)
	minShard.Count++

	if _, err := s.shardRepository.Save(minShard, tx); err != nil {
		return nil, err
	}

	return response.ToGetShard(200, minShard.ShardKey, &shards), nil
}

// getMinShard 最小のシャードを取得
func (s *shardService) getMinShard(ss *shardEntity.Shards, shards response.Shards) *shardEntity.Shard {
	minShard := (*ss)[0]

	for i, s := range *ss {
		shard := &response.Shard{
			ID:       s.ID,
			ShardKey: s.ShardKey,
			Name:     s.Name,
			Count:    s.Count,
		}
		shards[i] = *shard

		if s.Count < minShard.Count {
			minShard = s
		}
	}

	return &minShard
}
