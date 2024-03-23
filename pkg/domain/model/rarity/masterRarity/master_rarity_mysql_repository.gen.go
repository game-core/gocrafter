// Package masterRarity レアリティ
//
//go:generate mockgen -source=./master_rarity_mysql_repository.gen.go -destination=./master_rarity_mysql_repository_mock.gen.go -package=masterRarity
package masterRarity

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterRarityMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterRarity, error)
	FindOrNil(ctx context.Context, id int64) (*MasterRarity, error)
	FindByRarityType(ctx context.Context, rarityType enum.RarityType) (*MasterRarity, error)
	FindOrNilByRarityType(ctx context.Context, rarityType enum.RarityType) (*MasterRarity, error)
	FindList(ctx context.Context) (MasterRarities, error)
	FindListByRarityType(ctx context.Context, rarityType enum.RarityType) (MasterRarities, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterRarity) (*MasterRarity, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterRarities) (MasterRarities, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterRarity) (*MasterRarity, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterRarity) error
}
