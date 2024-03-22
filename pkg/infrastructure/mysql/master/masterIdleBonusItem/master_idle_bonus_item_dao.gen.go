// Package masterIdleBonusItem 放置ボーナスアイテム
package masterIdleBonusItem

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusItem"
)

type masterIdleBonusItemDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterIdleBonusItemDao(conn *database.MysqlHandler) masterIdleBonusItem.MasterIdleBonusItemRepository {
	return &masterIdleBonusItemDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterIdleBonusItemDao) Find(ctx context.Context, id int64) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusItem.MasterIdleBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusItemDao) FindOrNil(ctx context.Context, id int64) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusItem.MasterIdleBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusItemDao) FindByMasterIdleBonusScheduleId(ctx context.Context, masterIdleBonusScheduleId int64) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindByMasterIdleBonusScheduleId", fmt.Sprintf("%d_", masterIdleBonusScheduleId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusItem.MasterIdleBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_idle_bonus_schedule_id = ?", masterIdleBonusScheduleId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindByMasterIdleBonusScheduleId", fmt.Sprintf("%d_", masterIdleBonusScheduleId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusItemDao) FindByMasterItemId(ctx context.Context, masterItemId int64) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindByMasterItemId", fmt.Sprintf("%d_", masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusItem.MasterIdleBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindByMasterItemId", fmt.Sprintf("%d_", masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusItemDao) FindByMasterIdleBonusScheduleIdAndMasterItemId(ctx context.Context, masterIdleBonusScheduleId int64, masterItemId int64) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindByMasterIdleBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterIdleBonusScheduleId, masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusItem.MasterIdleBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_idle_bonus_schedule_id = ?", masterIdleBonusScheduleId).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindByMasterIdleBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterIdleBonusScheduleId, masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusItemDao) FindOrNilByMasterIdleBonusScheduleId(ctx context.Context, masterIdleBonusScheduleId int64) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindOrNilByMasterIdleBonusScheduleId", fmt.Sprintf("%d_", masterIdleBonusScheduleId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusItem.MasterIdleBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_idle_bonus_schedule_id = ?", masterIdleBonusScheduleId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindOrNilByMasterIdleBonusScheduleId", fmt.Sprintf("%d_", masterIdleBonusScheduleId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusItemDao) FindOrNilByMasterItemId(ctx context.Context, masterItemId int64) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindOrNilByMasterItemId", fmt.Sprintf("%d_", masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusItem.MasterIdleBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindOrNilByMasterItemId", fmt.Sprintf("%d_", masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusItemDao) FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId(ctx context.Context, masterIdleBonusScheduleId int64, masterItemId int64) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterIdleBonusScheduleId, masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterIdleBonusItem.MasterIdleBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterIdleBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_idle_bonus_schedule_id = ?", masterIdleBonusScheduleId).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterIdleBonusScheduleId, masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterIdleBonusItemDao) FindList(ctx context.Context) (masterIdleBonusItem.MasterIdleBonusItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonusItem.MasterIdleBonusItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonusItems()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonusItem.NewMasterIdleBonusItems()
	for _, t := range ts {
		ms = append(ms, masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusItemDao) FindListByMasterIdleBonusScheduleId(ctx context.Context, masterIdleBonusScheduleId int64) (masterIdleBonusItem.MasterIdleBonusItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindListByMasterIdleBonusScheduleId", fmt.Sprintf("%d_", masterIdleBonusScheduleId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonusItem.MasterIdleBonusItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonusItems()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_idle_bonus_schedule_id = ?", masterIdleBonusScheduleId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonusItem.NewMasterIdleBonusItems()
	for _, t := range ts {
		ms = append(ms, masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindListByMasterIdleBonusScheduleId", fmt.Sprintf("%d_", masterIdleBonusScheduleId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusItemDao) FindListByMasterItemId(ctx context.Context, masterItemId int64) (masterIdleBonusItem.MasterIdleBonusItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindListByMasterItemId", fmt.Sprintf("%d_", masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonusItem.MasterIdleBonusItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonusItems()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_item_id = ?", masterItemId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonusItem.NewMasterIdleBonusItems()
	for _, t := range ts {
		ms = append(ms, masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindListByMasterItemId", fmt.Sprintf("%d_", masterItemId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusItemDao) FindListByMasterIdleBonusScheduleIdAndMasterItemId(ctx context.Context, masterIdleBonusScheduleId int64, masterItemId int64) (masterIdleBonusItem.MasterIdleBonusItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_idle_bonus_item", "FindListByMasterIdleBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterIdleBonusScheduleId, masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterIdleBonusItem.MasterIdleBonusItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterIdleBonusItems()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_idle_bonus_schedule_id = ?", masterIdleBonusScheduleId).Where("master_item_id = ?", masterItemId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterIdleBonusItem.NewMasterIdleBonusItems()
	for _, t := range ts {
		ms = append(ms, masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_idle_bonus_item", "FindListByMasterIdleBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterIdleBonusScheduleId, masterItemId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterIdleBonusItemDao) Create(ctx context.Context, tx *gorm.DB, m *masterIdleBonusItem.MasterIdleBonusItem) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterIdleBonusItem{
		Id:                        m.Id,
		MasterIdleBonusScheduleId: m.MasterIdleBonusScheduleId,
		MasterItemId:              m.MasterItemId,
		Name:                      m.Name,
		Count:                     m.Count,
	}
	res := conn.Model(NewMasterIdleBonusItem()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count), nil
}

func (s *masterIdleBonusItemDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterIdleBonusItem.MasterIdleBonusItems) (masterIdleBonusItem.MasterIdleBonusItems, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterIdleBonusItems()
	for _, m := range ms {
		t := &MasterIdleBonusItem{
			Id:                        m.Id,
			MasterIdleBonusScheduleId: m.MasterIdleBonusScheduleId,
			MasterItemId:              m.MasterItemId,
			Name:                      m.Name,
			Count:                     m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterIdleBonusItem()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterIdleBonusItemDao) Update(ctx context.Context, tx *gorm.DB, m *masterIdleBonusItem.MasterIdleBonusItem) (*masterIdleBonusItem.MasterIdleBonusItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterIdleBonusItem{
		Id:                        m.Id,
		MasterIdleBonusScheduleId: m.MasterIdleBonusScheduleId,
		MasterItemId:              m.MasterItemId,
		Name:                      m.Name,
		Count:                     m.Count,
	}
	res := conn.Model(NewMasterIdleBonusItem()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterIdleBonusItem.SetMasterIdleBonusItem(t.Id, t.MasterIdleBonusScheduleId, t.MasterItemId, t.Name, t.Count), nil
}

func (s *masterIdleBonusItemDao) Delete(ctx context.Context, tx *gorm.DB, m *masterIdleBonusItem.MasterIdleBonusItem) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterIdleBonusItem()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterIdleBonusItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
