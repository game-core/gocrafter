// Package ranking ルームランキング
package ranking

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SetCommonRankingRoom(masterRankingId int64, roomId string, userId string, score int32, rankedAt *timestamppb.Timestamp) *CommonRankingRoom {
	return &CommonRankingRoom{
		MasterRankingId: masterRankingId,
		RoomId:          roomId,
		UserId:          userId,
		Score:           score,
		RankedAt:        rankedAt,
	}
}
