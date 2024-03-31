// Package ranking ランキング更新リクエスト
package ranking

func SetRankingUpdateRequest(userId string, masterRankingId int64, roomId string, score int32) *RankingUpdateRequest {
	return &RankingUpdateRequest{
		UserId:          userId,
		MasterRankingId: masterRankingId,
		RoomId:          roomId,
		Score:           score,
	}
}
