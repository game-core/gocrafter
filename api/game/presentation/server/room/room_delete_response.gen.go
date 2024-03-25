// Package room ルーム削除レスポンス
package room

func SetRoomDeleteResponse(commonRoom *CommonRoom) *RoomDeleteResponse {
	return &RoomDeleteResponse{
		CommonRoom: commonRoom,
	}
}
