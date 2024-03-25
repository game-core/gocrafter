// Package room ルーム参加リクエスト
package room

type RoomCheckInRequests []*RoomCheckInRequest

type RoomCheckInRequest struct {
	UserId string
	RoomId string
}

func NewRoomCheckInRequest() *RoomCheckInRequest {
	return &RoomCheckInRequest{}
}

func NewRoomCheckInRequests() RoomCheckInRequests {
	return RoomCheckInRequests{}
}

func SetRoomCheckInRequest(userId string, roomId string) *RoomCheckInRequest {
	return &RoomCheckInRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
