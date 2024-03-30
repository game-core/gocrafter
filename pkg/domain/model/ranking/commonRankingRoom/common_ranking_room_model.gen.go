// Package commonRankingRoom ルームランキング
package commonRankingRoom

type CommonRankingRooms []*CommonRankingRoom

type CommonRankingRoom struct {
	MasterRankingId int64
	RoomId          string
	UserId          string
	Score           int32
}

func NewCommonRankingRoom() *CommonRankingRoom {
	return &CommonRankingRoom{}
}

func NewCommonRankingRooms() CommonRankingRooms {
	return CommonRankingRooms{}
}

func SetCommonRankingRoom(masterRankingId int64, roomId string, userId string, score int32) *CommonRankingRoom {
	return &CommonRankingRoom{
		MasterRankingId: masterRankingId,
		RoomId:          roomId,
		UserId:          userId,
		Score:           score,
	}
}
