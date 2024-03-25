// Package room ルーム検索リクエスト
package room

func SetRoomSearchRequest(userId string, name string) *RoomSearchRequest {
	return &RoomSearchRequest{
		UserId: userId,
		Name:   name,
	}
}
