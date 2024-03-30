// Package masterRankingScope ランキング範囲
package masterRankingScope

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingScope"
)

type masterRankingScopeDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterRankingScopeDao(conn *database.MysqlHandler) masterRankingScope.MasterRankingScopeMysqlRepository {
	return &masterRankingScopeDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterRankingScopeDao) Find(ctx context.Context, id int64) (*masterRankingScope.MasterRankingScope, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking_scope", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRankingScope.MasterRankingScope); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRankingScope()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterRankingScope.SetMasterRankingScope(t.Id, t.Name, t.RankingType)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking_scope", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingScopeDao) FindOrNil(ctx context.Context, id int64) (*masterRankingScope.MasterRankingScope, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking_scope", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRankingScope.MasterRankingScope); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRankingScope()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterRankingScope.SetMasterRankingScope(t.Id, t.Name, t.RankingType)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking_scope", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingScopeDao) FindByRankingType(ctx context.Context, rankingType enum.RankingType) (*masterRankingScope.MasterRankingScope, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking_scope", "FindByRankingType", fmt.Sprintf("%d_", rankingType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRankingScope.MasterRankingScope); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRankingScope()
	res := s.ReadMysqlConn.WithContext(ctx).Where("ranking_type = ?", rankingType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterRankingScope.SetMasterRankingScope(t.Id, t.Name, t.RankingType)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking_scope", "FindByRankingType", fmt.Sprintf("%d_", rankingType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingScopeDao) FindOrNilByRankingType(ctx context.Context, rankingType enum.RankingType) (*masterRankingScope.MasterRankingScope, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking_scope", "FindOrNilByRankingType", fmt.Sprintf("%d_", rankingType)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRankingScope.MasterRankingScope); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRankingScope()
	res := s.ReadMysqlConn.WithContext(ctx).Where("ranking_type = ?", rankingType).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterRankingScope.SetMasterRankingScope(t.Id, t.Name, t.RankingType)
	s.Cache.Set(cashes.CreateCacheKey("master_ranking_scope", "FindOrNilByRankingType", fmt.Sprintf("%d_", rankingType)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRankingScopeDao) FindList(ctx context.Context) (masterRankingScope.MasterRankingScopes, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking_scope", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterRankingScope.MasterRankingScopes); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterRankingScopes()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterRankingScope.NewMasterRankingScopes()
	for _, t := range ts {
		ms = append(ms, masterRankingScope.SetMasterRankingScope(t.Id, t.Name, t.RankingType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_ranking_scope", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterRankingScopeDao) FindListByRankingType(ctx context.Context, rankingType enum.RankingType) (masterRankingScope.MasterRankingScopes, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_ranking_scope", "FindListByRankingType", fmt.Sprintf("%d_", rankingType)))
	if found {
		if cachedEntity, ok := cachedResult.(masterRankingScope.MasterRankingScopes); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterRankingScopes()
	res := s.ReadMysqlConn.WithContext(ctx).Where("ranking_type = ?", rankingType).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterRankingScope.NewMasterRankingScopes()
	for _, t := range ts {
		ms = append(ms, masterRankingScope.SetMasterRankingScope(t.Id, t.Name, t.RankingType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_ranking_scope", "FindListByRankingType", fmt.Sprintf("%d_", rankingType)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterRankingScopeDao) Create(ctx context.Context, tx *gorm.DB, m *masterRankingScope.MasterRankingScope) (*masterRankingScope.MasterRankingScope, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRankingScope{
		Id:          m.Id,
		Name:        m.Name,
		RankingType: m.RankingType,
	}
	res := conn.Model(NewMasterRankingScope()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRankingScope.SetMasterRankingScope(t.Id, t.Name, t.RankingType), nil
}

func (s *masterRankingScopeDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterRankingScope.MasterRankingScopes) (masterRankingScope.MasterRankingScopes, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterRankingScopes()
	for _, m := range ms {
		t := &MasterRankingScope{
			Id:          m.Id,
			Name:        m.Name,
			RankingType: m.RankingType,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterRankingScope()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterRankingScopeDao) Update(ctx context.Context, tx *gorm.DB, m *masterRankingScope.MasterRankingScope) (*masterRankingScope.MasterRankingScope, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRankingScope{
		Id:          m.Id,
		Name:        m.Name,
		RankingType: m.RankingType,
	}
	res := conn.Model(NewMasterRankingScope()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRankingScope.SetMasterRankingScope(t.Id, t.Name, t.RankingType), nil
}

func (s *masterRankingScopeDao) Delete(ctx context.Context, tx *gorm.DB, m *masterRankingScope.MasterRankingScope) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterRankingScope()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterRankingScope())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
