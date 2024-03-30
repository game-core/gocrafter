// Package commonRankingRoom ルームランキング
package commonRankingRoom

import (
	"time"
)

type CommonRankingRooms []*CommonRankingRoom

type CommonRankingRoom struct {
	MasterRankingId int64
	RoomId          string
	UserId          string
	Score           int32
	RankedAt        time.Time
}

func NewCommonRankingRoom() *CommonRankingRoom {
	return &CommonRankingRoom{}
}

func NewCommonRankingRooms() CommonRankingRooms {
	return CommonRankingRooms{}
}

func SetCommonRankingRoom(masterRankingId int64, roomId string, userId string, score int32, rankedAt time.Time) *CommonRankingRoom {
	return &CommonRankingRoom{
		MasterRankingId: masterRankingId,
		RoomId:          roomId,
		UserId:          userId,
		Score:           score,
		RankedAt:        rankedAt,
	}
}
