// Package masterActionTrigger アクショントリガー
package masterActionTrigger

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionTrigger"
)

type masterActionTriggerDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterActionTriggerDao(conn *database.SqlHandler) masterActionTrigger.MasterActionTriggerRepository {
	return &masterActionTriggerDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterActionTriggerDao) Find(ctx context.Context, id int64) (*masterActionTrigger.MasterActionTrigger, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionTrigger.MasterActionTrigger); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionTrigger()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionTrigger.SetMasterActionTrigger(t.Id, t.Name, t.ActionTriggerType)
	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionTriggerDao) FindOrNil(ctx context.Context, id int64) (*masterActionTrigger.MasterActionTrigger, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionTrigger.MasterActionTrigger); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionTrigger()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionTrigger.SetMasterActionTrigger(t.Id, t.Name, t.ActionTriggerType)
	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionTriggerDao) FindByActionTriggerType(ctx context.Context, actionTriggerType enum.ActionTriggerType) (*masterActionTrigger.MasterActionTrigger, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindByActionTriggerType", fmt.Sprintf("%d_", actionTriggerType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionTrigger.MasterActionTrigger); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionTrigger()
	res := s.ReadConn.WithContext(ctx).Where("action_trigger_type = ?", actionTriggerType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionTrigger.SetMasterActionTrigger(t.Id, t.Name, t.ActionTriggerType)
	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindByActionTriggerType", fmt.Sprintf("%d_", actionTriggerType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionTriggerDao) FindOrNilByActionTriggerType(ctx context.Context, actionTriggerType enum.ActionTriggerType) (*masterActionTrigger.MasterActionTrigger, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindOrNilByActionTriggerType", fmt.Sprintf("%d_", actionTriggerType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionTrigger.MasterActionTrigger); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionTrigger()
	res := s.ReadConn.WithContext(ctx).Where("action_trigger_type = ?", actionTriggerType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionTrigger.SetMasterActionTrigger(t.Id, t.Name, t.ActionTriggerType)
	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindOrNilByActionTriggerType", fmt.Sprintf("%d_", actionTriggerType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionTriggerDao) FindList(ctx context.Context) (masterActionTrigger.MasterActionTriggers, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionTrigger.MasterActionTriggers); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionTriggers()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionTrigger.NewMasterActionTriggers()
	for _, t := range ts {
		ms = append(ms, masterActionTrigger.SetMasterActionTrigger(t.Id, t.Name, t.ActionTriggerType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionTriggerDao) FindListByActionTriggerType(ctx context.Context, actionTriggerType enum.ActionTriggerType) (masterActionTrigger.MasterActionTriggers, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindListByActionTriggerType", fmt.Sprintf("%d_", actionTriggerType)))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionTrigger.MasterActionTriggers); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionTriggers()
	res := s.ReadConn.WithContext(ctx).Where("action_trigger_type = ?", actionTriggerType).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionTrigger.NewMasterActionTriggers()
	for _, t := range ts {
		ms = append(ms, masterActionTrigger.SetMasterActionTrigger(t.Id, t.Name, t.ActionTriggerType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindListByActionTriggerType", fmt.Sprintf("%d_", actionTriggerType)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionTriggerDao) Create(ctx context.Context, tx *gorm.DB, m *masterActionTrigger.MasterActionTrigger) (*masterActionTrigger.MasterActionTrigger, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterActionTrigger{
		Id:                m.Id,
		Name:              m.Name,
		ActionTriggerType: m.ActionTriggerType,
	}
	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionTrigger.SetMasterActionTrigger(t.Id, t.Name, t.ActionTriggerType), nil
}

func (s *masterActionTriggerDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterActionTrigger.MasterActionTriggers) (masterActionTrigger.MasterActionTriggers, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterActionTriggers()
	for _, m := range ms {
		t := &MasterActionTrigger{
			Id:                m.Id,
			Name:              m.Name,
			ActionTriggerType: m.ActionTriggerType,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionTriggerDao) Update(ctx context.Context, tx *gorm.DB, m *masterActionTrigger.MasterActionTrigger) (*masterActionTrigger.MasterActionTrigger, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterActionTrigger{
		Id:                m.Id,
		Name:              m.Name,
		ActionTriggerType: m.ActionTriggerType,
	}
	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionTrigger.SetMasterActionTrigger(t.Id, t.Name, t.ActionTriggerType), nil
}

func (s *masterActionTriggerDao) Delete(ctx context.Context, tx *gorm.DB, m *masterActionTrigger.MasterActionTrigger) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterActionTrigger())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
