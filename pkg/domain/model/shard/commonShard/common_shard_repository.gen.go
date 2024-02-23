// Package commonShard シャード管理
//
//go:generate mockgen -source=./common_shard_repository.gen.go -destination=./common_shard_repository_mock.gen.go -package=commonShard
package commonShard

import (
	context "context"

	"gorm.io/gorm"
)

type CommonShardRepository interface {
	Find(ctx context.Context, id int64) (*CommonShard, error)
	FindOrNil(ctx context.Context, id int64) (*CommonShard, error)
	FindByShardKey(ctx context.Context, shardKey string) (*CommonShard, error)
	FinOrNilByShardKey(ctx context.Context, shardKey string) (*CommonShard, error)
	FindList(ctx context.Context) (CommonShards, error)
	FindListByShardKey(ctx context.Context, shardKey string) (CommonShards, error)
	Create(ctx context.Context, tx *gorm.DB, m *CommonShard) (*CommonShard, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms CommonShards) (CommonShards, error)
	Update(ctx context.Context, tx *gorm.DB, m *CommonShard) (*CommonShard, error)
	Delete(ctx context.Context, tx *gorm.DB, m *CommonShard) error
}
