// Package room ルーム確認レスポンス
package room

import (
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
)

type RoomCheckResponses []*RoomCheckResponse

type RoomCheckResponse struct {
	CommonRoom     *commonRoom.CommonRoom
	CommonRoomUser *commonRoomUser.CommonRoomUser
}

func NewRoomCheckResponse() *RoomCheckResponse {
	return &RoomCheckResponse{}
}

func NewRoomCheckResponses() RoomCheckResponses {
	return RoomCheckResponses{}
}

func SetRoomCheckResponse(commonRoom *commonRoom.CommonRoom, commonRoomUser *commonRoomUser.CommonRoomUser) *RoomCheckResponse {
	return &RoomCheckResponse{
		CommonRoom:     commonRoom,
		CommonRoomUser: commonRoomUser,
	}
}
