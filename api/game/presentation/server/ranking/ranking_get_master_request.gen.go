// Package ranking ランキングマスター取得リクエスト
package ranking

func SetRankingGetMasterRequest(userId string, masterRankingEventId int64) *RankingGetMasterRequest {
	return &RankingGetMasterRequest{
		UserId:               userId,
		MasterRankingEventId: masterRankingEventId,
	}
}
