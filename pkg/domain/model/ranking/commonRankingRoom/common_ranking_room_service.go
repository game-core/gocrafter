package commonRankingRoom

import (
	"sort"
	"time"
)

// SortRanking ランキングをソートする
func (s CommonRankingRooms) SortRanking(lastEventAt time.Time) CommonRankingRooms {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Score > s[j].Score
	})

	filtered := make(CommonRankingRooms, 0, len(s))
	for _, entry := range s {
		if entry.RankedAt.After(lastEventAt) {
			filtered = append(filtered, entry)
		}
	}

	return filtered
}

// ExcludeRanking ランキング除外
func (s CommonRankingRooms) ExcludeRanking(masterRankingId int64, roomId string, userId string, score int32, now, lastEventAt time.Time) *CommonRankingRoom {
	sortRanking := append(s, SetCommonRankingRoom(masterRankingId, roomId, userId, score, now)).SortRanking(lastEventAt)

	return sortRanking[len(sortRanking)-1]
}
