package account

import (
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/auth/account"
	accountRepository "github.com/game-core/gocrafter/domain/repository/auth/account"
)

type accountDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewAccountDao(conn *database.SqlHandler) accountRepository.AccountRepository {
	return &accountDao{
		Read:  conn.Auth.ReadConn,
		Write: conn.Auth.WriteConn,
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

func (d *accountDao) FindByEmail(Email string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.Read.Where("email = ?", Email).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindByID(ID int64) (*account.Account, error) {
	entity := &account.Account{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindOrNilByEmail(Email string) (*account.Account, error) {
	entity := &account.Account{}
	res := d.Read.Where("email = ?", Email).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) FindOrNilByID(ID int64) (*account.Account, error) {
	entity := &account.Account{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) List(limit int) (*account.Accounts, error) {
	entity := &account.Accounts{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) ListByEmail(Email string) (*account.Accounts, error) {
	entity := &account.Accounts{}
	res := d.Read.Where("email = ?", Email).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *accountDao) Save(entity *account.Account, tx *gorm.DB) (*account.Account, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&account.Account{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
