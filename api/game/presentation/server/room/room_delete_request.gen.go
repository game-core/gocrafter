// Package room ルーム削除リクエスト
package room

func SetRoomDeleteRequest(userId string, roomId string) *RoomDeleteRequest {
	return &RoomDeleteRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
