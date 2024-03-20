// Package masterActionRun 実行されるアクション
//
//go:generate mockgen -source=./master_action_run_repository.gen.go -destination=./master_action_run_repository_mock.gen.go -package=masterActionRun
package masterActionRun

import (
	"context"

	"gorm.io/gorm"
)

type MasterActionRunRepository interface {
	Find(ctx context.Context, id int64) (*MasterActionRun, error)
	FindOrNil(ctx context.Context, id int64) (*MasterActionRun, error)
	FindByName(ctx context.Context, name string) (*MasterActionRun, error)
	FindByActionId(ctx context.Context, actionId int64) (*MasterActionRun, error)
	FindOrNilByName(ctx context.Context, name string) (*MasterActionRun, error)
	FindOrNilByActionId(ctx context.Context, actionId int64) (*MasterActionRun, error)
	FindList(ctx context.Context) (MasterActionRuns, error)
	FindListByName(ctx context.Context, name string) (MasterActionRuns, error)
	FindListByActionId(ctx context.Context, actionId int64) (MasterActionRuns, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterActionRun) (*MasterActionRun, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterActionRuns) (MasterActionRuns, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterActionRun) (*MasterActionRun, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterActionRun) error
}
