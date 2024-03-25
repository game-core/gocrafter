// Package room ルーム退出リクエスト
package room

func SetRoomCheckOutRequest(userId string, roomId string) *RoomCheckOutRequest {
	return &RoomCheckOutRequest{
		UserId: userId,
		RoomId: roomId,
	}
}
