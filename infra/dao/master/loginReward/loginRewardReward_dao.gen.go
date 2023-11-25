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

type loginRewardRewardDao struct {
	Read  *gorm.DB
	Write *gorm.DB
	Cache *cache.Cache
}

func NewLoginRewardRewardDao(conn *database.SqlHandler) loginRewardRepository.LoginRewardRewardRepository {
	return &loginRewardRewardDao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
		Cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (d *loginRewardRewardDao) Create(entity *loginReward.LoginRewardReward, tx *gorm.DB) (*loginReward.LoginRewardReward, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardReward{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardRewardDao) Delete(entity *loginReward.LoginRewardReward, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardReward{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *loginRewardRewardDao) FindByID(ID int64) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_reward", "FindByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_reward", "FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) FindByLoginRewardModelName(LoginRewardModelName string) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_reward", "FindByLoginRewardModelName", fmt.Sprintf("%s_", LoginRewardModelName)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
	res := d.Read.Where("login_reward_model_name = ?", LoginRewardModelName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_reward", "FindByLoginRewardModelName", fmt.Sprintf("%s_", LoginRewardModelName)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) FindByName(Name string) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_reward", "FindByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_reward", "FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) FindOrNilByID(ID int64) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_reward", "FindOrNilByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_reward", "FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) FindOrNilByLoginRewardModelName(LoginRewardModelName string) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_reward", "FindOrNilByLoginRewardModelName", fmt.Sprintf("%s_", LoginRewardModelName)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
	res := d.Read.Where("login_reward_model_name = ?", LoginRewardModelName).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_reward", "FindByLoginRewardModelName", fmt.Sprintf("%s_", LoginRewardModelName)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) FindOrNilByName(Name string) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_reward", "FindOrNilByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_reward", "FindByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) List(limit int) (*loginReward.LoginRewardRewards, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_reward", "List", ""))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardRewards); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardRewards{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_reward", "List", ""), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) ListByLoginRewardModelName(LoginRewardModelName string) (*loginReward.LoginRewardRewards, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_reward", "ListByLoginRewardModelName", fmt.Sprintf("%s_", LoginRewardModelName)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardRewards); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardRewards{}
	res := d.Read.Where("login_reward_model_name = ?", LoginRewardModelName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_reward", "ListByLoginRewardModelName", fmt.Sprintf("%s_", LoginRewardModelName)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) ListByName(Name string) (*loginReward.LoginRewardRewards, error) {
	cachedResult, found := d.Cache.Get(dbChashe.CreateCacheKey("login_reward_reward", "ListByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardRewards); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardRewards{}
	res := d.Read.Where("name = ?", Name).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(dbChashe.CreateCacheKey("login_reward_reward", "ListByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) Save(entity *loginReward.LoginRewardReward, tx *gorm.DB) (*loginReward.LoginRewardReward, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardReward{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
