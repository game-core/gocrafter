// Package masterRankingScope ランキング範囲
package masterRankingScope

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterRankingScopes []*MasterRankingScope

type MasterRankingScope struct {
	Id          int64
	Name        string
	RankingType enum.RankingType
}

func NewMasterRankingScope() *MasterRankingScope {
	return &MasterRankingScope{}
}

func NewMasterRankingScopes() MasterRankingScopes {
	return MasterRankingScopes{}
}

func SetMasterRankingScope(id int64, name string, rankingType enum.RankingType) *MasterRankingScope {
	return &MasterRankingScope{
		Id:          id,
		Name:        name,
		RankingType: rankingType,
	}
}

func (t *MasterRankingScope) TableName() string {
	return "master_ranking_scope"
}
