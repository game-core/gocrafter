// Package masterIdleBonus 放置ボーナス
package masterIdleBonus

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonus"
)

type masterIdleBonusDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterIdleBonusDao(conn *database.MysqlHandler) masterIdleBonus.MasterIdleBonusRepository {
	return &masterIdleBonusDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterIdleBonusDao) Find(ctx context.Context, id int64) (*masterIdleBonus.MasterIdleBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonus.MasterIdleBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonus()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonus.SetMasterIdleBonus(t.Id, t.MasterIdleBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusDao) FindOrNil(ctx context.Context, id int64) (*masterIdleBonus.MasterIdleBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonus.MasterIdleBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonus()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonus.SetMasterIdleBonus(t.Id, t.MasterIdleBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusDao) FindByMasterIdleBonusEventId(ctx context.Context, masterIdleBonusEventId int64) (*masterIdleBonus.MasterIdleBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus", "FindByMasterIdleBonusEventId", fmt.Sprintf("%d_", masterIdleBonusEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonus.MasterIdleBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonus()
	res := s.ReadConn.WithContext(ctx).Where("master_idle_bonus_event_id = ?", masterIdleBonusEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonus.SetMasterIdleBonus(t.Id, t.MasterIdleBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus", "FindByMasterIdleBonusEventId", fmt.Sprintf("%d_", masterIdleBonusEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusDao) FindOrNilByMasterIdleBonusEventId(ctx context.Context, masterIdleBonusEventId int64) (*masterIdleBonus.MasterIdleBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus", "FindOrNilByMasterIdleBonusEventId", fmt.Sprintf("%d_", masterIdleBonusEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonus.MasterIdleBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonus()
	res := s.ReadConn.WithContext(ctx).Where("master_idle_bonus_event_id = ?", masterIdleBonusEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonus.SetMasterIdleBonus(t.Id, t.MasterIdleBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus", "FindOrNilByMasterIdleBonusEventId", fmt.Sprintf("%d_", masterIdleBonusEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusDao) FindList(ctx context.Context) (masterIdleBonus.MasterIdleBonuses, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonus.MasterIdleBonuses); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonuses()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonus.NewMasterIdleBonuses()
	for _, t := range ts {
		ms = append(ms, masterIdleBonus.SetMasterIdleBonus(t.Id, t.MasterIdleBonusEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusDao) FindListByMasterIdleBonusEventId(ctx context.Context, masterIdleBonusEventId int64) (masterIdleBonus.MasterIdleBonuses, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus", "FindListByMasterIdleBonusEventId", fmt.Sprintf("%d_", masterIdleBonusEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonus.MasterIdleBonuses); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonuses()
	res := s.ReadConn.WithContext(ctx).Where("master_idle_bonus_event_id = ?", masterIdleBonusEventId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonus.NewMasterIdleBonuses()
	for _, t := range ts {
		ms = append(ms, masterIdleBonus.SetMasterIdleBonus(t.Id, t.MasterIdleBonusEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus", "FindListByMasterIdleBonusEventId", fmt.Sprintf("%d_", masterIdleBonusEventId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusDao) Create(ctx context.Context, tx *gorm.DB, m *masterIdleBonus.MasterIdleBonus) (*masterIdleBonus.MasterIdleBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterIdleBonus{
		Id:                     m.Id,
		MasterIdleBonusEventId: m.MasterIdleBonusEventId,
		Name:                   m.Name,
	}
	res := conn.Model(NewMasterIdleBonus()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterIdleBonus.SetMasterIdleBonus(t.Id, t.MasterIdleBonusEventId, t.Name), nil
}

func (s *masterIdleBonusDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterIdleBonus.MasterIdleBonuses) (masterIdleBonus.MasterIdleBonuses, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterIdleBonuses()
	for _, m := range ms {
		t := &MasterIdleBonus{
			Id:                     m.Id,
			MasterIdleBonusEventId: m.MasterIdleBonusEventId,
			Name:                   m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterIdleBonus()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterIdleBonusDao) Update(ctx context.Context, tx *gorm.DB, m *masterIdleBonus.MasterIdleBonus) (*masterIdleBonus.MasterIdleBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterIdleBonus{
		Id:                     m.Id,
		MasterIdleBonusEventId: m.MasterIdleBonusEventId,
		Name:                   m.Name,
	}
	res := conn.Model(NewMasterIdleBonus()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterIdleBonus.SetMasterIdleBonus(t.Id, t.MasterIdleBonusEventId, t.Name), nil
}

func (s *masterIdleBonusDao) Delete(ctx context.Context, tx *gorm.DB, m *masterIdleBonus.MasterIdleBonus) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterIdleBonus()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterIdleBonus())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
