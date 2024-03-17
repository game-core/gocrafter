// Package masterActionRun 実行されるアクション
package masterActionRun

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionRun"
)

type masterActionRunDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterActionRunDao(conn *database.SqlHandler) masterActionRun.MasterActionRunRepository {
	return &masterActionRunDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterActionRunDao) Find(ctx context.Context, id int64) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunDao) FindOrNil(ctx context.Context, id int64) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunDao) FindByName(ctx context.Context, name string) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindByName", fmt.Sprintf("%s_", name)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadConn.WithContext(ctx).Where("name = ?", name).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindByName", fmt.Sprintf("%s_", name)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunDao) FindByActionId(ctx context.Context, actionId int64) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindByActionId", fmt.Sprintf("%d_", actionId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadConn.WithContext(ctx).Where("action_id = ?", actionId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindByActionId", fmt.Sprintf("%d_", actionId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunDao) FinOrNilByName(ctx context.Context, name string) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindOrNilByName", fmt.Sprintf("%s_", name)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadConn.WithContext(ctx).Where("name = ?", name).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindOrNilByName", fmt.Sprintf("%s_", name)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunDao) FinOrNilByActionId(ctx context.Context, actionId int64) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindOrNilByActionId", fmt.Sprintf("%d_", actionId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadConn.WithContext(ctx).Where("action_id = ?", actionId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindOrNilByActionId", fmt.Sprintf("%d_", actionId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunDao) FindList(ctx context.Context) (masterActionRun.MasterActionRuns, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionRun.MasterActionRuns); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionRuns()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionRun.NewMasterActionRuns()
	for _, t := range ts {
		ms = append(ms, masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionRunDao) FindListByName(ctx context.Context, name string) (masterActionRun.MasterActionRuns, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindListByName", fmt.Sprintf("%s_", name)))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionRun.MasterActionRuns); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionRuns()
	res := s.ReadConn.WithContext(ctx).Where("name = ?", name).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionRun.NewMasterActionRuns()
	for _, t := range ts {
		ms = append(ms, masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindListByName", fmt.Sprintf("%s_", name)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionRunDao) FindListByActionId(ctx context.Context, actionId int64) (masterActionRun.MasterActionRuns, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindListByActionId", fmt.Sprintf("%d_", actionId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionRun.MasterActionRuns); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionRuns()
	res := s.ReadConn.WithContext(ctx).Where("action_id = ?", actionId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionRun.NewMasterActionRuns()
	for _, t := range ts {
		ms = append(ms, masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindListByActionId", fmt.Sprintf("%d_", actionId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionRunDao) Create(ctx context.Context, tx *gorm.DB, m *masterActionRun.MasterActionRun) (*masterActionRun.MasterActionRun, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterActionRun{
		Id:       m.Id,
		Name:     m.Name,
		ActionId: m.ActionId,
	}
	res := conn.Model(NewMasterActionRun()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId), nil
}

func (s *masterActionRunDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterActionRun.MasterActionRuns) (masterActionRun.MasterActionRuns, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterActionRuns()
	for _, m := range ms {
		t := &MasterActionRun{
			Id:       m.Id,
			Name:     m.Name,
			ActionId: m.ActionId,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterActionRun()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionRunDao) Update(ctx context.Context, tx *gorm.DB, m *masterActionRun.MasterActionRun) (*masterActionRun.MasterActionRun, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterActionRun{
		Id:       m.Id,
		Name:     m.Name,
		ActionId: m.ActionId,
	}
	res := conn.Model(NewMasterActionRun()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionRun.SetMasterActionRun(t.Id, t.Name, t.ActionId), nil
}

func (s *masterActionRunDao) Delete(ctx context.Context, tx *gorm.DB, m *masterActionRun.MasterActionRun) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterActionRun()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterActionRun())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
