// Package commonRoom ルーム
package commonRoom

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type CommonRooms []*CommonRoom

type CommonRoom struct {
	RoomId          string
	HostUserId      string
	RoomReleaseType enum.RoomReleaseType
	Name            string
	UserCount       int32
}

func NewCommonRoom() *CommonRoom {
	return &CommonRoom{}
}

func NewCommonRooms() CommonRooms {
	return CommonRooms{}
}

func SetCommonRoom(roomId string, hostUserId string, roomReleaseType enum.RoomReleaseType, name string, userCount int32) *CommonRoom {
	return &CommonRoom{
		RoomId:          roomId,
		HostUserId:      hostUserId,
		RoomReleaseType: roomReleaseType,
		Name:            name,
		UserCount:       userCount,
	}
}
