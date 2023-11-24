package loginReward

import (
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/user/loginReward"
	loginRewardRepository "github.com/game-core/gocrafter/domain/repository/user/loginReward"
)

type loginRewardReceiveStepDao struct {
	ShardConn *database.ShardConn
}

func NewLoginRewardReceiveStepDao(conn *database.SqlHandler) loginRewardRepository.LoginRewardReceiveStepRepository {
	return &loginRewardReceiveStepDao{
		ShardConn: conn.User,
	}
}

func (d *loginRewardReceiveStepDao) Create(entity *loginReward.LoginRewardReceiveStep, shardKey string, tx *gorm.DB) (*loginReward.LoginRewardReceiveStep, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardReceiveStep{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) Delete(entity *loginReward.LoginRewardReceiveStep, shardKey string, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardReceiveStep{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *loginRewardReceiveStepDao) FindByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error) {
	entity := &loginReward.LoginRewardReceiveStep{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) FindByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error) {
	entity := &loginReward.LoginRewardReceiveStep{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) FindByID(ID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error) {
	entity := &loginReward.LoginRewardReceiveStep{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) FindByLoginRewardStatusID(LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error) {
	entity := &loginReward.LoginRewardReceiveStep{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) FindOrNilByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error) {
	entity := &loginReward.LoginRewardReceiveStep{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) FindOrNilByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error) {
	entity := &loginReward.LoginRewardReceiveStep{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) FindOrNilByID(ID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error) {
	entity := &loginReward.LoginRewardReceiveStep{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) FindOrNilByLoginRewardStatusID(LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveStep, error) {
	entity := &loginReward.LoginRewardReceiveStep{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) List(limit int, shardKey string) (*loginReward.LoginRewardReceiveSteps, error) {
	entity := &loginReward.LoginRewardReceiveSteps{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) ListByAccountID(AccountID int64, shardKey string) (*loginReward.LoginRewardReceiveSteps, error) {
	entity := &loginReward.LoginRewardReceiveSteps{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) ListByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveSteps, error) {
	entity := &loginReward.LoginRewardReceiveSteps{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) ListByLoginRewardStatusID(LoginRewardStatusID int64, shardKey string) (*loginReward.LoginRewardReceiveSteps, error) {
	entity := &loginReward.LoginRewardReceiveSteps{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepDao) Save(entity *loginReward.LoginRewardReceiveStep, shardKey string, tx *gorm.DB) (*loginReward.LoginRewardReceiveStep, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardReceiveStep{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
