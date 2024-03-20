// Package masterLoginBonus ログインボーナス
package masterLoginBonus

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonus"
)

type masterLoginBonusDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterLoginBonusDao(conn *database.SqlHandler) masterLoginBonus.MasterLoginBonusRepository {
	return &masterLoginBonusDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterLoginBonusDao) Find(ctx context.Context, id int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusDao) FindOrNil(ctx context.Context, id int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusDao) FindByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadConn.WithContext(ctx).Where("master_login_bonus_event_id = ?", masterLoginBonusEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusDao) FindOrNilByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindOrNilByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadConn.WithContext(ctx).Where("master_login_bonus_event_id = ?", masterLoginBonusEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindOrNilByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusDao) FindList(ctx context.Context) (masterLoginBonus.MasterLoginBonuses, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonus.MasterLoginBonuses); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonuses()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonus.NewMasterLoginBonuses()
	for _, t := range ts {
		ms = append(ms, masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusDao) FindListByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (masterLoginBonus.MasterLoginBonuses, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindListByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonus.MasterLoginBonuses); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonuses()
	res := s.ReadConn.WithContext(ctx).Where("master_login_bonus_event_id = ?", masterLoginBonusEventId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonus.NewMasterLoginBonuses()
	for _, t := range ts {
		ms = append(ms, masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindListByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusDao) Create(ctx context.Context, tx *gorm.DB, m *masterLoginBonus.MasterLoginBonus) (*masterLoginBonus.MasterLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterLoginBonus{
		Id:                      m.Id,
		MasterLoginBonusEventId: m.MasterLoginBonusEventId,
		Name:                    m.Name,
	}
	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name), nil
}

func (s *masterLoginBonusDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterLoginBonus.MasterLoginBonuses) (masterLoginBonus.MasterLoginBonuses, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterLoginBonuses()
	for _, m := range ms {
		t := &MasterLoginBonus{
			Id:                      m.Id,
			MasterLoginBonusEventId: m.MasterLoginBonusEventId,
			Name:                    m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterLoginBonusDao) Update(ctx context.Context, tx *gorm.DB, m *masterLoginBonus.MasterLoginBonus) (*masterLoginBonus.MasterLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterLoginBonus{
		Id:                      m.Id,
		MasterLoginBonusEventId: m.MasterLoginBonusEventId,
		Name:                    m.Name,
	}
	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name), nil
}

func (s *masterLoginBonusDao) Delete(ctx context.Context, tx *gorm.DB, m *masterLoginBonus.MasterLoginBonus) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterLoginBonus())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
