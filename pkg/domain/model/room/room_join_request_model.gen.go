// Package room ルーム参加リクエスト
package room

type RoomJoinRequests []*RoomJoinRequest

type RoomJoinRequest struct {
	UserId string
	RoomId string
}

func NewRoomJoinRequest() *RoomJoinRequest {
	return &RoomJoinRequest{}
}

func NewRoomJoinRequests() RoomJoinRequests {
	return RoomJoinRequests{}
}

func SetRoomJoinRequest(userId string, roomId string) *RoomJoinRequest {
	return &RoomJoinRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
