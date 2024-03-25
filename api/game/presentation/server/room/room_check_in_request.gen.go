// Package room ルーム参加リクエスト
package room

func SetRoomCheckInRequest(userId string, roomId string) *RoomCheckInRequest {
	return &RoomCheckInRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
