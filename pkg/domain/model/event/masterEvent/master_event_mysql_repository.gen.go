// Package masterEvent イベント
//
//go:generate mockgen -source=./master_event_mysql_repository.gen.go -destination=./master_event_mysql_repository_mock.gen.go -package=masterEvent
package masterEvent

import (
	"context"

	"gorm.io/gorm"
)

type MasterEventMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterEvent, error)
	FindOrNil(ctx context.Context, id int64) (*MasterEvent, error)
	FindList(ctx context.Context) (MasterEvents, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterEvent) (*MasterEvent, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterEvents) (MasterEvents, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterEvent) (*MasterEvent, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterEvent) error
}
