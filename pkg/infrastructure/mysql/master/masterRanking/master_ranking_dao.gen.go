// Package masterRanking ランキング
package masterRanking

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRanking"
)

type masterRankingDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterRankingDao(conn *database.MysqlHandler) masterRanking.MasterRankingMysqlRepository {
	return &masterRankingDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterRankingDao) Find(ctx context.Context, id int64) (*masterRanking.MasterRanking, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRanking.MasterRanking); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRanking()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterRanking.SetMasterRanking(t.Id, t.MasterRankingEventId, t.Name, t.RankingScopeType, t.RankingLimit)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingDao) FindOrNil(ctx context.Context, id int64) (*masterRanking.MasterRanking, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRanking.MasterRanking); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRanking()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterRanking.SetMasterRanking(t.Id, t.MasterRankingEventId, t.Name, t.RankingScopeType, t.RankingLimit)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingDao) FindByMasterRankingEventId(ctx context.Context, masterRankingEventId int64) (*masterRanking.MasterRanking, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking", "FindByMasterRankingEventId", fmt.Sprintf("%d_", masterRankingEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRanking.MasterRanking); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRanking()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_event_id = ?", masterRankingEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterRanking.SetMasterRanking(t.Id, t.MasterRankingEventId, t.Name, t.RankingScopeType, t.RankingLimit)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking", "FindByMasterRankingEventId", fmt.Sprintf("%d_", masterRankingEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingDao) FindOrNilByMasterRankingEventId(ctx context.Context, masterRankingEventId int64) (*masterRanking.MasterRanking, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking", "FindOrNilByMasterRankingEventId", fmt.Sprintf("%d_", masterRankingEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRanking.MasterRanking); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRanking()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_event_id = ?", masterRankingEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterRanking.SetMasterRanking(t.Id, t.MasterRankingEventId, t.Name, t.RankingScopeType, t.RankingLimit)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking", "FindOrNilByMasterRankingEventId", fmt.Sprintf("%d_", masterRankingEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingDao) FindList(ctx context.Context) (masterRanking.MasterRankings, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterRanking.MasterRankings); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterRankings()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterRanking.NewMasterRankings()
	for _, t := range ts {
		ms = append(ms, masterRanking.SetMasterRanking(t.Id, t.MasterRankingEventId, t.Name, t.RankingScopeType, t.RankingLimit))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_ranking", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterRankingDao) FindListByMasterRankingEventId(ctx context.Context, masterRankingEventId int64) (masterRanking.MasterRankings, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking", "FindListByMasterRankingEventId", fmt.Sprintf("%d_", masterRankingEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterRanking.MasterRankings); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterRankings()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_ranking_event_id = ?", masterRankingEventId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterRanking.NewMasterRankings()
	for _, t := range ts {
		ms = append(ms, masterRanking.SetMasterRanking(t.Id, t.MasterRankingEventId, t.Name, t.RankingScopeType, t.RankingLimit))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_ranking", "FindListByMasterRankingEventId", fmt.Sprintf("%d_", masterRankingEventId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterRankingDao) Create(ctx context.Context, tx *gorm.DB, m *masterRanking.MasterRanking) (*masterRanking.MasterRanking, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRanking{
		Id:                   m.Id,
		MasterRankingEventId: m.MasterRankingEventId,
		Name:                 m.Name,
		RankingScopeType:     m.RankingScopeType,
		RankingLimit:         m.RankingLimit,
	}
	res := conn.Model(NewMasterRanking()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRanking.SetMasterRanking(t.Id, t.MasterRankingEventId, t.Name, t.RankingScopeType, t.RankingLimit), nil
}

func (s *masterRankingDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterRanking.MasterRankings) (masterRanking.MasterRankings, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterRankings()
	for _, m := range ms {
		t := &MasterRanking{
			Id:                   m.Id,
			MasterRankingEventId: m.MasterRankingEventId,
			Name:                 m.Name,
			RankingScopeType:     m.RankingScopeType,
			RankingLimit:         m.RankingLimit,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterRanking()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterRankingDao) Update(ctx context.Context, tx *gorm.DB, m *masterRanking.MasterRanking) (*masterRanking.MasterRanking, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRanking{
		Id:                   m.Id,
		MasterRankingEventId: m.MasterRankingEventId,
		Name:                 m.Name,
		RankingScopeType:     m.RankingScopeType,
		RankingLimit:         m.RankingLimit,
	}
	res := conn.Model(NewMasterRanking()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRanking.SetMasterRanking(t.Id, t.MasterRankingEventId, t.Name, t.RankingScopeType, t.RankingLimit), nil
}

func (s *masterRankingDao) Delete(ctx context.Context, tx *gorm.DB, m *masterRanking.MasterRanking) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterRanking()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterRanking())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
