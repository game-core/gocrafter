// Package masterAction アクション
//
//go:generate mockgen -source=./master_action_repository.gen.go -destination=./master_action_repository_mock.gen.go -package=masterAction
package masterAction

import (
	"context"

	"gorm.io/gorm"
)

type MasterActionRepository interface {
	Find(ctx context.Context, id int64) (*MasterAction, error)
	FindOrNil(ctx context.Context, id int64) (*MasterAction, error)
	FindByName(ctx context.Context, name string) (*MasterAction, error)
	FindByAnyId(ctx context.Context, anyId int64) (*MasterAction, error)
	FinOrNilByName(ctx context.Context, name string) (*MasterAction, error)
	FinOrNilByAnyId(ctx context.Context, anyId int64) (*MasterAction, error)
	FindList(ctx context.Context) (MasterActions, error)
	FindListByName(ctx context.Context, name string) (MasterActions, error)
	FindListByAnyId(ctx context.Context, anyId int64) (MasterActions, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterActions) (MasterActions, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterAction) error
}
