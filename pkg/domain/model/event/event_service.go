//go:generate mockgen -source=./event_service.go -destination=./event_service_mock.gen.go -package=event
package event

import (
	"context"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/event/masterEvent"
)

type EventService interface {
	Get(cxt context.Context, req *EventGetRequest) (*EventGetResponse, error)
}

type eventService struct {
	masterEventMysqlRepository masterEvent.MasterEventMysqlRepository
}

func NewEventService(
	masterEventMysqlRepository masterEvent.MasterEventMysqlRepository,
) EventService {
	return &eventService{
		masterEventMysqlRepository: masterEventMysqlRepository,
	}
}

// Get イベントを取得する
func (s *eventService) Get(cxt context.Context, req *EventGetRequest) (*EventGetResponse, error) {
	result, err := s.masterEventMysqlRepository.Find(cxt, req.MasterEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterEventMysqlRepository.Find", err)
	}

	return SetEventGetResponse(result), nil
}
