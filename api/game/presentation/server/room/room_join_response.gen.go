// Package room ルーム参加レスポンス
package room

func SetRoomJoinResponse(commonRoom *CommonRoom, commonRoomUser *CommonRoomUser) *RoomJoinResponse {
	return &RoomJoinResponse{
		CommonRoom:     commonRoom,
		CommonRoomUser: commonRoomUser,
	}
}
