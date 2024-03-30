// Package ranking ランキング更新レスポンス
package ranking

import (
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingWorld"
)

type RankingUpdateResponses []*RankingUpdateResponse

type RankingUpdateResponse struct {
	CommonRankingRooms  commonRankingRoom.CommonRankingRooms
	CommonRankingWorlds commonRankingWorld.CommonRankingWorlds
}

func NewRankingUpdateResponse() *RankingUpdateResponse {
	return &RankingUpdateResponse{}
}

func NewRankingUpdateResponses() RankingUpdateResponses {
	return RankingUpdateResponses{}
}

func SetRankingUpdateResponse(commonRankingRooms commonRankingRoom.CommonRankingRooms, commonRankingWorlds commonRankingWorld.CommonRankingWorlds) *RankingUpdateResponse {
	return &RankingUpdateResponse{
		CommonRankingRooms:  commonRankingRooms,
		CommonRankingWorlds: commonRankingWorlds,
	}
}
