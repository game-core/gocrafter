// Package ranking ランキング更新レスポンス
package ranking

func SetRankingUpdateResponse(commonRankingRooms []*CommonRankingRoom, commonRankingWorlds []*CommonRankingWorld) *RankingUpdateResponse {
	return &RankingUpdateResponse{
		CommonRankingRooms:  commonRankingRooms,
		CommonRankingWorlds: commonRankingWorlds,
	}
}
