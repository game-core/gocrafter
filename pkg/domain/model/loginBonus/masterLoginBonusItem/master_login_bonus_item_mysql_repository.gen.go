// Package masterLoginBonusItem ログインボーナスアイテム
//
//go:generate mockgen -source=./master_login_bonus_item_mysql_repository.gen.go -destination=./master_login_bonus_item_mysql_repository_mock.gen.go -package=masterLoginBonusItem
package masterLoginBonusItem

import (
	"context"

	"gorm.io/gorm"
)

type MasterLoginBonusItemMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterLoginBonusItem, error)
	FindOrNil(ctx context.Context, id int64) (*MasterLoginBonusItem, error)
	FindByMasterLoginBonusScheduleId(ctx context.Context, masterLoginBonusScheduleId int64) (*MasterLoginBonusItem, error)
	FindByMasterItemId(ctx context.Context, masterItemId int64) (*MasterLoginBonusItem, error)
	FindByMasterLoginBonusScheduleIdAndMasterItemId(ctx context.Context, masterLoginBonusScheduleId int64, masterItemId int64) (*MasterLoginBonusItem, error)
	FindOrNilByMasterLoginBonusScheduleId(ctx context.Context, masterLoginBonusScheduleId int64) (*MasterLoginBonusItem, error)
	FindOrNilByMasterItemId(ctx context.Context, masterItemId int64) (*MasterLoginBonusItem, error)
	FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId(ctx context.Context, masterLoginBonusScheduleId int64, masterItemId int64) (*MasterLoginBonusItem, error)
	FindList(ctx context.Context) (MasterLoginBonusItems, error)
	FindListByMasterLoginBonusScheduleId(ctx context.Context, masterLoginBonusScheduleId int64) (MasterLoginBonusItems, error)
	FindListByMasterItemId(ctx context.Context, masterItemId int64) (MasterLoginBonusItems, error)
	FindListByMasterLoginBonusScheduleIdAndMasterItemId(ctx context.Context, masterLoginBonusScheduleId int64, masterItemId int64) (MasterLoginBonusItems, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusItem) (*MasterLoginBonusItem, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonusItems) (MasterLoginBonusItems, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusItem) (*MasterLoginBonusItem, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusItem) error
}
