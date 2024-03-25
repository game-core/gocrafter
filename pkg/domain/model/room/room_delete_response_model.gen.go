// Package room ルーム削除レスポンス
package room

import (
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
)

type RoomDeleteResponses []*RoomDeleteResponse

type RoomDeleteResponse struct {
	CommonRoom *commonRoom.CommonRoom
}

func NewRoomDeleteResponse() *RoomDeleteResponse {
	return &RoomDeleteResponse{}
}

func NewRoomDeleteResponses() RoomDeleteResponses {
	return RoomDeleteResponses{}
}

func SetRoomDeleteResponse(commonRoom *commonRoom.CommonRoom) *RoomDeleteResponse {
	return &RoomDeleteResponse{
		CommonRoom: commonRoom,
	}
}
