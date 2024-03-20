// Package masterIdleBonus 放置ボーナス
//
//go:generate mockgen -source=./master_idle_bonus_repository.gen.go -destination=./master_idle_bonus_repository_mock.gen.go -package=masterIdleBonus
package masterIdleBonus

import (
	"context"

	"gorm.io/gorm"
)

type MasterIdleBonusRepository interface {
	Find(ctx context.Context, id int64) (*MasterIdleBonus, error)
	FindOrNil(ctx context.Context, id int64) (*MasterIdleBonus, error)
	FindByMasterIdleBonusEventId(ctx context.Context, masterIdleBonusEventId int64) (*MasterIdleBonus, error)
	FindOrNilByMasterIdleBonusEventId(ctx context.Context, masterIdleBonusEventId int64) (*MasterIdleBonus, error)
	FindList(ctx context.Context) (MasterIdleBonuses, error)
	FindListByMasterIdleBonusEventId(ctx context.Context, masterIdleBonusEventId int64) (MasterIdleBonuses, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterIdleBonus) (*MasterIdleBonus, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterIdleBonuses) (MasterIdleBonuses, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterIdleBonus) (*MasterIdleBonus, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterIdleBonus) error
}
