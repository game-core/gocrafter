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

type loginRewardItemDao struct {
	Read  *gorm.DB
	Write *gorm.DB
	Cache *cache.Cache
}

func NewLoginRewardItemDao(conn *database.SqlHandler) loginRewardRepository.LoginRewardItemRepository {
	return &loginRewardItemDao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
		Cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (d *loginRewardItemDao) Create(entity *loginReward.LoginRewardItem, tx *gorm.DB) (*loginReward.LoginRewardItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardItem{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardItemDao) Delete(entity *loginReward.LoginRewardItem, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardItem{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *loginRewardItemDao) FindByID(ID int64) (*loginReward.LoginRewardItem, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_item", "FindByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardItem); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardItem{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_item", "FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardItemDao) FindByName(Name string) (*loginReward.LoginRewardItem, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_item", "FindByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardItem); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardItem{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_item", "FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardItemDao) FindOrNilByID(ID int64) (*loginReward.LoginRewardItem, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_item", "FindOrNilByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardItem); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardItem{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_item", "FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardItemDao) FindOrNilByName(Name string) (*loginReward.LoginRewardItem, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_item", "FindOrNilByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardItem); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardItem{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_item", "FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardItemDao) List(limit int) (*loginReward.LoginRewardItems, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_item", "List", ""))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardItems); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardItems{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_item", "List", ""), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardItemDao) ListByName(Name string) (*loginReward.LoginRewardItems, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_item", "ListByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardItems); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardItems{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_item", "ListByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardItemDao) Save(entity *loginReward.LoginRewardItem, tx *gorm.DB) (*loginReward.LoginRewardItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardItem{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
