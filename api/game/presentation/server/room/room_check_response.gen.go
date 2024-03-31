// Package room ルーム確認レスポンス
package room

func SetRoomCheckResponse(commonRoom *CommonRoom, commonRoomUser *CommonRoomUser) *RoomCheckResponse {
	return &RoomCheckResponse{
		CommonRoom:     commonRoom,
		CommonRoomUser: commonRoomUser,
	}
}
