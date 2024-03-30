// Package masterRanking ランキング
//
//go:generate mockgen -source=./master_ranking_mysql_repository.gen.go -destination=./master_ranking_mysql_repository_mock.gen.go -package=masterRanking
package masterRanking

import (
	"context"

	"gorm.io/gorm"
)

type MasterRankingMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterRanking, error)
	FindOrNil(ctx context.Context, id int64) (*MasterRanking, error)
	FindByMasterRankingEventId(ctx context.Context, masterRankingEventId int64) (*MasterRanking, error)
	FindOrNilByMasterRankingEventId(ctx context.Context, masterRankingEventId int64) (*MasterRanking, error)
	FindList(ctx context.Context) (MasterRankings, error)
	FindListByMasterRankingEventId(ctx context.Context, masterRankingEventId int64) (MasterRankings, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterRanking) (*MasterRanking, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterRankings) (MasterRankings, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterRanking) (*MasterRanking, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterRanking) error
}
