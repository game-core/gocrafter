package account

import (
	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/model/user/account"
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

func (e *accountDao) Create(entity *account.Account, tx *gorm.DB) (*account.Account, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = e.Write
	}

	res := conn.Model(&account.Account{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *accountDao) Delete(entity *account.Account, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = e.Write
	}

	res := conn.Model(&account.Account{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (e *accountDao) FindByID(ID int64) (*account.Account, error) {
	entity := &account.Account{}
	res := e.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *accountDao) FindByIDAndUUID(ID int64, UUID int64) (*account.Account, error) {
	entity := &account.Account{}
	res := e.Read.Where("id = ?", ID).Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *accountDao) FindByUUID(UUID int64) (*account.Account, error) {
	entity := &account.Account{}
	res := e.Read.Where("uuid = ?", UUID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *accountDao) List(limit int64) (*account.Accounts, error) {
	entity := &account.Accounts{}
	res := e.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (e *accountDao) Update(entity *account.Account, tx *gorm.DB) (*account.Account, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = e.Write
	}

	res := conn.Model(&account.Account{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
