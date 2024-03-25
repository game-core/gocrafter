// Package room ルーム作成リクエスト
package room

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type RoomCreateRequests []*RoomCreateRequest

type RoomCreateRequest struct {
	UserId          string
	Name            string
	RoomReleaseType enum.RoomReleaseType
}

func NewRoomCreateRequest() *RoomCreateRequest {
	return &RoomCreateRequest{}
}

func NewRoomCreateRequests() RoomCreateRequests {
	return RoomCreateRequests{}
}

func SetRoomCreateRequest(userId string, name string, roomReleaseType enum.RoomReleaseType) *RoomCreateRequest {
	return &RoomCreateRequest{
		UserId:          userId,
		Name:            name,
		RoomReleaseType: roomReleaseType,
	}
}
