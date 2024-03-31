// Package ranking ランキングマスター取得リクエスト
package ranking

type RankingGetMasterRequests []*RankingGetMasterRequest

type RankingGetMasterRequest struct {
	MasterRankingId int64
}

func NewRankingGetMasterRequest() *RankingGetMasterRequest {
	return &RankingGetMasterRequest{}
}

func NewRankingGetMasterRequests() RankingGetMasterRequests {
	return RankingGetMasterRequests{}
}

func SetRankingGetMasterRequest(masterRankingId int64) *RankingGetMasterRequest {
	return &RankingGetMasterRequest{
		MasterRankingId: masterRankingId,
	}
}
