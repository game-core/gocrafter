// Package masterLoginBonus ログインボーナス
package masterLoginBonus

import (
	"context"

	"gorm.io/gorm"
)

type MasterLoginBonusRepository interface {
	Find(ctx context.Context, id int64) (*MasterLoginBonus, error)
	FindOrNil(ctx context.Context, id int64) (*MasterLoginBonus, error)
	FindByMasterEventId(ctx context.Context, masterEventId int64) (*MasterLoginBonus, error)
	FinOrNilByMasterEventId(ctx context.Context, masterEventId int64) (*MasterLoginBonus, error)
	FindList(ctx context.Context) (MasterLoginBonuses, error)
	FindListByMasterEventId(ctx context.Context, masterEventId int64) (MasterLoginBonuses, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) (MasterLoginBonuses, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) error
}
