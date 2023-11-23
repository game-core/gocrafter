//go:generate mockgen -source=./event_repository.gen.go -destination=./event_repository_mock.gen.go -package=event
package event

import (
	"github.com/game-core/gocrafter/domain/entity/master/event"
	"github.com/jinzhu/gorm"
)

type EventRepository interface {
	Create(entity *event.Event, tx *gorm.DB) (*event.Event, error)

	Delete(entity *event.Event, tx *gorm.DB) error

	FindByID(ID int64) (*event.Event, error)

	FindByName(Name string) (*event.Event, error)

	FindOrNilByID(ID int64) (*event.Event, error)

	FindOrNilByName(Name string) (*event.Event, error)

	List(limit int64) (*event.Events, error)

	ListByName(Name string) (*event.Events, error)

	Update(entity *event.Event, tx *gorm.DB) (*event.Event, error)
}
