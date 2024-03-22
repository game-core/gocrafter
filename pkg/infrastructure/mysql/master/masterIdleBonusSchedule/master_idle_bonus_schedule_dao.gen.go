// Package masterIdleBonusSchedule 放置ボーナススケジュール
package masterIdleBonusSchedule

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusSchedule"
)

type masterIdleBonusScheduleDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterIdleBonusScheduleDao(conn *database.MysqlHandler) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
	return &masterIdleBonusScheduleDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterIdleBonusScheduleDao) Find(ctx context.Context, id int64) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusSchedule.MasterIdleBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusSchedule()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusScheduleDao) FindOrNil(ctx context.Context, id int64) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusSchedule.MasterIdleBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusSchedule()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusScheduleDao) FindByMasterIdleBonusId(ctx context.Context, masterIdleBonusId int64) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindByMasterIdleBonusId", fmt.Sprintf("%d_", masterIdleBonusId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusSchedule.MasterIdleBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusSchedule()
	res := s.ReadConn.WithContext(ctx).Where("master_idle_bonus_id = ?", masterIdleBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindByMasterIdleBonusId", fmt.Sprintf("%d_", masterIdleBonusId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusScheduleDao) FindByStep(ctx context.Context, step int32) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindByStep", fmt.Sprintf("%d_", step)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusSchedule.MasterIdleBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusSchedule()
	res := s.ReadConn.WithContext(ctx).Where("step = ?", step).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindByStep", fmt.Sprintf("%d_", step)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusScheduleDao) FindByMasterIdleBonusIdAndStep(ctx context.Context, masterIdleBonusId int64, step int32) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindByMasterIdleBonusIdAndStep", fmt.Sprintf("%d_%d_", masterIdleBonusId, step)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusSchedule.MasterIdleBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusSchedule()
	res := s.ReadConn.WithContext(ctx).Where("master_idle_bonus_id = ?", masterIdleBonusId).Where("step = ?", step).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindByMasterIdleBonusIdAndStep", fmt.Sprintf("%d_%d_", masterIdleBonusId, step)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusScheduleDao) FindOrNilByMasterIdleBonusId(ctx context.Context, masterIdleBonusId int64) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindOrNilByMasterIdleBonusId", fmt.Sprintf("%d_", masterIdleBonusId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusSchedule.MasterIdleBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusSchedule()
	res := s.ReadConn.WithContext(ctx).Where("master_idle_bonus_id = ?", masterIdleBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindOrNilByMasterIdleBonusId", fmt.Sprintf("%d_", masterIdleBonusId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusScheduleDao) FindOrNilByStep(ctx context.Context, step int32) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindOrNilByStep", fmt.Sprintf("%d_", step)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusSchedule.MasterIdleBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusSchedule()
	res := s.ReadConn.WithContext(ctx).Where("step = ?", step).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindOrNilByStep", fmt.Sprintf("%d_", step)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusScheduleDao) FindOrNilByMasterIdleBonusIdAndStep(ctx context.Context, masterIdleBonusId int64, step int32) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindOrNilByMasterIdleBonusIdAndStep", fmt.Sprintf("%d_%d_", masterIdleBonusId, step)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusSchedule.MasterIdleBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusSchedule()
	res := s.ReadConn.WithContext(ctx).Where("master_idle_bonus_id = ?", masterIdleBonusId).Where("step = ?", step).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindOrNilByMasterIdleBonusIdAndStep", fmt.Sprintf("%d_%d_", masterIdleBonusId, step)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusScheduleDao) FindList(ctx context.Context) (masterIdleBonusSchedule.MasterIdleBonusSchedules, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonusSchedule.MasterIdleBonusSchedules); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonusSchedules()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonusSchedule.NewMasterIdleBonusSchedules()
	for _, t := range ts {
		ms = append(ms, masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusScheduleDao) FindListByMasterIdleBonusId(ctx context.Context, masterIdleBonusId int64) (masterIdleBonusSchedule.MasterIdleBonusSchedules, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindListByMasterIdleBonusId", fmt.Sprintf("%d_", masterIdleBonusId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonusSchedule.MasterIdleBonusSchedules); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonusSchedules()
	res := s.ReadConn.WithContext(ctx).Where("master_idle_bonus_id = ?", masterIdleBonusId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonusSchedule.NewMasterIdleBonusSchedules()
	for _, t := range ts {
		ms = append(ms, masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindListByMasterIdleBonusId", fmt.Sprintf("%d_", masterIdleBonusId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusScheduleDao) FindListByStep(ctx context.Context, step int32) (masterIdleBonusSchedule.MasterIdleBonusSchedules, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindListByStep", fmt.Sprintf("%d_", step)))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonusSchedule.MasterIdleBonusSchedules); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonusSchedules()
	res := s.ReadConn.WithContext(ctx).Where("step = ?", step).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonusSchedule.NewMasterIdleBonusSchedules()
	for _, t := range ts {
		ms = append(ms, masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindListByStep", fmt.Sprintf("%d_", step)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusScheduleDao) FindListByMasterIdleBonusIdAndStep(ctx context.Context, masterIdleBonusId int64, step int32) (masterIdleBonusSchedule.MasterIdleBonusSchedules, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindListByMasterIdleBonusIdAndStep", fmt.Sprintf("%d_%d_", masterIdleBonusId, step)))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonusSchedule.MasterIdleBonusSchedules); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonusSchedules()
	res := s.ReadConn.WithContext(ctx).Where("master_idle_bonus_id = ?", masterIdleBonusId).Where("step = ?", step).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonusSchedule.NewMasterIdleBonusSchedules()
	for _, t := range ts {
		ms = append(ms, masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_schedule", "FindListByMasterIdleBonusIdAndStep", fmt.Sprintf("%d_%d_", masterIdleBonusId, step)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusScheduleDao) Create(ctx context.Context, tx *gorm.DB, m *masterIdleBonusSchedule.MasterIdleBonusSchedule) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterIdleBonusSchedule{
		Id:                m.Id,
		MasterIdleBonusId: m.MasterIdleBonusId,
		Step:              m.Step,
		Name:              m.Name,
	}
	res := conn.Model(NewMasterIdleBonusSchedule()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name), nil
}

func (s *masterIdleBonusScheduleDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterIdleBonusSchedule.MasterIdleBonusSchedules) (masterIdleBonusSchedule.MasterIdleBonusSchedules, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterIdleBonusSchedules()
	for _, m := range ms {
		t := &MasterIdleBonusSchedule{
			Id:                m.Id,
			MasterIdleBonusId: m.MasterIdleBonusId,
			Step:              m.Step,
			Name:              m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterIdleBonusSchedule()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterIdleBonusScheduleDao) Update(ctx context.Context, tx *gorm.DB, m *masterIdleBonusSchedule.MasterIdleBonusSchedule) (*masterIdleBonusSchedule.MasterIdleBonusSchedule, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterIdleBonusSchedule{
		Id:                m.Id,
		MasterIdleBonusId: m.MasterIdleBonusId,
		Step:              m.Step,
		Name:              m.Name,
	}
	res := conn.Model(NewMasterIdleBonusSchedule()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterIdleBonusSchedule.SetMasterIdleBonusSchedule(t.Id, t.MasterIdleBonusId, t.Step, t.Name), nil
}

func (s *masterIdleBonusScheduleDao) Delete(ctx context.Context, tx *gorm.DB, m *masterIdleBonusSchedule.MasterIdleBonusSchedule) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterIdleBonusSchedule()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterIdleBonusSchedule())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
