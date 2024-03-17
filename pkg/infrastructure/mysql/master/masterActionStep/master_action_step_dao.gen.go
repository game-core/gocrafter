// Package masterActionStep アクションステップ
package masterActionStep

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionStep"
)

type masterActionStepDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterActionStepDao(conn *database.SqlHandler) masterActionStep.MasterActionStepRepository {
	return &masterActionStepDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterActionStepDao) Find(ctx context.Context, id int64) (*masterActionStep.MasterActionStep, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionStep.MasterActionStep); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionStep()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionStep.SetMasterActionStep(t.Id, t.Name, t.ActionStepType)
	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionStepDao) FindOrNil(ctx context.Context, id int64) (*masterActionStep.MasterActionStep, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionStep.MasterActionStep); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionStep()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionStep.SetMasterActionStep(t.Id, t.Name, t.ActionStepType)
	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionStepDao) FindByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*masterActionStep.MasterActionStep, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindByActionStepType", fmt.Sprintf("%d_", actionStepType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionStep.MasterActionStep); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionStep()
	res := s.ReadConn.WithContext(ctx).Where("action_step_type = ?", actionStepType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionStep.SetMasterActionStep(t.Id, t.Name, t.ActionStepType)
	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindByActionStepType", fmt.Sprintf("%d_", actionStepType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionStepDao) FinOrNilByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*masterActionStep.MasterActionStep, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindOrNilByActionStepType", fmt.Sprintf("%d_", actionStepType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionStep.MasterActionStep); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionStep()
	res := s.ReadConn.WithContext(ctx).Where("action_step_type = ?", actionStepType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionStep.SetMasterActionStep(t.Id, t.Name, t.ActionStepType)
	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindOrNilByActionStepType", fmt.Sprintf("%d_", actionStepType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionStepDao) FindList(ctx context.Context) (masterActionStep.MasterActionSteps, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionStep.MasterActionSteps); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionSteps()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionStep.NewMasterActionSteps()
	for _, t := range ts {
		ms = append(ms, masterActionStep.SetMasterActionStep(t.Id, t.Name, t.ActionStepType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionStepDao) FindListByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (masterActionStep.MasterActionSteps, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindListByActionStepType", fmt.Sprintf("%d_", actionStepType)))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionStep.MasterActionSteps); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionSteps()
	res := s.ReadConn.WithContext(ctx).Where("action_step_type = ?", actionStepType).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionStep.NewMasterActionSteps()
	for _, t := range ts {
		ms = append(ms, masterActionStep.SetMasterActionStep(t.Id, t.Name, t.ActionStepType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindListByActionStepType", fmt.Sprintf("%d_", actionStepType)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionStepDao) Create(ctx context.Context, tx *gorm.DB, m *masterActionStep.MasterActionStep) (*masterActionStep.MasterActionStep, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterActionStep{
		Id:             m.Id,
		Name:           m.Name,
		ActionStepType: m.ActionStepType,
	}
	res := conn.Model(NewMasterActionStep()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionStep.SetMasterActionStep(t.Id, t.Name, t.ActionStepType), nil
}

func (s *masterActionStepDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterActionStep.MasterActionSteps) (masterActionStep.MasterActionSteps, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterActionSteps()
	for _, m := range ms {
		t := &MasterActionStep{
			Id:             m.Id,
			Name:           m.Name,
			ActionStepType: m.ActionStepType,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterActionStep()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionStepDao) Update(ctx context.Context, tx *gorm.DB, m *masterActionStep.MasterActionStep) (*masterActionStep.MasterActionStep, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterActionStep{
		Id:             m.Id,
		Name:           m.Name,
		ActionStepType: m.ActionStepType,
	}
	res := conn.Model(NewMasterActionStep()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionStep.SetMasterActionStep(t.Id, t.Name, t.ActionStepType), nil
}

func (s *masterActionStepDao) Delete(ctx context.Context, tx *gorm.DB, m *masterActionStep.MasterActionStep) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterActionStep()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterActionStep())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
