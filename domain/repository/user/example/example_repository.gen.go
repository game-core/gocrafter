package example

import (
	"github.com/game-core/gocrafter/domain/model/user/example"
	"github.com/jinzhu/gorm"
)

type ExampleRepository interface {
	Create(example *example.Example, tx *gorm.DB) (*example.Example, error)

	Delete(example *example.Example, tx *gorm.DB) error

	FindByID(ID int64) (*example.Example, error)

	FindByIDAndUserID(ID int64, UserID int64) (*example.Example, error)

	FindByUserID(UserID int64) (*example.Example, error)

	List(limit int64) (*example.Examples, error)

	Update(example *example.Example, tx *gorm.DB) (*example.Example, error)
}
