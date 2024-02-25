// Package masterItem アイテム
//
//go:generate mockgen -source=./master_item_repository.gen.go -destination=./master_item_repository_mock.gen.go -package=masterItem
package masterItem

import (
	context "context"

	"gorm.io/gorm"
)

type MasterItemRepository interface {
	Find(ctx context.Context, id int64) (*MasterItem, error)
	FindOrNil(ctx context.Context, id int64) (*MasterItem, error)
	FindByName(ctx context.Context, name string) (*MasterItem, error)
	FinOrNilByName(ctx context.Context, name string) (*MasterItem, error)
	FindList(ctx context.Context) (MasterItems, error)
	FindListByName(ctx context.Context, name string) (MasterItems, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterItem) (*MasterItem, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterItems) (MasterItems, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterItem) (*MasterItem, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterItem) error
}
