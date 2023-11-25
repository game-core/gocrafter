//go:generate mockgen -source=./event_service.go -destination=./event_service_mock.gen.go -package=event
package event

import (
	"time"

	eventEntity "github.com/game-core/gocrafter/domain/entity/master/event"
	eventRepository "github.com/game-core/gocrafter/domain/repository/master/event"
)

type EventService interface {
	GetEventToEntity(name string, now time.Time) (*eventEntity.Event, error)
}

type eventService struct {
	eventRepository eventRepository.EventRepository
}

func NewEventService(
	eventRepository eventRepository.EventRepository,
) EventService {
	return &eventService{
		eventRepository: eventRepository,
	}
}

// GetEventToEntity イベントをEntityで取得する
func (s *eventService) GetEventToEntity(name string, now time.Time) (*eventEntity.Event, error) {
	e, err := s.eventRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return e, nil
}
