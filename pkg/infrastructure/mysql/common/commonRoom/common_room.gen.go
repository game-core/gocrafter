// Package commonRoom ルーム
package commonRoom

import (
	"time"
)

type CommonRooms []*CommonRoom

type CommonRoom struct {
	RoomId     string
	HostUserId string
	RoomNumber int32
	Name       string
	UserCount  int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewCommonRoom() *CommonRoom {
	return &CommonRoom{}
}

func NewCommonRooms() CommonRooms {
	return CommonRooms{}
}

func SetCommonRoom(roomId string, hostUserId string, roomNumber int32, name string, userCount int32, createdAt time.Time, updatedAt time.Time) *CommonRoom {
	return &CommonRoom{
		RoomId:     roomId,
		HostUserId: hostUserId,
		RoomNumber: roomNumber,
		Name:       name,
		UserCount:  userCount,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}

func (t *CommonRoom) TableName() string {
	return "common_room"
}
