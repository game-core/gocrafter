//go:generate mockgen -source=./itemBox_repository.gen.go -destination=./itemBox_repository_mock.gen.go -package=item
package item

import (
	"github.com/game-core/gocrafter/domain/entity/user/item"
	"gorm.io/gorm"
)

type ItemBoxRepository interface {
	Create(entity *item.ItemBox, shardKey string, tx *gorm.DB) (*item.ItemBox, error)

	Delete(entity *item.ItemBox, shardKey string, tx *gorm.DB) error

	FindByAccountID(AccountID int64, shardKey string) (*item.ItemBox, error)

	FindByAccountIDAndItemName(AccountID int64, ItemName string, shardKey string) (*item.ItemBox, error)

	FindByID(ID int64, shardKey string) (*item.ItemBox, error)

	FindByItemName(ItemName string, shardKey string) (*item.ItemBox, error)

	FindOrNilByAccountID(AccountID int64, shardKey string) (*item.ItemBox, error)

	FindOrNilByAccountIDAndItemName(AccountID int64, ItemName string, shardKey string) (*item.ItemBox, error)

	FindOrNilByID(ID int64, shardKey string) (*item.ItemBox, error)

	FindOrNilByItemName(ItemName string, shardKey string) (*item.ItemBox, error)

	List(limit int, shardKey string) (*item.ItemBoxs, error)

	ListByAccountID(AccountID int64, shardKey string) (*item.ItemBoxs, error)

	ListByAccountIDAndItemName(AccountID int64, ItemName string, shardKey string) (*item.ItemBoxs, error)

	ListByItemName(ItemName string, shardKey string) (*item.ItemBoxs, error)

	Save(entity *item.ItemBox, shardKey string, tx *gorm.DB) (*item.ItemBox, error)
}
