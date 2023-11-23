//go:generate mockgen -source=./itemBox_repository.gen.go -destination=./itemBox_repository_mock.gen.go -package=item
package item

import (
	"github.com/game-core/gocrafter/domain/entity/user/item"
	"github.com/jinzhu/gorm"
)

type ItemBoxRepository interface {
	Create(entity *item.ItemBox, shardKey int, tx *gorm.DB) (*item.ItemBox, error)

	Delete(entity *item.ItemBox, shardKey int, tx *gorm.DB) error

	FindByAccountID(AccountID int64, shardKey int) (*item.ItemBox, error)

	FindByAccountIDAndItemName(AccountID int64, ItemName string, shardKey int) (*item.ItemBox, error)

	FindByID(ID int64, shardKey int) (*item.ItemBox, error)

	FindByItemName(ItemName string, shardKey int) (*item.ItemBox, error)

	FindOrNilByAccountID(AccountID int64, shardKey int) (*item.ItemBox, error)

	FindOrNilByAccountIDAndItemName(AccountID int64, ItemName string, shardKey int) (*item.ItemBox, error)

	FindOrNilByID(ID int64, shardKey int) (*item.ItemBox, error)

	FindOrNilByItemName(ItemName string, shardKey int) (*item.ItemBox, error)

	List(limit int64, shardKey int) (*item.ItemBoxs, error)

	ListByAccountID(AccountID int64, shardKey int) (*item.ItemBoxs, error)

	ListByAccountIDAndItemName(AccountID int64, ItemName string, shardKey int) (*item.ItemBoxs, error)

	ListByItemName(ItemName string, shardKey int) (*item.ItemBoxs, error)

	Update(entity *item.ItemBox, shardKey int, tx *gorm.DB) (*item.ItemBox, error)
}
