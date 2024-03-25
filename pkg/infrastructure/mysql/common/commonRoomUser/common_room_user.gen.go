// Package commonRoomUser ルームユーザー
package commonRoomUser

import (
	"time"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type CommonRoomUsers []*CommonRoomUser

type CommonRoomUser struct {
	RoomId               string
	HostUserId           string
	RoomUserPositionType enum.RoomUserPositionType
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func NewCommonRoomUser() *CommonRoomUser {
	return &CommonRoomUser{}
}

func NewCommonRoomUsers() CommonRoomUsers {
	return CommonRoomUsers{}
}

func SetCommonRoomUser(roomId string, hostUserId string, roomUserPositionType enum.RoomUserPositionType, createdAt time.Time, updatedAt time.Time) *CommonRoomUser {
	return &CommonRoomUser{
		RoomId:               roomId,
		HostUserId:           hostUserId,
		RoomUserPositionType: roomUserPositionType,
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
	}
}

func (t *CommonRoomUser) TableName() string {
	return "common_room_user"
}
