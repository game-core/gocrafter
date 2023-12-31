package event

import (
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	dbChashe "github.com/game-core/gocrafter/config/cashe"
	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/master/event"
	eventRepository "github.com/game-core/gocrafter/domain/repository/master/event"
)

type eventDao struct {
	Read  *gorm.DB
	Write *gorm.DB
	Cache *cache.Cache
}

func NewEventDao(conn *database.SqlHandler) eventRepository.EventRepository {
	return &eventDao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
		Cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (d *eventDao) Create(entity *event.Event, tx *gorm.DB) (*event.Event, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&event.Event{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *eventDao) Delete(entity *event.Event, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&event.Event{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *eventDao) FindByID(ID int64) (*event.Event, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("event", "FindByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*event.Event); ok {
			return cachedEntity, nil
		}
	}

	entity := &event.Event{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("event", "FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *eventDao) FindByName(Name string) (*event.Event, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("event", "FindByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*event.Event); ok {
			return cachedEntity, nil
		}
	}

	entity := &event.Event{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("event", "FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *eventDao) FindOrNilByID(ID int64) (*event.Event, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("event", "FindOrNilByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*event.Event); ok {
			return cachedEntity, nil
		}
	}

	entity := &event.Event{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("event", "FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *eventDao) FindOrNilByName(Name string) (*event.Event, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("event", "FindOrNilByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*event.Event); ok {
			return cachedEntity, nil
		}
	}

	entity := &event.Event{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("event", "FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *eventDao) List(limit int) (*event.Events, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("event", "List", ""))
	if found {
		if cachedEntity, ok := cachedResult.(*event.Events); ok {
			return cachedEntity, nil
		}
	}

	entity := &event.Events{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("event", "List", ""), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *eventDao) ListByName(Name string) (*event.Events, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("event", "ListByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*event.Events); ok {
			return cachedEntity, nil
		}
	}

	entity := &event.Events{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("event", "ListByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *eventDao) Save(entity *event.Event, tx *gorm.DB) (*event.Event, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&event.Event{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
