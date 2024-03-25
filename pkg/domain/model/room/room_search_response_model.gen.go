// Package room ルーム検索リクエスト
package room

import (
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
)

type RoomSearchResponses []*RoomSearchResponse

type RoomSearchResponse struct {
	CommonRooms commonRoom.CommonRooms
}

func NewRoomSearchResponse() *RoomSearchResponse {
	return &RoomSearchResponse{}
}

func NewRoomSearchResponses() RoomSearchResponses {
	return RoomSearchResponses{}
}

func SetRoomSearchResponse(commonRooms commonRoom.CommonRooms) *RoomSearchResponse {
	return &RoomSearchResponse{
		CommonRooms: commonRooms,
	}
}
