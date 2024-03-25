package room

import (
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
)

func SetCommonRooms(commonRooms commonRoom.CommonRooms) []*CommonRoom {
	var results []*CommonRoom

	for _, commonRoom := range commonRooms {
		results = append(
			results,
			SetCommonRoom(
				commonRoom.RoomId,
				commonRoom.HostUserId,
				RoomReleaseType(commonRoom.RoomReleaseType),
				commonRoom.Name,
				commonRoom.UserCount,
			),
		)
	}

	return results
}
