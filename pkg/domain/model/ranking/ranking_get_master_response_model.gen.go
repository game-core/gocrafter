// Package ranking ランキングマスター取得レスポンス
package ranking

import (
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRanking"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingScope"
)

type RankingGetMasterResponses []*RankingGetMasterResponse

type RankingGetMasterResponse struct {
	MasterRanking      *masterRanking.MasterRanking
	MasterRankingEvent *masterRankingEvent.MasterRankingEvent
	MasterRankingScope *masterRankingScope.MasterRankingScope
}

func NewRankingGetMasterResponse() *RankingGetMasterResponse {
	return &RankingGetMasterResponse{}
}

func NewRankingGetMasterResponses() RankingGetMasterResponses {
	return RankingGetMasterResponses{}
}

func SetRankingGetMasterResponse(masterRanking *masterRanking.MasterRanking, masterRankingEvent *masterRankingEvent.MasterRankingEvent, masterRankingScope *masterRankingScope.MasterRankingScope) *RankingGetMasterResponse {
	return &RankingGetMasterResponse{
		MasterRanking:      masterRanking,
		MasterRankingEvent: masterRankingEvent,
		MasterRankingScope: masterRankingScope,
	}
}
