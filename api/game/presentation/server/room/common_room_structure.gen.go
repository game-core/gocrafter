// Package room ルーム
package room

func SetCommonRoom(roomId string, hostUserId string, roomReleaseType RoomReleaseType, name string, userCount int32) *CommonRoom {
	return &CommonRoom{
		RoomId:          roomId,
		HostUserId:      hostUserId,
		RoomReleaseType: roomReleaseType,
		Name:            name,
		UserCount:       userCount,
	}
}
