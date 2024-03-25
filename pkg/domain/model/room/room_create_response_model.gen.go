// Package room ルーム作成レスポンス
package room

import (
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
)

type RoomCreateResponses []*RoomCreateResponse

type RoomCreateResponse struct {
	CommonRoom *commonRoom.CommonRoom
}

func NewRoomCreateResponse() *RoomCreateResponse {
	return &RoomCreateResponse{}
}

func NewRoomCreateResponses() RoomCreateResponses {
	return RoomCreateResponses{}
}

func SetRoomCreateResponse(commonRoom *commonRoom.CommonRoom) *RoomCreateResponse {
	return &RoomCreateResponse{
		CommonRoom: commonRoom,
	}
}
