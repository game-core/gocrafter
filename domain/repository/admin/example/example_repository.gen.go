//go:generate mockgen -source=./example_repository.gen.go -destination=./example_repository_mock.gen.go -package=example
package example

import (
	"github.com/game-core/gocrafter/domain/entity/admin/example"
	"gorm.io/gorm"
)

type ExampleRepository interface {
	Create(entity *example.Example, tx *gorm.DB) (*example.Example, error)

	Delete(entity *example.Example, tx *gorm.DB) error

	FindByID(ID int64) (*example.Example, error)

	FindByIDAndName(ID int64, Name string) (*example.Example, error)

	FindByName(Name string) (*example.Example, error)

	FindOrNilByID(ID int64) (*example.Example, error)

	FindOrNilByIDAndName(ID int64, Name string) (*example.Example, error)

	FindOrNilByName(Name string) (*example.Example, error)

	List(limit int) (*example.Examples, error)

	ListByIDAndName(ID int64, Name string) (*example.Examples, error)

	ListByName(Name string) (*example.Examples, error)

	Save(entity *example.Example, tx *gorm.DB) (*example.Example, error)
}
