// Package masterActionTrigger アクショントリガー
//
//go:generate mockgen -source=./master_action_trigger_repository.gen.go -destination=./master_action_trigger_repository_mock.gen.go -package=masterActionTrigger
package masterActionTrigger

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterActionTriggerRepository interface {
	Find(ctx context.Context, id int64) (*MasterActionTrigger, error)
	FindOrNil(ctx context.Context, id int64) (*MasterActionTrigger, error)
	FindByActionTriggerType(ctx context.Context, actionTriggerType enum.ActionTriggerType) (*MasterActionTrigger, error)
	FinOrNilByActionTriggerType(ctx context.Context, actionTriggerType enum.ActionTriggerType) (*MasterActionTrigger, error)
	FindList(ctx context.Context) (MasterActionTriggers, error)
	FindListByActionTriggerType(ctx context.Context, actionTriggerType enum.ActionTriggerType) (MasterActionTriggers, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterActionTrigger) (*MasterActionTrigger, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterActionTriggers) (MasterActionTriggers, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterActionTrigger) (*MasterActionTrigger, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterActionTrigger) error
}
