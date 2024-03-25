// Package room ルームユーザー
package room

func SetCommonRoomUser(roomId string, userId string, roomUserPositionType RoomUserPositionType) *CommonRoomUser {
	return &CommonRoomUser{
		RoomId:               roomId,
		UserId:               userId,
		RoomUserPositionType: roomUserPositionType,
	}
}
