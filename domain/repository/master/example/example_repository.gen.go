package example

import (
	"github.com/game-core/gocrafter/domain/entity/master/example"
	"github.com/jinzhu/gorm"
)

type ExampleRepository interface {
	Create(example *example.Example, tx *gorm.DB) (*example.Example, error)

	Delete(example *example.Example, tx *gorm.DB) error

	FindByID(ID int64) (*example.Example, error)

	FindByIDAndName(ID int64, Name string) (*example.Example, error)

	FindByName(Name string) (*example.Example, error)

	List(limit int64) (*example.Examples, error)

	Update(example *example.Example, tx *gorm.DB) (*example.Example, error)
}
