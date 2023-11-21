package example

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/master/example"
	exampleRepository "github.com/game-core/gocrafter/domain/repository/master/example"
)

type exampleDao struct {
	Read  *gorm.DB
	Write *gorm.DB
	Cache *cache.Cache
}

func NewExampleDao(conn *database.SqlHandler) exampleRepository.ExampleRepository {
	return &exampleDao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
		Cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (d *exampleDao) Create(entity *example.Example, tx *gorm.DB) (*example.Example, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&example.Example{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *exampleDao) Delete(entity *example.Example, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&example.Example{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *exampleDao) FindByID(ID int64) (*example.Example, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*example.Example); ok {
			return cachedEntity, nil
		}
	}

	entity := &example.Example{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *exampleDao) FindByIDAndName(ID int64, Name string) (*example.Example, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByIDAndName", fmt.Sprintf("%d_%s_", ID, Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*example.Example); ok {
			return cachedEntity, nil
		}
	}

	entity := &example.Example{}
	res := d.Read.Where("id = ?", ID).Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByIDAndName", fmt.Sprintf("%d_%s_", ID, Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *exampleDao) FindByName(Name string) (*example.Example, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*example.Example); ok {
			return cachedEntity, nil
		}
	}

	entity := &example.Example{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *exampleDao) List(limit int64) (*example.Examples, error) {
	cachedResult, found := d.Cache.Get(cacheKey("List", ""))
	if found {
		if cachedEntity, ok := cachedResult.(*example.Examples); ok {
			return cachedEntity, nil
		}
	}

	entity := &example.Examples{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("List", ""), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *exampleDao) ListByIDAndName(ID int64, Name string) (*example.Examples, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByIDAndName", fmt.Sprintf("%d_%s_", ID, Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*example.Examples); ok {
			return cachedEntity, nil
		}
	}

	entity := &example.Examples{}
	res := d.Read.Where("id = ?", ID).Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("ListByIDAndName", fmt.Sprintf("%d_%s_", ID, Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *exampleDao) ListByName(Name string) (*example.Examples, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*example.Examples); ok {
			return cachedEntity, nil
		}
	}

	entity := &example.Examples{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("ListByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *exampleDao) Update(entity *example.Example, tx *gorm.DB) (*example.Example, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&example.Example{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func cacheKey(method string, key string) string {
	return fmt.Sprintf("example:%s:%s", method, key)
}
