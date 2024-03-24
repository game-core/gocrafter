// Package masterConfig 設定
package masterConfig

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterConfigs []*MasterConfig

type MasterConfig struct {
	Id         int64
	Name       string
	ConfigType enum.ConfigType
	Value      string
}

func NewMasterConfig() *MasterConfig {
	return &MasterConfig{}
}

func NewMasterConfigs() MasterConfigs {
	return MasterConfigs{}
}

func SetMasterConfig(id int64, name string, configType enum.ConfigType, value string) *MasterConfig {
	return &MasterConfig{
		Id:         id,
		Name:       name,
		ConfigType: configType,
		Value:      value,
	}
}

func (t *MasterConfig) TableName() string {
	return "master_config"
}
