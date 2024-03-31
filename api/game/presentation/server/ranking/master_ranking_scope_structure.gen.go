// Package ranking ランキング範囲
package ranking

func SetMasterRankingScope(id int64, name string, rankingScopeType RankingScopeType) *MasterRankingScope {
	return &MasterRankingScope{
		Id:               id,
		Name:             name,
		RankingScopeType: rankingScopeType,
	}
}
