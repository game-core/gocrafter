package loginReward

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"

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
	cachedResult, found := d.Cache.Get(cacheKey("FindByID", fmt.Sprintf("%d_", ID)))
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

	d.Cache.Set(cacheKey("FindByID", fmt.Sprintf("%d_", ID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) FindOrNilByItemID(ItemID int64) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByItemID", fmt.Sprintf("%d_", ItemID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
	res := d.Read.Where("item_id = ?", ItemID).Find(entity)
	if res.RecordNotFound() {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByItemID", fmt.Sprintf("%d_", ItemID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) FindOrNilByLoginRewardID(LoginRewardID int64) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByLoginRewardID", fmt.Sprintf("%d_", LoginRewardID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
	res := d.Read.Where("login_reward_id = ?", LoginRewardID).Find(entity)
	if res.RecordNotFound() {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByLoginRewardID", fmt.Sprintf("%d_", LoginRewardID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) FindOrNilByLoginRewardIDAndItemID(LoginRewardID int64, ItemID int64) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByLoginRewardIDAndItemID", fmt.Sprintf("%d_%d_", LoginRewardID, ItemID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
	res := d.Read.Where("login_reward_id = ?", LoginRewardID).Where("item_id = ?", ItemID).Find(entity)
	if res.RecordNotFound() {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("FindByLoginRewardIDAndItemID", fmt.Sprintf("%d_%d_", LoginRewardID, ItemID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) FindOrNilByName(Name string) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByName", fmt.Sprintf("%s_", Name)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
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

func (d *loginRewardRewardDao) FindOrNilByID(ID int64) (*loginReward.LoginRewardReward, error) {
	cachedResult, found := d.Cache.Get(cacheKey("FindByID", fmt.Sprintf("%d_", ID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardReward); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardReward{}
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

func (d *loginRewardRewardDao) List(limit int64) (*loginReward.LoginRewardRewards, error) {
	cachedResult, found := d.Cache.Get(cacheKey("List", ""))
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

	d.Cache.Set(cacheKey("List", ""), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) ListByItemID(ItemID int64) (*loginReward.LoginRewardRewards, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByItemID", fmt.Sprintf("%d_", ItemID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardRewards); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardRewards{}
	res := d.Read.Where("item_id = ?", ItemID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("ListByItemID", fmt.Sprintf("%d_", ItemID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) ListByLoginRewardID(LoginRewardID int64) (*loginReward.LoginRewardRewards, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByLoginRewardID", fmt.Sprintf("%d_", LoginRewardID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardRewards); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardRewards{}
	res := d.Read.Where("login_reward_id = ?", LoginRewardID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("ListByLoginRewardID", fmt.Sprintf("%d_", LoginRewardID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) ListByLoginRewardIDAndItemID(LoginRewardID int64, ItemID int64) (*loginReward.LoginRewardRewards, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByLoginRewardIDAndItemID", fmt.Sprintf("%d_%d_", LoginRewardID, ItemID)))
	if found {
		if cachedEntity, ok := cachedResult.(*loginReward.LoginRewardRewards); ok {
			return cachedEntity, nil
		}
	}

	entity := &loginReward.LoginRewardRewards{}
	res := d.Read.Where("login_reward_id = ?", LoginRewardID).Where("item_id = ?", ItemID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	d.Cache.Set(cacheKey("ListByLoginRewardIDAndItemID", fmt.Sprintf("%d_%d_", LoginRewardID, ItemID)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) ListByName(Name string) (*loginReward.LoginRewardRewards, error) {
	cachedResult, found := d.Cache.Get(cacheKey("ListByName", fmt.Sprintf("%s_", Name)))
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

	d.Cache.Set(cacheKey("ListByName", fmt.Sprintf("%s_", Name)), entity, cache.DefaultExpiration)

	return entity, nil
}

func (d *loginRewardRewardDao) Update(entity *loginReward.LoginRewardReward, tx *gorm.DB) (*loginReward.LoginRewardReward, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&loginReward.LoginRewardReward{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func cacheKey(method string, key string) string {
	return fmt.Sprintf("loginreward:%s:%s", method, key)
}
