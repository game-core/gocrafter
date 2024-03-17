// Package masterActionStep アクションステップ
//
//go:generate mockgen -source=./master_action_step_repository.gen.go -destination=./master_action_step_repository_mock.gen.go -package=masterActionStep
package masterActionStep

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterActionStepRepository interface {
	Find(ctx context.Context, id int64) (*MasterActionStep, error)
	FindOrNil(ctx context.Context, id int64) (*MasterActionStep, error)
	FindByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*MasterActionStep, error)
	FinOrNilByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*MasterActionStep, error)
	FindList(ctx context.Context) (MasterActionSteps, error)
	FindListByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (MasterActionSteps, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterActionStep) (*MasterActionStep, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterActionSteps) (MasterActionSteps, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterActionStep) (*MasterActionStep, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterActionStep) error
}
