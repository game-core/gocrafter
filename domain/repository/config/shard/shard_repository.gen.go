//go:generate mockgen -source=./shard_repository.gen.go -destination=./shard_repository_mock.gen.go -package=shard
package shard

import (
	"github.com/game-core/gocrafter/domain/entity/config/shard"
	"gorm.io/gorm"
)

type ShardRepository interface {
	Create(entity *shard.Shard, tx *gorm.DB) (*shard.Shard, error)

	Delete(entity *shard.Shard, tx *gorm.DB) error

	FindByID(ID int64) (*shard.Shard, error)

	FindByShardKey(ShardKey string) (*shard.Shard, error)

	FindOrNilByID(ID int64) (*shard.Shard, error)

	FindOrNilByShardKey(ShardKey string) (*shard.Shard, error)

	List(limit int) (*shard.Shards, error)

	ListByShardKey(ShardKey string) (*shard.Shards, error)

	Save(entity *shard.Shard, tx *gorm.DB) (*shard.Shard, error)
}
