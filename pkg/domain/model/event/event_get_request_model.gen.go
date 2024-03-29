// Package event イベント取得リクエスト
package event

type EventGetRequests []*EventGetRequest

type EventGetRequest struct {
	EventId int64
}

func NewEventGetRequest() *EventGetRequest {
	return &EventGetRequest{}
}

func NewEventGetRequests() EventGetRequests {
	return EventGetRequests{}
}

func SetEventGetRequest(eventId int64) *EventGetRequest {
	return &EventGetRequest{
		EventId: eventId,
	}
}
