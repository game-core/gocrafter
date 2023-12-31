package account

import (
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/user/account"
	accountRepository "github.com/game-core/gocrafter/domain/repository/user/account"
)

type accountDao struct {
	ShardConn *database.ShardConn
}

func NewAccountDao(conn *database.SqlHandler) accountRepository.AccountRepository {
	return &accountDao{
		ShardConn: conn.User,
	}
}

func (d *accountDao) Create(entity *account.Account, shardKey string, tx *gorm.DB) (*account.Account, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&account.Account{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) Delete(entity *account.Account, shardKey string, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&account.Account{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *accountDao) FindByID(ID int64, shardKey string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindByIDAndUUID(ID int64, UUID string, shardKey string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindByUUID(UUID string, shardKey string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindOrNilByID(ID int64, shardKey string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindOrNilByIDAndUUID(ID int64, UUID string, shardKey string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Where("uuid = ?", UUID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindOrNilByUUID(UUID string, shardKey string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("uuid = ?", UUID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) List(limit int, shardKey string) (*account.Accounts, error) {
	entity := &account.Accounts{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) ListByIDAndUUID(ID int64, UUID string, shardKey string) (*account.Accounts, error) {
	entity := &account.Accounts{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) ListByUUID(UUID string, shardKey string) (*account.Accounts, error) {
	entity := &account.Accounts{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) Save(entity *account.Account, shardKey string, tx *gorm.DB) (*account.Account, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&account.Account{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
