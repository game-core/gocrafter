// Package ranking ランキング取得レスポンス
package ranking

import (
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingWorld"
)

type RankingGetResponses []*RankingGetResponse

type RankingGetResponse struct {
	CommonRankingRooms  commonRankingRoom.CommonRankingRooms
	CommonRankingWorlds commonRankingWorld.CommonRankingWorlds
}

func NewRankingGetResponse() *RankingGetResponse {
	return &RankingGetResponse{}
}

func NewRankingGetResponses() RankingGetResponses {
	return RankingGetResponses{}
}

func SetRankingGetResponse(commonRankingRooms commonRankingRoom.CommonRankingRooms, commonRankingWorlds commonRankingWorld.CommonRankingWorlds) *RankingGetResponse {
	return &RankingGetResponse{
		CommonRankingRooms:  commonRankingRooms,
		CommonRankingWorlds: commonRankingWorlds,
	}
}
