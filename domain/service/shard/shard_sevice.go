package shard

import (
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

	srs, err := s.shardRepository.List(64)
	if err != nil {
		return nil, err
	}

	shards := make(response.Shards, len(*srs))
	minShard := &shardEntity.Shard{}

	for i, sr := range *srs {
		shard := &response.Shard{
			ID:       sr.ID,
			ShardKey: sr.ShardKey,
			Name:     sr.Name,
			Count:    sr.Count,
		}
		shards[i] = *shard

		if i == 0 || sr.ShardKey < minShard.ShardKey {
			minShard = &sr
		}
	}

	if _, err := s.shardRepository.Update(minShard, tx); err != nil {
		return nil, err
	}

	return &response.GetShard{
		Status:       200,
		NextShardKey: minShard.ShardKey,
		Shards:       &shards,
	}, nil
}
