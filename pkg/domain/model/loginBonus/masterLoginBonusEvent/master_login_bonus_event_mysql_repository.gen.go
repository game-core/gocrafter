// Package masterLoginBonusEvent ログインボーナスイベント
//
//go:generate mockgen -source=./master_login_bonus_event_mysql_repository.gen.go -destination=./master_login_bonus_event_mysql_repository_mock.gen.go -package=masterLoginBonusEvent
package masterLoginBonusEvent

import (
	"context"

	"gorm.io/gorm"
)

type MasterLoginBonusEventMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterLoginBonusEvent, error)
	FindOrNil(ctx context.Context, id int64) (*MasterLoginBonusEvent, error)
	FindList(ctx context.Context) (MasterLoginBonusEvents, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusEvent) (*MasterLoginBonusEvent, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonusEvents) (MasterLoginBonusEvents, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusEvent) (*MasterLoginBonusEvent, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusEvent) error
}
