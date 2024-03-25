//go:generate mockgen -source=./config_service.go -destination=./config_service_mock.gen.go -package=config
package config

import (
	"context"

	"github.com/game-core/gocrafter/internal/changes"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/config/masterConfig"
)

type ConfigService interface {
	GetAll(cxt context.Context) (masterConfig.MasterConfigs, error)
	Get(cxt context.Context, configType enum.ConfigType) (*masterConfig.MasterConfig, error)
	GetInt32(cxt context.Context, configType enum.ConfigType) (int32, error)
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

// Get 設定を取得する
func (s *configService) Get(cxt context.Context, configType enum.ConfigType) (*masterConfig.MasterConfig, error) {
	result, err := s.masterConfigMysqlRepository.FindByConfigType(cxt, configType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterConfigMysqlRepository.FindByConfigType", err)
	}

	return result, nil
}

// GetInt32 設定を取得する
func (s *configService) GetInt32(cxt context.Context, configType enum.ConfigType) (int32, error) {
	masterConfigModel, err := s.masterConfigMysqlRepository.FindByConfigType(cxt, configType)
	if err != nil {
		return 0, errors.NewMethodError("s.masterConfigMysqlRepository.FindByConfigType", err)
	}

	value, err := changes.StringToInt32(masterConfigModel.Value)
	if err != nil {
		return 0, errors.NewMethodError("changes.StringToInt32", err)
	}

	return value, nil
}
