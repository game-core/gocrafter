// Package room ルーム退出レスポンス
package room

import (
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
)

type RoomCheckOutResponses []*RoomCheckOutResponse

type RoomCheckOutResponse struct {
	CommonRoom     *commonRoom.CommonRoom
	CommonRoomUser *commonRoomUser.CommonRoomUser
}

func NewRoomCheckOutResponse() *RoomCheckOutResponse {
	return &RoomCheckOutResponse{}
}

func NewRoomCheckOutResponses() RoomCheckOutResponses {
	return RoomCheckOutResponses{}
}

func SetRoomCheckOutResponse(commonRoom *commonRoom.CommonRoom, commonRoomUser *commonRoomUser.CommonRoomUser) *RoomCheckOutResponse {
	return &RoomCheckOutResponse{
		CommonRoom:     commonRoom,
		CommonRoomUser: commonRoomUser,
	}
}
