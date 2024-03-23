// Package masterIdleBonusItem 放置ボーナスアイテム
//
//go:generate mockgen -source=./master_idle_bonus_item_mysql_repository.gen.go -destination=./master_idle_bonus_item_mysql_repository_mock.gen.go -package=masterIdleBonusItem
package masterIdleBonusItem

import (
	"context"

	"gorm.io/gorm"
)

type MasterIdleBonusItemMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterIdleBonusItem, error)
	FindOrNil(ctx context.Context, id int64) (*MasterIdleBonusItem, error)
	FindByMasterIdleBonusScheduleId(ctx context.Context, masterIdleBonusScheduleId int64) (*MasterIdleBonusItem, error)
	FindByMasterItemId(ctx context.Context, masterItemId int64) (*MasterIdleBonusItem, error)
	FindByMasterIdleBonusScheduleIdAndMasterItemId(ctx context.Context, masterIdleBonusScheduleId int64, masterItemId int64) (*MasterIdleBonusItem, error)
	FindOrNilByMasterIdleBonusScheduleId(ctx context.Context, masterIdleBonusScheduleId int64) (*MasterIdleBonusItem, error)
	FindOrNilByMasterItemId(ctx context.Context, masterItemId int64) (*MasterIdleBonusItem, error)
	FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId(ctx context.Context, masterIdleBonusScheduleId int64, masterItemId int64) (*MasterIdleBonusItem, error)
	FindList(ctx context.Context) (MasterIdleBonusItems, error)
	FindListByMasterIdleBonusScheduleId(ctx context.Context, masterIdleBonusScheduleId int64) (MasterIdleBonusItems, error)
	FindListByMasterItemId(ctx context.Context, masterItemId int64) (MasterIdleBonusItems, error)
	FindListByMasterIdleBonusScheduleIdAndMasterItemId(ctx context.Context, masterIdleBonusScheduleId int64, masterItemId int64) (MasterIdleBonusItems, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusItem) (*MasterIdleBonusItem, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterIdleBonusItems) (MasterIdleBonusItems, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusItem) (*MasterIdleBonusItem, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusItem) error
}
