// Package masterLoginBonus ログインボーナス
//
//go:generate mockgen -source=./master_login_bonus_repository.gen.go -destination=./master_login_bonus_repository_mock.gen.go -package=masterLoginBonus
package masterLoginBonus

import (
	"context"

	"gorm.io/gorm"
)

type MasterLoginBonusRepository interface {
	Find(ctx context.Context, id int64) (*MasterLoginBonus, error)
	FindOrNil(ctx context.Context, id int64) (*MasterLoginBonus, error)
	FindByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (*MasterLoginBonus, error)
	FinOrNilByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (*MasterLoginBonus, error)
	FindList(ctx context.Context) (MasterLoginBonuses, error)
	FindListByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (MasterLoginBonuses, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) (MasterLoginBonuses, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) error
}
