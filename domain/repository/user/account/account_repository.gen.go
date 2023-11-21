//go:generate mockgen -source=./account_repository.gen.go -destination=./account_repository_mock.gen.go -package=account
package account

import (
	"github.com/game-core/gocrafter/domain/entity/user/account"
	"github.com/jinzhu/gorm"
)

type AccountRepository interface {
	Create(entity *account.Account, tx *gorm.DB) (*account.Account, error)

	Delete(entity *account.Account, tx *gorm.DB) error

	FindByID(ID int64) (*account.Account, error)

	FindByIDAndUUID(ID int64, UUID string) (*account.Account, error)

	FindByUUID(UUID string) (*account.Account, error)

	List(limit int64) (*account.Accounts, error)

	ListByIDAndUUID(ID int64, UUID string) (*account.Accounts, error)

	ListByUUID(UUID string) (*account.Accounts, error)

	Update(entity *account.Account, tx *gorm.DB) (*account.Account, error)
}
