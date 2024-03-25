// Package room ルーム作成レスポンス
package room

func SetRoomCreateResponse(commonRoom *CommonRoom) *RoomCreateResponse {
	return &RoomCreateResponse{
		CommonRoom: commonRoom,
	}
}
