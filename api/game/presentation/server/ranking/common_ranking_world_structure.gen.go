// Package ranking ワールドランキング
package ranking

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SetCommonRankingWorld(masterRankingId int64, userId string, score int32, rankedAt *timestamppb.Timestamp) *CommonRankingWorld {
	return &CommonRankingWorld{
		MasterRankingId: masterRankingId,
		UserId:          userId,
		Score:           score,
		RankedAt:        rankedAt,
	}
}
