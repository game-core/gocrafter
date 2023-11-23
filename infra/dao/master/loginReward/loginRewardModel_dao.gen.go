package loginReward

import (
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	dbChashe "github.com/game-core/gocrafter/config/cashe"
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

func (d *loginRewardModelDao) FindByEventName(EventName string) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "FindByEventName", fmt.Sprintf("%s_", EventName)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("event_name = ?", EventName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "FindByEventName", fmt.Sprintf("%s_", EventName)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindByID(ID int64) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "FindByID", fmt.Sprintf("%d_", ID)))
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

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindByName(Name string) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "FindByName", fmt.Sprintf("%s_", Name)))
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

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindByNameAndEventName(Name string, EventName string) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "FindByNameAndEventName", fmt.Sprintf("%s_%s_", Name, EventName)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("name = ?", Name).Where("event_name = ?", EventName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "FindByNameAndEventName", fmt.Sprintf("%s_%s_", Name, EventName)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindOrNilByEventName(EventName string) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "FindOrNilByEventName", fmt.Sprintf("%s_", EventName)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("event_name = ?", EventName).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "FindByEventName", fmt.Sprintf("%s_", EventName)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindOrNilByID(ID int64) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "FindOrNilByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindOrNilByName(Name string) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "FindOrNilByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) FindOrNilByNameAndEventName(Name string, EventName string) (*loginReward.LoginRewardModel, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "FindOrNilByNameAndEventName", fmt.Sprintf("%s_%s_", Name, EventName)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModel); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModel{}
	res := d.Read.Where("name = ?", Name).Where("event_name = ?", EventName).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "FindByNameAndEventName", fmt.Sprintf("%s_%s_", Name, EventName)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) List(limit int) (*loginReward.LoginRewardModels, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "List", ""))
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

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "List", ""), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) ListByEventName(EventName string) (*loginReward.LoginRewardModels, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "ListByEventName", fmt.Sprintf("%s_", EventName)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModels); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModels{}
	res := d.Read.Where("event_name = ?", EventName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "ListByEventName", fmt.Sprintf("%s_", EventName)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) ListByName(Name string) (*loginReward.LoginRewardModels, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "ListByName", fmt.Sprintf("%s_", Name)))
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

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "ListByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) ListByNameAndEventName(Name string, EventName string) (*loginReward.LoginRewardModels, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_model", "ListByNameAndEventName", fmt.Sprintf("%s_%s_", Name, EventName)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardModels); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardModels{}
	res := d.Read.Where("name = ?", Name).Where("event_name = ?", EventName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_model", "ListByNameAndEventName", fmt.Sprintf("%s_%s_", Name, EventName)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardModelDao) Save(entity *loginReward.LoginRewardModel, tx *gorm.DB) (*loginReward.LoginRewardModel, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardModel{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
