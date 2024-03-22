// Package masterIdleBonusEvent 放置ボーナスイベント
package masterIdleBonusEvent

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusEvent"
)

type masterIdleBonusEventDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterIdleBonusEventDao(conn *database.MysqlHandler) masterIdleBonusEvent.MasterIdleBonusEventRepository {
	return &masterIdleBonusEventDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterIdleBonusEventDao) Find(ctx context.Context, id int64) (*masterIdleBonusEvent.MasterIdleBonusEvent, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_event", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusEvent.MasterIdleBonusEvent); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusEvent()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonusEvent.SetMasterIdleBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_event", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusEventDao) FindOrNil(ctx context.Context, id int64) (*masterIdleBonusEvent.MasterIdleBonusEvent, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_event", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusEvent.MasterIdleBonusEvent); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusEvent()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonusEvent.SetMasterIdleBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_event", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusEventDao) FindList(ctx context.Context) (masterIdleBonusEvent.MasterIdleBonusEvents, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_event", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonusEvent.MasterIdleBonusEvents); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonusEvents()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonusEvent.NewMasterIdleBonusEvents()
	for _, t := range ts {
		ms = append(ms, masterIdleBonusEvent.SetMasterIdleBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_event", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusEventDao) Create(ctx context.Context, tx *gorm.DB, m *masterIdleBonusEvent.MasterIdleBonusEvent) (*masterIdleBonusEvent.MasterIdleBonusEvent, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterIdleBonusEvent{
		Id:            m.Id,
		Name:          m.Name,
		ResetHour:     m.ResetHour,
		IntervalHour:  m.IntervalHour,
		RepeatSetting: m.RepeatSetting,
		StartAt:       m.StartAt,
		EndAt:         m.EndAt,
	}
	res := conn.Model(NewMasterIdleBonusEvent()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterIdleBonusEvent.SetMasterIdleBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt), nil
}

func (s *masterIdleBonusEventDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterIdleBonusEvent.MasterIdleBonusEvents) (masterIdleBonusEvent.MasterIdleBonusEvents, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterIdleBonusEvents()
	for _, m := range ms {
		t := &MasterIdleBonusEvent{
			Id:            m.Id,
			Name:          m.Name,
			ResetHour:     m.ResetHour,
			IntervalHour:  m.IntervalHour,
			RepeatSetting: m.RepeatSetting,
			StartAt:       m.StartAt,
			EndAt:         m.EndAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterIdleBonusEvent()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterIdleBonusEventDao) Update(ctx context.Context, tx *gorm.DB, m *masterIdleBonusEvent.MasterIdleBonusEvent) (*masterIdleBonusEvent.MasterIdleBonusEvent, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterIdleBonusEvent{
		Id:            m.Id,
		Name:          m.Name,
		ResetHour:     m.ResetHour,
		IntervalHour:  m.IntervalHour,
		RepeatSetting: m.RepeatSetting,
		StartAt:       m.StartAt,
		EndAt:         m.EndAt,
	}
	res := conn.Model(NewMasterIdleBonusEvent()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterIdleBonusEvent.SetMasterIdleBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt), nil
}

func (s *masterIdleBonusEventDao) Delete(ctx context.Context, tx *gorm.DB, m *masterIdleBonusEvent.MasterIdleBonusEvent) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterIdleBonusEvent()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterIdleBonusEvent())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
