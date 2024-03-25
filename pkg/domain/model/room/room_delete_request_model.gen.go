// Package room ルーム削除リクエスト
package room

type RoomDeleteRequests []*RoomDeleteRequest

type RoomDeleteRequest struct {
	UserId string
	RoomId string
}

func NewRoomDeleteRequest() *RoomDeleteRequest {
	return &RoomDeleteRequest{}
}

func NewRoomDeleteRequests() RoomDeleteRequests {
	return RoomDeleteRequests{}
}

func SetRoomDeleteRequest(userId string, roomId string) *RoomDeleteRequest {
	return &RoomDeleteRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
