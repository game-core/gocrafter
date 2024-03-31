// Package ranking ランキング取得リクエスト
package ranking

func SetRankingGetRequest(userId string, masterRankingId int64, roomId string) *RankingGetRequest {
	return &RankingGetRequest{
		UserId:          userId,
		MasterRankingId: masterRankingId,
		RoomId:          roomId,
	}
}
