// Package masterIdleBonusEvent 放置ボーナスイベント
//
//go:generate mockgen -source=./master_idle_bonus_event_repository.gen.go -destination=./master_idle_bonus_event_repository_mock.gen.go -package=masterIdleBonusEvent
package masterIdleBonusEvent

import (
	"context"

	"gorm.io/gorm"
)

type MasterIdleBonusEventRepository interface {
	Find(ctx context.Context, id int64) (*MasterIdleBonusEvent, error)
	FindOrNil(ctx context.Context, id int64) (*MasterIdleBonusEvent, error)
	FindList(ctx context.Context) (MasterIdleBonusEvents, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusEvent) (*MasterIdleBonusEvent, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterIdleBonusEvents) (MasterIdleBonusEvents, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusEvent) (*MasterIdleBonusEvent, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusEvent) error
}
