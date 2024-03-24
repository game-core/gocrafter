// Package masterConfig 設定
//
//go:generate mockgen -source=./master_config_mysql_repository.gen.go -destination=./master_config_mysql_repository_mock.gen.go -package=masterConfig
package masterConfig

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterConfigMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterConfig, error)
	FindOrNil(ctx context.Context, id int64) (*MasterConfig, error)
	FindByConfigType(ctx context.Context, configType enum.ConfigType) (*MasterConfig, error)
	FindOrNilByConfigType(ctx context.Context, configType enum.ConfigType) (*MasterConfig, error)
	FindList(ctx context.Context) (MasterConfigs, error)
	FindListByConfigType(ctx context.Context, configType enum.ConfigType) (MasterConfigs, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterConfig) (*MasterConfig, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterConfigs) (MasterConfigs, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterConfig) (*MasterConfig, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterConfig) error
}
