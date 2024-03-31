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

// AggregateRanking 集計する（戻り値: 更新する値, 削除する値）
func (s CommonRankingWorlds) AggregateRanking(masterRankingId int64, userId string, score int32, now, lastEventAt time.Time, rankingLimit int32) (CommonRankingWorlds, *CommonRankingWorld, *CommonRankingWorld) {
	sortRanking := s.SortRanking(lastEventAt)

	commonRankingWorld := sortRanking.updateScoreByUserId(userId, score, now)
	if commonRankingWorld != nil {
		return sortRanking.aggregateByUserId(commonRankingWorld), commonRankingWorld, nil
	}

	if len(sortRanking) >= int(rankingLimit) {
		if sortRanking[len(sortRanking)-1].Score < score {
			updateModel := SetCommonRankingWorld(masterRankingId, userId, score, now)
			deleteModel := sortRanking[len(sortRanking)-1]
			sortRanking[len(sortRanking)-1] = updateModel
			return sortRanking, updateModel, deleteModel
		}
	} else {
		updateModel := SetCommonRankingWorld(masterRankingId, userId, score, now)
		ranking := append(sortRanking, updateModel)
		return ranking, updateModel, nil
	}

	return sortRanking, nil, nil
}

// CheckRankingByUserId ランキングに対象のユーザーが存在するか確認する
func (s CommonRankingWorlds) CheckRankingByUserId(userId string) bool {
	for _, commonRankingWorld := range s {
		if commonRankingWorld.UserId == userId {
			return true
		}
	}

	return false
}

// updateScoreByUserId ユーザーIDからスコアを更新する
func (s CommonRankingWorlds) updateScoreByUserId(userId string, score int32, now time.Time) *CommonRankingWorld {
	for _, commonRankingWorld := range s {
		if commonRankingWorld.UserId == userId && commonRankingWorld.Score <= score {
			commonRankingWorld.Score = score
			commonRankingWorld.RankedAt = now
			return commonRankingWorld
		}
	}

	return nil
}

// aggregateByUserId ユーザーIDから集計する
func (s CommonRankingWorlds) aggregateByUserId(commonRankingWorld *CommonRankingWorld) CommonRankingWorlds {
	for i, model := range s {
		if model.UserId == commonRankingWorld.UserId && model.Score <= commonRankingWorld.Score {
			result := s
			result[i] = commonRankingWorld
			return result
		}
	}

	return s
}
