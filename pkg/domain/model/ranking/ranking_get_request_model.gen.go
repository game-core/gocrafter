// Package ranking ランキング取得リクエスト
package ranking

type RankingGetRequests []*RankingGetRequest

type RankingGetRequest struct {
	UserId          string
	MasterRankingId int64
	RoomId          string
}

func NewRankingGetRequest() *RankingGetRequest {
	return &RankingGetRequest{}
}

func NewRankingGetRequests() RankingGetRequests {
	return RankingGetRequests{}
}

func SetRankingGetRequest(userId string, masterRankingId int64, roomId string) *RankingGetRequest {
	return &RankingGetRequest{
		UserId:          userId,
		MasterRankingId: masterRankingId,
		RoomId:          roomId,
	}
}
