// Package commonShard シャード管理
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
