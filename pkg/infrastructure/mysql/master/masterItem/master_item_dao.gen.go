// Package masterItem アイテム
package masterItem

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/item/masterItem"
)

type masterItemDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterItemDao(conn *database.MysqlHandler) masterItem.MasterItemRepository {
	return &masterItemDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterItemDao) Find(ctx context.Context, id int64) (*masterItem.MasterItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_item", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterItem.MasterItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterItem()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterItem.SetMasterItem(t.Id, t.Name, t.ResourceType, t.RarityType, t.Content)
	s.Cache.Set(cashes.CreateCacheKey("master_item", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterItemDao) FindOrNil(ctx context.Context, id int64) (*masterItem.MasterItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_item", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterItem.MasterItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterItem()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterItem.SetMasterItem(t.Id, t.Name, t.ResourceType, t.RarityType, t.Content)
	s.Cache.Set(cashes.CreateCacheKey("master_item", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterItemDao) FindByName(ctx context.Context, name string) (*masterItem.MasterItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_item", "FindByName", fmt.Sprintf("%s_", name)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterItem.MasterItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterItem()
	res := s.ReadConn.WithContext(ctx).Where("name = ?", name).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterItem.SetMasterItem(t.Id, t.Name, t.ResourceType, t.RarityType, t.Content)
	s.Cache.Set(cashes.CreateCacheKey("master_item", "FindByName", fmt.Sprintf("%s_", name)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterItemDao) FindOrNilByName(ctx context.Context, name string) (*masterItem.MasterItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_item", "FindOrNilByName", fmt.Sprintf("%s_", name)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterItem.MasterItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterItem()
	res := s.ReadConn.WithContext(ctx).Where("name = ?", name).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterItem.SetMasterItem(t.Id, t.Name, t.ResourceType, t.RarityType, t.Content)
	s.Cache.Set(cashes.CreateCacheKey("master_item", "FindOrNilByName", fmt.Sprintf("%s_", name)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterItemDao) FindList(ctx context.Context) (masterItem.MasterItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_item", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterItem.MasterItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterItems()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterItem.NewMasterItems()
	for _, t := range ts {
		ms = append(ms, masterItem.SetMasterItem(t.Id, t.Name, t.ResourceType, t.RarityType, t.Content))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_item", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterItemDao) FindListByName(ctx context.Context, name string) (masterItem.MasterItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_item", "FindListByName", fmt.Sprintf("%s_", name)))
	if found {
		if cachedEntity, ok := cachedResult.(masterItem.MasterItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterItems()
	res := s.ReadConn.WithContext(ctx).Where("name = ?", name).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterItem.NewMasterItems()
	for _, t := range ts {
		ms = append(ms, masterItem.SetMasterItem(t.Id, t.Name, t.ResourceType, t.RarityType, t.Content))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_item", "FindListByName", fmt.Sprintf("%s_", name)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterItemDao) Create(ctx context.Context, tx *gorm.DB, m *masterItem.MasterItem) (*masterItem.MasterItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterItem{
		Id:           m.Id,
		Name:         m.Name,
		ResourceType: m.ResourceType,
		RarityType:   m.RarityType,
		Content:      m.Content,
	}
	res := conn.Model(NewMasterItem()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterItem.SetMasterItem(t.Id, t.Name, t.ResourceType, t.RarityType, t.Content), nil
}

func (s *masterItemDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterItem.MasterItems) (masterItem.MasterItems, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterItems()
	for _, m := range ms {
		t := &MasterItem{
			Id:           m.Id,
			Name:         m.Name,
			ResourceType: m.ResourceType,
			RarityType:   m.RarityType,
			Content:      m.Content,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterItem()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterItemDao) Update(ctx context.Context, tx *gorm.DB, m *masterItem.MasterItem) (*masterItem.MasterItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterItem{
		Id:           m.Id,
		Name:         m.Name,
		ResourceType: m.ResourceType,
		RarityType:   m.RarityType,
		Content:      m.Content,
	}
	res := conn.Model(NewMasterItem()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterItem.SetMasterItem(t.Id, t.Name, t.ResourceType, t.RarityType, t.Content), nil
}

func (s *masterItemDao) Delete(ctx context.Context, tx *gorm.DB, m *masterItem.MasterItem) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterItem()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
