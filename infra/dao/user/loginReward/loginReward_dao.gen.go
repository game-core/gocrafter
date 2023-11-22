package loginReward

import (
	"github.com/jinzhu/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/user/loginReward"
	loginRewardRepository "github.com/game-core/gocrafter/domain/repository/user/loginReward"
)

type loginRewardDao struct {
	ShardConn *database.ShardConn
}

func NewLoginRewardStatusDao(conn *database.SqlHandler) loginRewardRepository.LoginRewardStatusRepository {
	return &loginRewardDao{
		ShardConn: conn.User,
	}
}

func (d *loginRewardDao) Create(entity *loginReward.LoginRewardStatus, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardStatus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardStatus{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardDao) Delete(entity *loginReward.LoginRewardStatus, shardKey int, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardStatus{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *loginRewardDao) FindByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardDao) FindByAccountIDAndLoginRewardModelID(AccountID int64, LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Where("login_reward_model_id = ?", LoginRewardModelID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardDao) FindByID(ID int64, shardKey int) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardDao) FindByLoginRewardModelID(LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_id = ?", LoginRewardModelID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardDao) List(limit int64, shardKey int) (*loginReward.LoginRewardStatuss, error) {
	entity := &loginReward.LoginRewardStatuss{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardDao) ListByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardStatuss, error) {
	entity := &loginReward.LoginRewardStatuss{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardDao) ListByAccountIDAndLoginRewardModelID(AccountID int64, LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatuss, error) {
	entity := &loginReward.LoginRewardStatuss{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Where("login_reward_model_id = ?", LoginRewardModelID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardDao) ListByLoginRewardModelID(LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatuss, error) {
	entity := &loginReward.LoginRewardStatuss{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_id = ?", LoginRewardModelID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardDao) Update(entity *loginReward.LoginRewardStatus, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardStatus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardStatus{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
