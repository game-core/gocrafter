// Package masterRankingEvent ランキングイベント
//
//go:generate mockgen -source=./master_ranking_event_mysql_repository.gen.go -destination=./master_ranking_event_mysql_repository_mock.gen.go -package=masterRankingEvent
package masterRankingEvent

import (
	"context"

	"gorm.io/gorm"
)

type MasterRankingEventMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterRankingEvent, error)
	FindOrNil(ctx context.Context, id int64) (*MasterRankingEvent, error)
	FindList(ctx context.Context) (MasterRankingEvents, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterRankingEvent) (*MasterRankingEvent, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterRankingEvents) (MasterRankingEvents, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterRankingEvent) (*MasterRankingEvent, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterRankingEvent) error
}
