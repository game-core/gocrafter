//go:generate mockgen -source=./config_service.go -destination=./config_service_mock.gen.go -package=config
package config

import (
	"context"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/config/masterConfig"
)

type ConfigService interface {
	GetAll(cxt context.Context) (masterConfig.MasterConfigs, error)
	GetByConfigType(cxt context.Context, configType enum.ConfigType) (*masterConfig.MasterConfig, error)
}

type configService struct {
	masterConfigMysqlRepository masterConfig.MasterConfigMysqlRepository
}

func NewConfigService(
	masterConfigMysqlRepository masterConfig.MasterConfigMysqlRepository,
) ConfigService {
	return &configService{
		masterConfigMysqlRepository: masterConfigMysqlRepository,
	}
}

// GetAll 設定一覧を取得する
func (s *configService) GetAll(cxt context.Context) (masterConfig.MasterConfigs, error) {
	results, err := s.masterConfigMysqlRepository.FindList(cxt)
	if err != nil {
		return nil, errors.NewMethodError("s.masterConfigMysqlRepository.FindList", err)
	}

	return results, nil
}

// GetByConfigType 設定を取得する
func (s *configService) GetByConfigType(cxt context.Context, configType enum.ConfigType) (*masterConfig.MasterConfig, error) {
	result, err := s.masterConfigMysqlRepository.FindByConfigType(cxt, configType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterConfigMysqlRepository.FindByConfigType", err)
	}

	return result, nil
}
