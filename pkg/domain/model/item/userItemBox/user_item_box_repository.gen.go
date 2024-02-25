// Package userItemBox ユーザーアイテムボックス
//
//go:generate mockgen -source=./user_item_box_repository.gen.go -destination=./user_item_box_repository_mock.gen.go -package=userItemBox
package userItemBox

import (
	"context"

	"gorm.io/gorm"
)

type UserItemBoxRepository interface {
	Find(ctx context.Context, userId string, masterItemId int64) (*UserItemBox, error)
	FindOrNil(ctx context.Context, userId string, masterItemId int64) (*UserItemBox, error)
	FindList(ctx context.Context, userId string) (UserItemBoxs, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserItemBox) (*UserItemBox, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserItemBoxs) (UserItemBoxs, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserItemBox) (*UserItemBox, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserItemBox) error
}
