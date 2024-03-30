// Package commonRankingWorld ワールドランキング
package commonRankingWorld

import (
	"time"
)

type CommonRankingWorlds []*CommonRankingWorld

type CommonRankingWorld struct {
	MasterRankingId int64
	UserId          string
	Score           int32
	RankedAt        time.Time
}

func NewCommonRankingWorld() *CommonRankingWorld {
	return &CommonRankingWorld{}
}

func NewCommonRankingWorlds() CommonRankingWorlds {
	return CommonRankingWorlds{}
}

func SetCommonRankingWorld(masterRankingId int64, userId string, score int32, rankedAt time.Time) *CommonRankingWorld {
	return &CommonRankingWorld{
		MasterRankingId: masterRankingId,
		UserId:          userId,
		Score:           score,
		RankedAt:        rankedAt,
	}
}
