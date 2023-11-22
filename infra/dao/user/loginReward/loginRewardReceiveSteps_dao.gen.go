package loginReward

import (
	"github.com/jinzhu/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/user/loginReward"
	loginRewardRepository "github.com/game-core/gocrafter/domain/repository/user/loginReward"
)

type loginRewardReceiveStepsDao struct {
	ShardConn *database.ShardConn
}

func NewLoginRewardReceiveStepsDao(conn *database.SqlHandler) loginRewardRepository.LoginRewardReceiveStepsRepository {
	return &loginRewardReceiveStepsDao{
		ShardConn: conn.User,
	}
}

func (d *loginRewardReceiveStepsDao) Create(entity *loginReward.LoginRewardReceiveSteps, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardReceiveSteps, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardReceiveSteps{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepsDao) Delete(entity *loginReward.LoginRewardReceiveSteps, shardKey int, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardReceiveSteps{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *loginRewardReceiveStepsDao) FindByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error) {
	entity := &loginReward.LoginRewardReceiveSteps{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepsDao) FindByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error) {
	entity := &loginReward.LoginRewardReceiveSteps{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepsDao) FindByID(ID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error) {
	entity := &loginReward.LoginRewardReceiveSteps{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepsDao) FindByLoginRewardStatusID(LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveSteps, error) {
	entity := &loginReward.LoginRewardReceiveSteps{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepsDao) List(limit int64, shardKey int) (*loginReward.LoginRewardReceiveStepss, error) {
	entity := &loginReward.LoginRewardReceiveStepss{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepsDao) ListByAccountID(AccountID int64, shardKey int) (*loginReward.LoginRewardReceiveStepss, error) {
	entity := &loginReward.LoginRewardReceiveStepss{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepsDao) ListByAccountIDAndLoginRewardStatusID(AccountID int64, LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStepss, error) {
	entity := &loginReward.LoginRewardReceiveStepss{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("user_id = ?", AccountID).Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepsDao) ListByLoginRewardStatusID(LoginRewardStatusID int64, shardKey int) (*loginReward.LoginRewardReceiveStepss, error) {
	entity := &loginReward.LoginRewardReceiveStepss{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("login_reward_model_id = ?", LoginRewardStatusID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *loginRewardReceiveStepsDao) Update(entity *loginReward.LoginRewardReceiveSteps, shardKey int, tx *gorm.DB) (*loginReward.LoginRewardReceiveSteps, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&loginReward.LoginRewardReceiveSteps{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
