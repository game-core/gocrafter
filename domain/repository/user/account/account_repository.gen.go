//go:generate mockgen -source=./account_repository.gen.go -destination=./account_repository_mock.gen.go -package=account
package account

import (
	"github.com/game-core/gocrafter/domain/entity/user/account"
	"github.com/jinzhu/gorm"
)

type AccountRepository interface {
	Create(entity *account.Account, shardKey int, tx *gorm.DB) (*account.Account, error)

	Delete(entity *account.Account, shardKey int, tx *gorm.DB) error

	FindByID(ID int64, shardKey int) (*account.Account, error)

	FindOrNilByIDAndUUID(ID int64, UUID string, shardKey int) (*account.Account, error)

	FindOrNilByUUID(UUID string, shardKey int) (*account.Account, error)

	FindOrNilByID(ID int64, shardKey int) (*account.Account, error)

	List(limit int64, shardKey int) (*account.Accounts, error)

	ListByIDAndUUID(ID int64, UUID string, shardKey int) (*account.Accounts, error)

	ListByUUID(UUID string, shardKey int) (*account.Accounts, error)

	Update(entity *account.Account, shardKey int, tx *gorm.DB) (*account.Account, error)
}
