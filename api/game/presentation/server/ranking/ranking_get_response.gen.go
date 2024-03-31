// Package ranking ランキング取得レスポンス
package ranking

func SetRankingGetResponse(commonRankingRooms []*CommonRankingRoom, commonRankingWorlds []*CommonRankingWorld) *RankingGetResponse {
	return &RankingGetResponse{
		CommonRankingRooms:  commonRankingRooms,
		CommonRankingWorlds: commonRankingWorlds,
	}
}
