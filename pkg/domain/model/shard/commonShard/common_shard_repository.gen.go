// Package commonShard シャード管理
//
//go:generate mockgen -source=./common_shard_repository.gen.go -destination=./common_shard_repository_mock.gen.go -package=commonShard
package commonShard

import (
	"context"

	"gorm.io/gorm"
)

type CommonShardRepository interface {
	FindList(ctx context.Context) (CommonShards, error)
	Create(ctx context.Context, tx *gorm.DB, m *CommonShard) (*CommonShard, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms CommonShards) (CommonShards, error)
}
