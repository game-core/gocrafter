//go:generate mockgen -source=./item_repository.gen.go -destination=./item_repository_mock.gen.go -package=item
package item

import (
	"github.com/game-core/gocrafter/domain/entity/master/item"
	"github.com/jinzhu/gorm"
)

type ItemRepository interface {
	Create(entity *item.Item, tx *gorm.DB) (*item.Item, error)

	Delete(entity *item.Item, tx *gorm.DB) error

	FindByID(ID int64) (*item.Item, error)

	FindByName(Name string) (*item.Item, error)

	FindOrNilByID(ID int64) (*item.Item, error)

	FindOrNilByName(Name string) (*item.Item, error)

	List(limit int64) (*item.Items, error)

	ListByName(Name string) (*item.Items, error)

	Update(entity *item.Item, tx *gorm.DB) (*item.Item, error)
}
