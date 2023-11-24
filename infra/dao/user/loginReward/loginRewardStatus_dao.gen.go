package loginReward

import (
	"gorm.io/gorm"

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

func (d *loginRewardStatusDao) Create(entity *loginReward.LoginRewardStatus, shardKey string, tx *gorm.DB) (*loginReward.LoginRewardStatus, error) {
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

func (d *loginRewardStatusDao) Delete(entity *loginReward.LoginRewardStatus, shardKey string, tx *gorm.DB) error {
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

func (d *loginRewardStatusDao) FindByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindByAccountIDAndLoginRewardModelName(AccountID int64, LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Where("login_reward_model_Name = ?", LoginRewardModelName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindByID(ID int64, shardKey string) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindByLoginRewardModelName(LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_Name = ?", LoginRewardModelName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindOrNilByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindOrNilByAccountIDAndLoginRewardModelName(AccountID int64, LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Where("login_reward_model_Name = ?", LoginRewardModelName).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindOrNilByID(ID int64, shardKey string) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) FindOrNilByLoginRewardModelName(LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatus, error) {
	entity := &loginReward.LoginRewardStatus{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_Name = ?", LoginRewardModelName).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) List(limit int, shardKey string) (*loginReward.LoginRewardStatuses, error) {
	entity := &loginReward.LoginRewardStatuses{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) ListByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardStatuses, error) {
	entity := &loginReward.LoginRewardStatuses{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) ListByAccountIDAndLoginRewardModelName(AccountID int64, LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatuses, error) {
	entity := &loginReward.LoginRewardStatuses{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Where("login_reward_model_Name = ?", LoginRewardModelName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) ListByLoginRewardModelName(LoginRewardModelName string, shardKey string) (*loginReward.LoginRewardStatuses, error) {
	entity := &loginReward.LoginRewardStatuses{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_Name = ?", LoginRewardModelName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardStatusDao) Save(entity *loginReward.LoginRewardStatus, shardKey string, tx *gorm.DB) (*loginReward.LoginRewardStatus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardStatus{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
