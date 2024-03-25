// Package commonRoomUser ルームユーザー
package commonRoomUser

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type CommonRoomUsers []*CommonRoomUser

type CommonRoomUser struct {
	RoomId               string
	UserId               string
	RoomUserPositionType enum.RoomUserPositionType
}

func NewCommonRoomUser() *CommonRoomUser {
	return &CommonRoomUser{}
}

func NewCommonRoomUsers() CommonRoomUsers {
	return CommonRoomUsers{}
}

func SetCommonRoomUser(roomId string, userId string, roomUserPositionType enum.RoomUserPositionType) *CommonRoomUser {
	return &CommonRoomUser{
		RoomId:               roomId,
		UserId:               userId,
		RoomUserPositionType: roomUserPositionType,
	}
}
