// Package room ルーム確認リクエスト
package room

type RoomCheckRequests []*RoomCheckRequest

type RoomCheckRequest struct {
	UserId string
	RoomId string
}

func NewRoomCheckRequest() *RoomCheckRequest {
	return &RoomCheckRequest{}
}

func NewRoomCheckRequests() RoomCheckRequests {
	return RoomCheckRequests{}
}

func SetRoomCheckRequest(userId string, roomId string) *RoomCheckRequest {
	return &RoomCheckRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
