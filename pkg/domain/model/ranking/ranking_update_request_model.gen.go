// Package ranking ランキング更新リクエスト
package ranking

type RankingUpdateRequests []*RankingUpdateRequest

type RankingUpdateRequest struct {
	UserId               string
	MasterRankingEventId int64
	RoomId               string
	Score                int32
}

func NewRankingUpdateRequest() *RankingUpdateRequest {
	return &RankingUpdateRequest{}
}

func NewRankingUpdateRequests() RankingUpdateRequests {
	return RankingUpdateRequests{}
}

func SetRankingUpdateRequest(userId string, masterRankingEventId int64, roomId string, score int32) *RankingUpdateRequest {
	return &RankingUpdateRequest{
		UserId:               userId,
		MasterRankingEventId: masterRankingEventId,
		RoomId:               roomId,
		Score:                score,
	}
}
