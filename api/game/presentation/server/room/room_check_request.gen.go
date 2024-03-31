// Package room ルーム確認リクエスト
package room

func SetRoomCheckRequest(userId string, roomId string) *RoomCheckRequest {
	return &RoomCheckRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
