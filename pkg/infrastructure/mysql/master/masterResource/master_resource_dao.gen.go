// Package masterResource レアリティ
package masterResource

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/resource/masterResource"
)

type masterResourceDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterResourceDao(conn *database.MysqlHandler) masterResource.MasterResourceRepository {
	return &masterResourceDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterResourceDao) Find(ctx context.Context, id int64) (*masterResource.MasterResource, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterResource.MasterResource); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterResource()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterResource.SetMasterResource(t.Id, t.Name, t.ResourceType)
	s.Cache.Set(cashes.CreateCacheKey("master_resource", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterResourceDao) FindOrNil(ctx context.Context, id int64) (*masterResource.MasterResource, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterResource.MasterResource); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterResource()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterResource.SetMasterResource(t.Id, t.Name, t.ResourceType)
	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterResourceDao) FindByResourceType(ctx context.Context, resourceType enum.ResourceType) (*masterResource.MasterResource, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindByResourceType", fmt.Sprintf("%d_", resourceType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterResource.MasterResource); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterResource()
	res := s.ReadMysqlConn.WithContext(ctx).Where("resource_type = ?", resourceType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterResource.SetMasterResource(t.Id, t.Name, t.ResourceType)
	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindByResourceType", fmt.Sprintf("%d_", resourceType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterResourceDao) FindOrNilByResourceType(ctx context.Context, resourceType enum.ResourceType) (*masterResource.MasterResource, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindOrNilByResourceType", fmt.Sprintf("%d_", resourceType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterResource.MasterResource); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterResource()
	res := s.ReadMysqlConn.WithContext(ctx).Where("resource_type = ?", resourceType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterResource.SetMasterResource(t.Id, t.Name, t.ResourceType)
	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindOrNilByResourceType", fmt.Sprintf("%d_", resourceType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterResourceDao) FindList(ctx context.Context) (masterResource.MasterResources, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterResource.MasterResources); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterResources()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterResource.NewMasterResources()
	for _, t := range ts {
		ms = append(ms, masterResource.SetMasterResource(t.Id, t.Name, t.ResourceType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterResourceDao) FindListByResourceType(ctx context.Context, resourceType enum.ResourceType) (masterResource.MasterResources, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindListByResourceType", fmt.Sprintf("%d_", resourceType)))
	if found {
		if cachedEntity, ok := cachedResult.(masterResource.MasterResources); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterResources()
	res := s.ReadMysqlConn.WithContext(ctx).Where("resource_type = ?", resourceType).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterResource.NewMasterResources()
	for _, t := range ts {
		ms = append(ms, masterResource.SetMasterResource(t.Id, t.Name, t.ResourceType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindListByResourceType", fmt.Sprintf("%d_", resourceType)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterResourceDao) Create(ctx context.Context, tx *gorm.DB, m *masterResource.MasterResource) (*masterResource.MasterResource, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterResource{
		Id:           m.Id,
		Name:         m.Name,
		ResourceType: m.ResourceType,
	}
	res := conn.Model(NewMasterResource()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterResource.SetMasterResource(t.Id, t.Name, t.ResourceType), nil
}

func (s *masterResourceDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterResource.MasterResources) (masterResource.MasterResources, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterResources()
	for _, m := range ms {
		t := &MasterResource{
			Id:           m.Id,
			Name:         m.Name,
			ResourceType: m.ResourceType,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterResource()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterResourceDao) Update(ctx context.Context, tx *gorm.DB, m *masterResource.MasterResource) (*masterResource.MasterResource, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterResource{
		Id:           m.Id,
		Name:         m.Name,
		ResourceType: m.ResourceType,
	}
	res := conn.Model(NewMasterResource()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterResource.SetMasterResource(t.Id, t.Name, t.ResourceType), nil
}

func (s *masterResourceDao) Delete(ctx context.Context, tx *gorm.DB, m *masterResource.MasterResource) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterResource()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterResource())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
