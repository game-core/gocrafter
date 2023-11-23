package item

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/master/item"
	itemRepository "github.com/game-core/gocrafter/domain/repository/master/item"
)

type itemDao struct {
	Read  *gorm.DB
	Write *gorm.DB
	Cache *cache.Cache
}

func NewItemDao(conn *database.SqlHandler) itemRepository.ItemRepository {
	return &itemDao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
		Cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (d *itemDao) Create(entity *item.Item, tx *gorm.DB) (*item.Item, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&item.Item{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemDao) Delete(entity *item.Item, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&item.Item{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *itemDao) FindByID(ID int64) (*item.Item, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*item.Item); ok {
			return cachedEntity, nil
		}
	}

	entity := &item.Item{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *itemDao) FindOrNilByName(Name string) (*item.Item, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*item.Item); ok {
			return cachedEntity, nil
		}
	}

	entity := &item.Item{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if res.RecordNotFound() {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *itemDao) FindOrNilByID(ID int64) (*item.Item, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*item.Item); ok {
			return cachedEntity, nil
		}
	}

	entity := &item.Item{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if res.RecordNotFound() {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *itemDao) List(limit int64) (*item.Items, error) {
	cachedResult, found := d.Cache.Get(cacheKey("List", ""))
	if found {
		if cachedEntity, ok := cachedResult.(*item.Items); ok {
			return cachedEntity, nil
		}
	}

	entity := &item.Items{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("List", ""), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *itemDao) ListByName(Name string) (*item.Items, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*item.Items); ok {
			return cachedEntity, nil
		}
	}

	entity := &item.Items{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("ListByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *itemDao) Update(entity *item.Item, tx *gorm.DB) (*item.Item, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&item.Item{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func cacheKey(method string, key string) string {
	return fmt.Sprintf("item:%s:%s", method, key)
}
