// Package room ルーム参加レスポンス
package room

import (
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
)

type RoomCheckInResponses []*RoomCheckInResponse

type RoomCheckInResponse struct {
	CommonRoom     *commonRoom.CommonRoom
	CommonRoomUser *commonRoomUser.CommonRoomUser
}

func NewRoomCheckInResponse() *RoomCheckInResponse {
	return &RoomCheckInResponse{}
}

func NewRoomCheckInResponses() RoomCheckInResponses {
	return RoomCheckInResponses{}
}

func SetRoomCheckInResponse(commonRoom *commonRoom.CommonRoom, commonRoomUser *commonRoomUser.CommonRoomUser) *RoomCheckInResponse {
	return &RoomCheckInResponse{
		CommonRoom:     commonRoom,
		CommonRoomUser: commonRoomUser,
	}
}
