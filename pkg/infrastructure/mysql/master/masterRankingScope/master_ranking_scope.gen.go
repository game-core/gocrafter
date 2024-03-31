// Package masterRankingScope ランキング範囲
package masterRankingScope

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterRankingScopes []*MasterRankingScope

type MasterRankingScope struct {
	Id               int64
	Name             string
	RankingScopeType enum.RankingScopeType
}

func NewMasterRankingScope() *MasterRankingScope {
	return &MasterRankingScope{}
}

func NewMasterRankingScopes() MasterRankingScopes {
	return MasterRankingScopes{}
}

func SetMasterRankingScope(id int64, name string, rankingScopeType enum.RankingScopeType) *MasterRankingScope {
	return &MasterRankingScope{
		Id:               id,
		Name:             name,
		RankingScopeType: rankingScopeType,
	}
}

func (t *MasterRankingScope) TableName() string {
	return "master_ranking_scope"
}
