//go:generate mockgen -source=./shard_repository.gen.go -destination=./shard_repository_mock.gen.go -package=shard
package shard

import (
	"github.com/game-core/gocrafter/domain/entity/config/shard"
	"github.com/jinzhu/gorm"
)

type ShardRepository interface {
	Create(entity *shard.Shard, tx *gorm.DB) (*shard.Shard, error)

	Delete(entity *shard.Shard, tx *gorm.DB) error

	FindByID(ID int64) (*shard.Shard, error)

	FindByShardKey(ShardKey int) (*shard.Shard, error)

	List(limit int64) (*shard.Shards, error)

	ListByShardKey(ShardKey int) (*shard.Shards, error)

	Update(entity *shard.Shard, tx *gorm.DB) (*shard.Shard, error)
}
