//go:generate mockgen -source=./itemBox_repository.gen.go -destination=./itemBox_repository_mock.gen.go -package=item
package item

import (
	"github.com/game-core/gocrafter/domain/entity/user/item"
	"github.com/jinzhu/gorm"
)

type ItemBoxRepository interface {
	Create(entity *item.ItemBox, shardKey int, tx *gorm.DB) (*item.ItemBox, error)

	Delete(entity *item.ItemBox, shardKey int, tx *gorm.DB) error

	FindByID(ID int64, shardKey int) (*item.ItemBox, error)

	FindOrNilByItemID(ItemID int64, shardKey int) (*item.ItemBox, error)

	FindOrNilByUserID(UserID int64, shardKey int) (*item.ItemBox, error)

	FindOrNilByUserIDAndItemID(UserID int64, ItemID int64, shardKey int) (*item.ItemBox, error)

	FindOrNilByID(ID int64, shardKey int) (*item.ItemBox, error)

	List(limit int64, shardKey int) (*item.ItemBoxs, error)

	ListByItemID(ItemID int64, shardKey int) (*item.ItemBoxs, error)

	ListByUserID(UserID int64, shardKey int) (*item.ItemBoxs, error)

	ListByUserIDAndItemID(UserID int64, ItemID int64, shardKey int) (*item.ItemBoxs, error)

	Update(entity *item.ItemBox, shardKey int, tx *gorm.DB) (*item.ItemBox, error)
}
