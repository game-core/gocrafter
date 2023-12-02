//go:generate mockgen -source=./account_repository.gen.go -destination=./account_repository_mock.gen.go -package=account
package account

import (
	"github.com/game-core/gocrafter/domain/entity/auth/account"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(entity *account.Account, tx *gorm.DB) (*account.Account, error)

	Delete(entity *account.Account, tx *gorm.DB) error

	FindByEmail(Email string) (*account.Account, error)

	FindByID(ID int64) (*account.Account, error)

	FindOrNilByEmail(Email string) (*account.Account, error)

	FindOrNilByID(ID int64) (*account.Account, error)

	List(limit int) (*account.Accounts, error)

	ListByEmail(Email string) (*account.Accounts, error)

	Save(entity *account.Account, tx *gorm.DB) (*account.Account, error)
}
