// Package room ルーム参加レスポンス
package room

func SetRoomCheckInResponse(commonRoom *CommonRoom, commonRoomUser *CommonRoomUser) *RoomCheckInResponse {
	return &RoomCheckInResponse{
		CommonRoom:     commonRoom,
		CommonRoomUser: commonRoomUser,
	}
}
