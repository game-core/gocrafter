package loginReward

import (
	"github.com/jinzhu/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/user/loginReward"
	loginRewardRepository "github.com/game-core/gocrafter/domain/repository/user/loginReward"
)

type loginRewardStatusDao struct {
	ShardConn *database.ShardConn
}

func NewLoginRewardStatusDao(conn *database.SqlHandler) loginRewardRepository.LoginRewardStatusRepository {
	return &loginRewardStatusDao{
		ShardConn: conn.User,
	}
}

func (d *loginRewardStatusDao) Create(entity *loginReward.LoginRewardStatus, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardStatus, error) {
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

func (d *loginRewardStatusDao) Delete(entity *loginReward.LoginRewardStatus, shardKey int, tx *gorm.DB) error {
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

func (d *loginRewardStatusDao) FindByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindByAccountIDAndLoginRewardModelID(AccountID int64, LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Where("login_reward_model_id = ?", LoginRewardModelID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindByID(ID int64, shardKey int) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindByLoginRewardModelID(LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_id = ?", LoginRewardModelID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) List(limit int64, shardKey int) (*loginReward.LoginRewardStatuses, error) {
	entity := &loginReward.LoginRewardStatuses{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) ListByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardStatuses, error) {
	entity := &loginReward.LoginRewardStatuses{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) ListByAccountIDAndLoginRewardModelID(AccountID int64, LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatuses, error) {
	entity := &loginReward.LoginRewardStatuses{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Where("login_reward_model_id = ?", LoginRewardModelID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) ListByLoginRewardModelID(LoginRewardModelID int64, shardKey int) (*loginReward.LoginRewardStatuses, error) {
	entity := &loginReward.LoginRewardStatuses{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_id = ?", LoginRewardModelID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) Update(entity *loginReward.LoginRewardStatus, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardStatus, error) {
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
