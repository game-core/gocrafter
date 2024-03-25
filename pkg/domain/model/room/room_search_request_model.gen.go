// Package room ルーム検索リクエスト
package room

type RoomSearchRequests []*RoomSearchRequest

type RoomSearchRequest struct {
	UserId string
	Name   string
}

func NewRoomSearchRequest() *RoomSearchRequest {
	return &RoomSearchRequest{}
}

func NewRoomSearchRequests() RoomSearchRequests {
	return RoomSearchRequests{}
}

func SetRoomSearchRequest(userId string, name string) *RoomSearchRequest {
	return &RoomSearchRequest{
		UserId: userId,
		Name:   name,
	}
}
