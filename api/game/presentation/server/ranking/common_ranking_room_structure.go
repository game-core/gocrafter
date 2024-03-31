package ranking

import (
	"github.com/game-core/gocrafter/internal/times"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingRoom"
)

func SetCommonRankingRooms(commonRankingRoomModels commonRankingRoom.CommonRankingRooms) []*CommonRankingRoom {
	var results []*CommonRankingRoom

	for _, model := range commonRankingRoomModels {
		results = append(
			results,
			SetCommonRankingRoom(
				model.MasterRankingId,
				model.RoomId,
				model.UserId,
				model.Score,
				times.TimeToPb(&model.RankedAt),
			),
		)
	}

	return results
}
