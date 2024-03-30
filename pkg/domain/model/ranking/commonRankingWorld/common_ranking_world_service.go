package commonRankingWorld

import (
	"sort"
	"time"
)

// SortRanking ランキングをソートする
func (s CommonRankingWorlds) SortRanking(lastEventAt time.Time) CommonRankingWorlds {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Score > s[j].Score
	})

	filtered := make(CommonRankingWorlds, 0, len(s))
	for _, entry := range s {
		if entry.RankedAt.After(lastEventAt) {
			filtered = append(filtered, entry)
		}
	}

	return filtered
}

// ExcludeRanking ランキング除外
func (s CommonRankingWorlds) ExcludeRanking(masterRankingId int64, userId string, score int32, now, lastEventAt time.Time) *CommonRankingWorld {
	sortRanking := append(s, SetCommonRankingWorld(masterRankingId, userId, score, now)).SortRanking(lastEventAt)

	return sortRanking[len(sortRanking)-1]
}
