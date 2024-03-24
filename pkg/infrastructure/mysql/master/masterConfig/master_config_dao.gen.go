// Package masterConfig 設定
package masterConfig

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/config/masterConfig"
)

type masterConfigDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterConfigDao(conn *database.MysqlHandler) masterConfig.MasterConfigMysqlRepository {
	return &masterConfigDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterConfigDao) Find(ctx context.Context, id int64) (*masterConfig.MasterConfig, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_config", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterConfig.MasterConfig); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterConfig()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterConfig.SetMasterConfig(t.Id, t.Name, t.ConfigType, t.Value)
	s.Cache.Set(cashes.CreateCacheKey("master_config", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterConfigDao) FindOrNil(ctx context.Context, id int64) (*masterConfig.MasterConfig, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_config", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterConfig.MasterConfig); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterConfig()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterConfig.SetMasterConfig(t.Id, t.Name, t.ConfigType, t.Value)
	s.Cache.Set(cashes.CreateCacheKey("master_config", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterConfigDao) FindByConfigType(ctx context.Context, configType enum.ConfigType) (*masterConfig.MasterConfig, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_config", "FindByConfigType", fmt.Sprintf("%d_", configType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterConfig.MasterConfig); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterConfig()
	res := s.ReadMysqlConn.WithContext(ctx).Where("config_type = ?", configType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterConfig.SetMasterConfig(t.Id, t.Name, t.ConfigType, t.Value)
	s.Cache.Set(cashes.CreateCacheKey("master_config", "FindByConfigType", fmt.Sprintf("%d_", configType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterConfigDao) FindOrNilByConfigType(ctx context.Context, configType enum.ConfigType) (*masterConfig.MasterConfig, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_config", "FindOrNilByConfigType", fmt.Sprintf("%d_", configType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterConfig.MasterConfig); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterConfig()
	res := s.ReadMysqlConn.WithContext(ctx).Where("config_type = ?", configType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterConfig.SetMasterConfig(t.Id, t.Name, t.ConfigType, t.Value)
	s.Cache.Set(cashes.CreateCacheKey("master_config", "FindOrNilByConfigType", fmt.Sprintf("%d_", configType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterConfigDao) FindList(ctx context.Context) (masterConfig.MasterConfigs, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_config", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterConfig.MasterConfigs); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterConfigs()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterConfig.NewMasterConfigs()
	for _, t := range ts {
		ms = append(ms, masterConfig.SetMasterConfig(t.Id, t.Name, t.ConfigType, t.Value))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_config", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterConfigDao) FindListByConfigType(ctx context.Context, configType enum.ConfigType) (masterConfig.MasterConfigs, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_config", "FindListByConfigType", fmt.Sprintf("%d_", configType)))
	if found {
		if cachedEntity, ok := cachedResult.(masterConfig.MasterConfigs); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterConfigs()
	res := s.ReadMysqlConn.WithContext(ctx).Where("config_type = ?", configType).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterConfig.NewMasterConfigs()
	for _, t := range ts {
		ms = append(ms, masterConfig.SetMasterConfig(t.Id, t.Name, t.ConfigType, t.Value))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_config", "FindListByConfigType", fmt.Sprintf("%d_", configType)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterConfigDao) Create(ctx context.Context, tx *gorm.DB, m *masterConfig.MasterConfig) (*masterConfig.MasterConfig, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterConfig{
		Id:         m.Id,
		Name:       m.Name,
		ConfigType: m.ConfigType,
		Value:      m.Value,
	}
	res := conn.Model(NewMasterConfig()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterConfig.SetMasterConfig(t.Id, t.Name, t.ConfigType, t.Value), nil
}

func (s *masterConfigDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterConfig.MasterConfigs) (masterConfig.MasterConfigs, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterConfigs()
	for _, m := range ms {
		t := &MasterConfig{
			Id:         m.Id,
			Name:       m.Name,
			ConfigType: m.ConfigType,
			Value:      m.Value,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterConfig()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterConfigDao) Update(ctx context.Context, tx *gorm.DB, m *masterConfig.MasterConfig) (*masterConfig.MasterConfig, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterConfig{
		Id:         m.Id,
		Name:       m.Name,
		ConfigType: m.ConfigType,
		Value:      m.Value,
	}
	res := conn.Model(NewMasterConfig()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterConfig.SetMasterConfig(t.Id, t.Name, t.ConfigType, t.Value), nil
}

func (s *masterConfigDao) Delete(ctx context.Context, tx *gorm.DB, m *masterConfig.MasterConfig) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterConfig()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterConfig())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
