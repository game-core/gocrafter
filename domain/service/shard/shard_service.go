//go:generate mockgen -source=./shard_service.go -destination=./shard_service_mock.gen.go -package=shard
package shard

import (
	"errors"
	"log"

	response "github.com/game-core/gocrafter/api/presentation/response/shard"
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

	srs, err := s.shardRepository.List(64)
	if err != nil {
		return nil, err
	}

	if len(*srs) == 0 {
		return nil, errors.New("failed to shardRepository.List")
	}

	shards := make(response.Shards, len(*srs))
	minShard := (*srs)[0]

	for i, sr := range *srs {
		shard := &response.Shard{
			ID:       sr.ID,
			ShardKey: sr.ShardKey,
			Name:     sr.Name,
			Count:    sr.Count,
		}
		shards[i] = *shard

		if sr.Count < minShard.Count {
			minShard = sr
		}
	}

	minShard.Count++
	if _, err := s.shardRepository.Save(&minShard, tx); err != nil {
		return nil, err
	}

	return &response.GetShard{
		Status:       200,
		NextShardKey: minShard.ShardKey,
		Shards:       &shards,
	}, nil
}
