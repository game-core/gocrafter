// Package masterRarity レアリティ
package masterRarity

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/rarity/masterRarity"
)

type masterRarityDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterRarityDao(conn *database.MysqlHandler) masterRarity.MasterRarityRepository {
	return &masterRarityDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterRarityDao) Find(ctx context.Context, id int64) (*masterRarity.MasterRarity, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRarity.MasterRarity); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRarity()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterRarity.SetMasterRarity(t.Id, t.Name, t.RarityType)
	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRarityDao) FindOrNil(ctx context.Context, id int64) (*masterRarity.MasterRarity, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRarity.MasterRarity); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRarity()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterRarity.SetMasterRarity(t.Id, t.Name, t.RarityType)
	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRarityDao) FindByRarityType(ctx context.Context, rarityType enum.RarityType) (*masterRarity.MasterRarity, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindByRarityType", fmt.Sprintf("%d_", rarityType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRarity.MasterRarity); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRarity()
	res := s.ReadMysqlConn.WithContext(ctx).Where("rarity_type = ?", rarityType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterRarity.SetMasterRarity(t.Id, t.Name, t.RarityType)
	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindByRarityType", fmt.Sprintf("%d_", rarityType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRarityDao) FindOrNilByRarityType(ctx context.Context, rarityType enum.RarityType) (*masterRarity.MasterRarity, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindOrNilByRarityType", fmt.Sprintf("%d_", rarityType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRarity.MasterRarity); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRarity()
	res := s.ReadMysqlConn.WithContext(ctx).Where("rarity_type = ?", rarityType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterRarity.SetMasterRarity(t.Id, t.Name, t.RarityType)
	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindOrNilByRarityType", fmt.Sprintf("%d_", rarityType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRarityDao) FindList(ctx context.Context) (masterRarity.MasterRarities, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterRarity.MasterRarities); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterRarities()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterRarity.NewMasterRarities()
	for _, t := range ts {
		ms = append(ms, masterRarity.SetMasterRarity(t.Id, t.Name, t.RarityType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterRarityDao) FindListByRarityType(ctx context.Context, rarityType enum.RarityType) (masterRarity.MasterRarities, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindListByRarityType", fmt.Sprintf("%d_", rarityType)))
	if found {
		if cachedEntity, ok := cachedResult.(masterRarity.MasterRarities); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterRarities()
	res := s.ReadMysqlConn.WithContext(ctx).Where("rarity_type = ?", rarityType).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterRarity.NewMasterRarities()
	for _, t := range ts {
		ms = append(ms, masterRarity.SetMasterRarity(t.Id, t.Name, t.RarityType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindListByRarityType", fmt.Sprintf("%d_", rarityType)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterRarityDao) Create(ctx context.Context, tx *gorm.DB, m *masterRarity.MasterRarity) (*masterRarity.MasterRarity, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRarity{
		Id:         m.Id,
		Name:       m.Name,
		RarityType: m.RarityType,
	}
	res := conn.Model(NewMasterRarity()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRarity.SetMasterRarity(t.Id, t.Name, t.RarityType), nil
}

func (s *masterRarityDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterRarity.MasterRarities) (masterRarity.MasterRarities, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterRarities()
	for _, m := range ms {
		t := &MasterRarity{
			Id:         m.Id,
			Name:       m.Name,
			RarityType: m.RarityType,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterRarity()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterRarityDao) Update(ctx context.Context, tx *gorm.DB, m *masterRarity.MasterRarity) (*masterRarity.MasterRarity, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRarity{
		Id:         m.Id,
		Name:       m.Name,
		RarityType: m.RarityType,
	}
	res := conn.Model(NewMasterRarity()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRarity.SetMasterRarity(t.Id, t.Name, t.RarityType), nil
}

func (s *masterRarityDao) Delete(ctx context.Context, tx *gorm.DB, m *masterRarity.MasterRarity) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterRarity()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterRarity())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
