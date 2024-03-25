// Package room ルーム参加レスポンス
package room

import (
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
)

type RoomJoinResponses []*RoomJoinResponse

type RoomJoinResponse struct {
	CommonRoom     *commonRoom.CommonRoom
	CommonRoomUser *commonRoomUser.CommonRoomUser
}

func NewRoomJoinResponse() *RoomJoinResponse {
	return &RoomJoinResponse{}
}

func NewRoomJoinResponses() RoomJoinResponses {
	return RoomJoinResponses{}
}

func SetRoomJoinResponse(commonRoom *commonRoom.CommonRoom, commonRoomUser *commonRoomUser.CommonRoomUser) *RoomJoinResponse {
	return &RoomJoinResponse{
		CommonRoom:     commonRoom,
		CommonRoomUser: commonRoomUser,
	}
}
