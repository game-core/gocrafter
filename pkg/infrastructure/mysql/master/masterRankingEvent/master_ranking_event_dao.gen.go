// Package masterRankingEvent ランキングイベント
package masterRankingEvent

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingEvent"
)

type masterRankingEventDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterRankingEventDao(conn *database.MysqlHandler) masterRankingEvent.MasterRankingEventMysqlRepository {
	return &masterRankingEventDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterRankingEventDao) Find(ctx context.Context, id int64) (*masterRankingEvent.MasterRankingEvent, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking_event", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRankingEvent.MasterRankingEvent); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRankingEvent()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterRankingEvent.SetMasterRankingEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking_event", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingEventDao) FindOrNil(ctx context.Context, id int64) (*masterRankingEvent.MasterRankingEvent, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking_event", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRankingEvent.MasterRankingEvent); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRankingEvent()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterRankingEvent.SetMasterRankingEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking_event", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingEventDao) FindList(ctx context.Context) (masterRankingEvent.MasterRankingEvents, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking_event", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterRankingEvent.MasterRankingEvents); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterRankingEvents()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterRankingEvent.NewMasterRankingEvents()
	for _, t := range ts {
		ms = append(ms, masterRankingEvent.SetMasterRankingEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_ranking_event", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterRankingEventDao) Create(ctx context.Context, tx *gorm.DB, m *masterRankingEvent.MasterRankingEvent) (*masterRankingEvent.MasterRankingEvent, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRankingEvent{
		Id:            m.Id,
		Name:          m.Name,
		ResetHour:     m.ResetHour,
		IntervalHour:  m.IntervalHour,
		RepeatSetting: m.RepeatSetting,
		StartAt:       m.StartAt,
		EndAt:         m.EndAt,
	}
	res := conn.Model(NewMasterRankingEvent()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRankingEvent.SetMasterRankingEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt), nil
}

func (s *masterRankingEventDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterRankingEvent.MasterRankingEvents) (masterRankingEvent.MasterRankingEvents, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterRankingEvents()
	for _, m := range ms {
		t := &MasterRankingEvent{
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

	res := conn.Model(NewMasterRankingEvent()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterRankingEventDao) Update(ctx context.Context, tx *gorm.DB, m *masterRankingEvent.MasterRankingEvent) (*masterRankingEvent.MasterRankingEvent, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRankingEvent{
		Id:            m.Id,
		Name:          m.Name,
		ResetHour:     m.ResetHour,
		IntervalHour:  m.IntervalHour,
		RepeatSetting: m.RepeatSetting,
		StartAt:       m.StartAt,
		EndAt:         m.EndAt,
	}
	res := conn.Model(NewMasterRankingEvent()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRankingEvent.SetMasterRankingEvent(t.Id, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt), nil
}

func (s *masterRankingEventDao) Delete(ctx context.Context, tx *gorm.DB, m *masterRankingEvent.MasterRankingEvent) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterRankingEvent()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterRankingEvent())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
