// Package ranking ランキングマスター取得リクエスト
package ranking

type RankingGetMasterRequests []*RankingGetMasterRequest

type RankingGetMasterRequest struct {
	MasterRankingEventId int64
}

func NewRankingGetMasterRequest() *RankingGetMasterRequest {
	return &RankingGetMasterRequest{}
}

func NewRankingGetMasterRequests() RankingGetMasterRequests {
	return RankingGetMasterRequests{}
}

func SetRankingGetMasterRequest(masterRankingEventId int64) *RankingGetMasterRequest {
	return &RankingGetMasterRequest{
		MasterRankingEventId: masterRankingEventId,
	}
}
