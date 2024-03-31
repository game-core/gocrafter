// Package ranking ランキングマスター取得リクエスト
package ranking

func SetRankingGetMasterRequest(userId string, masterRankingId int64) *RankingGetMasterRequest {
	return &RankingGetMasterRequest{
		UserId:          userId,
		MasterRankingId: masterRankingId,
	}
}
