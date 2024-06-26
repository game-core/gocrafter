// Package masterRanking ランキング
package masterRanking

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterRankings []*MasterRanking

type MasterRanking struct {
	Id                   int64
	MasterRankingEventId int64
	Name                 string
	RankingScopeType     enum.RankingScopeType
	RankingLimit         int32
}

func NewMasterRanking() *MasterRanking {
	return &MasterRanking{}
}

func NewMasterRankings() MasterRankings {
	return MasterRankings{}
}

func SetMasterRanking(id int64, masterRankingEventId int64, name string, rankingScopeType enum.RankingScopeType, rankingLimit int32) *MasterRanking {
	return &MasterRanking{
		Id:                   id,
		MasterRankingEventId: masterRankingEventId,
		Name:                 name,
		RankingScopeType:     rankingScopeType,
		RankingLimit:         rankingLimit,
	}
}
