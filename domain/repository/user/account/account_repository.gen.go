package account

import (
	"github.com/game-core/gocrafter/domain/model/user/account"
	"github.com/jinzhu/gorm"
)

type AccountRepository interface {
	Create(account *account.Account, tx *gorm.DB) (*account.Account, error)

	Delete(account *account.Account, tx *gorm.DB) error

	FindByID(ID int64) (*account.Account, error)

	FindByIDAndUUID(ID int64, UUID int64) (*account.Account, error)

	FindByUUID(UUID int64) (*account.Account, error)

	List(limit int64) (*account.Accounts, error)

	Update(account *account.Account, tx *gorm.DB) (*account.Account, error)
}
