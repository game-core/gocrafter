// Package room ルーム退出リクエスト
package room

type RoomCheckOutRequests []*RoomCheckOutRequest

type RoomCheckOutRequest struct {
	UserId string
	RoomId string
}

func NewRoomCheckOutRequest() *RoomCheckOutRequest {
	return &RoomCheckOutRequest{}
}

func NewRoomCheckOutRequests() RoomCheckOutRequests {
	return RoomCheckOutRequests{}
}

func SetRoomCheckOutRequest(userId string, roomId string) *RoomCheckOutRequest {
	return &RoomCheckOutRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
