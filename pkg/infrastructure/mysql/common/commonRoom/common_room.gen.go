// Package commonRoom ルーム
package commonRoom

import (
	"time"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type CommonRooms []*CommonRoom

type CommonRoom struct {
	RoomId          string
	HostUserId      string
	RoomReleaseType enum.RoomReleaseType
	Name            string
	UserCount       int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewCommonRoom() *CommonRoom {
	return &CommonRoom{}
}

func NewCommonRooms() CommonRooms {
	return CommonRooms{}
}

func SetCommonRoom(roomId string, hostUserId string, roomReleaseType enum.RoomReleaseType, name string, userCount int32, createdAt time.Time, updatedAt time.Time) *CommonRoom {
	return &CommonRoom{
		RoomId:          roomId,
		HostUserId:      hostUserId,
		RoomReleaseType: roomReleaseType,
		Name:            name,
		UserCount:       userCount,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}

func (t *CommonRoom) TableName() string {
	return "common_room"
}
