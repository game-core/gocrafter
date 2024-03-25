// Package room ルーム検索リクエスト
package room

func SetRoomSearchResponse(commonRooms []*CommonRoom) *RoomSearchResponse {
	return &RoomSearchResponse{
		CommonRooms: commonRooms,
	}
}
