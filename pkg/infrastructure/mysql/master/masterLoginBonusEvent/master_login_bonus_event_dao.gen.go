// Package masterLoginBonusEvent ログインボーナスイベント
package masterLoginBonusEvent

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusEvent"
)

type masterLoginBonusEventDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterLoginBonusEventDao(conn *database.SqlHandler) masterLoginBonusEvent.MasterLoginBonusEventRepository {
	return &masterLoginBonusEventDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterLoginBonusEventDao) Find(ctx context.Context, id int64) (*masterLoginBonusEvent.MasterLoginBonusEvent, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_event", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusEvent.MasterLoginBonusEvent); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusEvent()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonusEvent.SetMasterLoginBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_event", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusEventDao) FindOrNil(ctx context.Context, id int64) (*masterLoginBonusEvent.MasterLoginBonusEvent, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_event", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusEvent.MasterLoginBonusEvent); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusEvent()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonusEvent.SetMasterLoginBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_event", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusEventDao) FindList(ctx context.Context) (masterLoginBonusEvent.MasterLoginBonusEvents, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_event", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonusEvent.MasterLoginBonusEvents); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonusEvents()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonusEvent.NewMasterLoginBonusEvents()
	for _, t := range ts {
		ms = append(ms, masterLoginBonusEvent.SetMasterLoginBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_event", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusEventDao) Create(ctx context.Context, tx *gorm.DB, m *masterLoginBonusEvent.MasterLoginBonusEvent) (*masterLoginBonusEvent.MasterLoginBonusEvent, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterLoginBonusEvent{
		Id:            m.Id,
		Name:          m.Name,
		ResetHour:     m.ResetHour,
		IntervalHour:  m.IntervalHour,
		RepeatSetting: m.RepeatSetting,
		StartAt:       m.StartAt,
		EndAt:         m.EndAt,
	}
	res := conn.Model(NewMasterLoginBonusEvent()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonusEvent.SetMasterLoginBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt), nil
}

func (s *masterLoginBonusEventDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterLoginBonusEvent.MasterLoginBonusEvents) (masterLoginBonusEvent.MasterLoginBonusEvents, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterLoginBonusEvents()
	for _, m := range ms {
		t := &MasterLoginBonusEvent{
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

	res := conn.Model(NewMasterLoginBonusEvent()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterLoginBonusEventDao) Update(ctx context.Context, tx *gorm.DB, m *masterLoginBonusEvent.MasterLoginBonusEvent) (*masterLoginBonusEvent.MasterLoginBonusEvent, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterLoginBonusEvent{
		Id:            m.Id,
		Name:          m.Name,
		ResetHour:     m.ResetHour,
		IntervalHour:  m.IntervalHour,
		RepeatSetting: m.RepeatSetting,
		StartAt:       m.StartAt,
		EndAt:         m.EndAt,
	}
	res := conn.Model(NewMasterLoginBonusEvent()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonusEvent.SetMasterLoginBonusEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt), nil
}

func (s *masterLoginBonusEventDao) Delete(ctx context.Context, tx *gorm.DB, m *masterLoginBonusEvent.MasterLoginBonusEvent) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterLoginBonusEvent()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterLoginBonusEvent())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
