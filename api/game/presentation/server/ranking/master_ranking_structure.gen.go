// Package ranking ランキング
package ranking

func SetMasterRanking(id int64, masterRankingEventId int64, name string, rankingScopeType RankingScopeType, limit int32) *MasterRanking {
	return &MasterRanking{
		Id:                   id,
		MasterRankingEventId: masterRankingEventId,
		Name:                 name,
		RankingScopeType:     rankingScopeType,
		Limit:                limit,
	}
}
