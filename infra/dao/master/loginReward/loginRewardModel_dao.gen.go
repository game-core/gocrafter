package loginReward

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/master/loginReward"
	loginRewardRepository "github.com/game-core/gocrafter/domain/repository/master/loginReward"
)

type loginRewardModelDao struct {
	Read  *gorm.DB
	Write *gorm.DB
	Cache *cache.Cache
}

func NewLoginRewardModelDao(conn *database.SqlHandler) loginRewardRepository.LoginRewardModelRepository {
	return &loginRewardModelDao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
		Cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (d *loginRewardModelDao) Create(entity *loginReward.LoginRewardModel, tx *gorm.DB) (*loginReward.LoginRewardModel, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardModel{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardModelDao) Delete(entity *loginReward.LoginRewardModel, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardModel{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *loginRewardModelDao) FindByEventID(EventID int64) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByEventID", fmt.Sprintf("%d_", EventID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("event_id = ?", EventID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByEventID", fmt.Sprintf("%d_", EventID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindByID(ID int64) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindByName(Name string) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindByNameAndEventID(Name string, EventID int64) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByNameAndEventID", fmt.Sprintf("%s_%d_", Name, EventID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("name = ?", Name).Where("event_id = ?", EventID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByNameAndEventID", fmt.Sprintf("%s_%d_", Name, EventID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) List(limit int64) (*loginReward.LoginRewardModels, error) {
	cachedResult, found := d.Cache.Get(cacheKey("List", ""))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModels); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModels{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("List", ""), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) ListByEventID(EventID int64) (*loginReward.LoginRewardModels, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByEventID", fmt.Sprintf("%d_", EventID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModels); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModels{}
	res := d.Read.Where("event_id = ?", EventID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("ListByEventID", fmt.Sprintf("%d_", EventID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) ListByName(Name string) (*loginReward.LoginRewardModels, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModels); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModels{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("ListByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) ListByNameAndEventID(Name string, EventID int64) (*loginReward.LoginRewardModels, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByNameAndEventID", fmt.Sprintf("%s_%d_", Name, EventID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModels); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModels{}
	res := d.Read.Where("name = ?", Name).Where("event_id = ?", EventID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("ListByNameAndEventID", fmt.Sprintf("%s_%d_", Name, EventID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) Update(entity *loginReward.LoginRewardModel, tx *gorm.DB) (*loginReward.LoginRewardModel, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardModel{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func cacheKey(method string, key string) string {
	return fmt.Sprintf("loginreward:%s:%s", method, key)
}
