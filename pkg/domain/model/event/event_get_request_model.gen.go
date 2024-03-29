// Package event イベント取得リクエスト
package event

type EventGetRequests []*EventGetRequest

type EventGetRequest struct {
	MasterEventId int64
}

func NewEventGetRequest() *EventGetRequest {
	return &EventGetRequest{}
}

func NewEventGetRequests() EventGetRequests {
	return EventGetRequests{}
}

func SetEventGetRequest(masterEventId int64) *EventGetRequest {
	return &EventGetRequest{
		MasterEventId: masterEventId,
	}
}
