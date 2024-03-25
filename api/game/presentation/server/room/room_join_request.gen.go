// Package room ルーム参加リクエスト
package room

func SetRoomJoinRequest(userId string, roomId string) *RoomJoinRequest {
	return &RoomJoinRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
