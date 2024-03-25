// Package room ルーム退出レスポンス
package room

func SetRoomCheckOutResponse(commonRoom *CommonRoom, commonRoomUser *CommonRoomUser) *RoomCheckOutResponse {
	return &RoomCheckOutResponse{
		CommonRoom:     commonRoom,
		CommonRoomUser: commonRoomUser,
	}
}
