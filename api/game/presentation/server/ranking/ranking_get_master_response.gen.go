// Package ranking ランキングマスター取得レスポンス
package ranking

func SetRankingGetMasterResponse(masterRanking *MasterRanking, masterRankingEvent *MasterRankingEvent, masterRankingScope *MasterRankingScope) *RankingGetMasterResponse {
	return &RankingGetMasterResponse{
		MasterRanking:      masterRanking,
		MasterRankingEvent: masterRankingEvent,
		MasterRankingScope: masterRankingScope,
	}
}
