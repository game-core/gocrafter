// Package room ルーム作成リクエスト
package room

func SetRoomCreateRequest(userId string, name string, roomReleaseType RoomReleaseType) *RoomCreateRequest {
	return &RoomCreateRequest{
		UserId:          userId,
		Name:            name,
		RoomReleaseType: roomReleaseType,
	}
}
