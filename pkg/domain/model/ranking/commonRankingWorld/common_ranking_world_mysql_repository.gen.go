// Package commonRankingWorld ワールドランキング
//
//go:generate mockgen -source=./common_ranking_world_mysql_repository.gen.go -destination=./common_ranking_world_mysql_repository_mock.gen.go -package=commonRankingWorld
package commonRankingWorld

import (
	"context"

	"gorm.io/gorm"
)

type CommonRankingWorldMysqlRepository interface {
	Find(ctx context.Context, masterRankingId int64, userId string) (*CommonRankingWorld, error)
	FindOrNil(ctx context.Context, masterRankingId int64, userId string) (*CommonRankingWorld, error)
	FindByMasterRankingId(ctx context.Context, masterRankingId int64) (*CommonRankingWorld, error)
	FindOrNilByMasterRankingId(ctx context.Context, masterRankingId int64) (*CommonRankingWorld, error)
	FindList(ctx context.Context) (CommonRankingWorlds, error)
	FindListByMasterRankingId(ctx context.Context, masterRankingId int64) (CommonRankingWorlds, error)
	Create(ctx context.Context, tx *gorm.DB, m *CommonRankingWorld) (*CommonRankingWorld, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms CommonRankingWorlds) (CommonRankingWorlds, error)
	Update(ctx context.Context, tx *gorm.DB, m *CommonRankingWorld) (*CommonRankingWorld, error)
	Delete(ctx context.Context, tx *gorm.DB, m *CommonRankingWorld) error
}
