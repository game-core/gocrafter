package account

import (
	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/user/account"
	accountRepository "github.com/game-core/gocrafter/domain/repository/user/account"
	"github.com/jinzhu/gorm"
)

type accountDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewAccountDao(conn *database.SqlHandler) accountRepository.AccountRepository {
	return &accountDao{
		Read:  conn.User.ReadConn,
		Write: conn.User.WriteConn,
	}
}

func (d *accountDao) Create(entity *account.Account, tx *gorm.DB) (*account.Account, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&account.Account{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) Delete(entity *account.Account, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&account.Account{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *accountDao) FindByID(ID int64) (*account.Account, error) {
	entity := &account.Account{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindByIDAndUUID(ID int64, UUID string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.Read.Where("id = ?", ID).Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindByUUID(UUID string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.Read.Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) List(limit int64) (*account.Accounts, error) {
	entity := &account.Accounts{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) ListByIDAndUUID(ID int64, UUID string) (*account.Accounts, error) {
	entity := &account.Accounts{}
	res := d.Read.Where("id = ?", ID).Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) ListByUUID(UUID string) (*account.Accounts, error) {
	entity := &account.Accounts{}
	res := d.Read.Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) Update(entity *account.Account, tx *gorm.DB) (*account.Account, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&account.Account{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
