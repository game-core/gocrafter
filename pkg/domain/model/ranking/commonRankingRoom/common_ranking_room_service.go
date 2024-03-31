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

// AggregateRanking 集計する（戻り値: 更新する値, 削除する値）
func (s CommonRankingRooms) AggregateRanking(masterRankingId int64, roomId string, userId string, score int32, now, lastEventAt time.Time, rankingLimit int32) (CommonRankingRooms, *CommonRankingRoom, *CommonRankingRoom) {
	sortRanking := s.SortRanking(lastEventAt)

	commonRankingRoom := sortRanking.updateScoreByUserId(userId, score, now)
	if commonRankingRoom != nil {
		return sortRanking.aggregateByUserId(commonRankingRoom), commonRankingRoom, nil
	}

	if len(sortRanking) >= int(rankingLimit) {
		if sortRanking[len(sortRanking)-1].Score < score {
			updateModel := SetCommonRankingRoom(masterRankingId, roomId, userId, score, now)
			deleteModel := sortRanking[len(sortRanking)-1]
			sortRanking[len(sortRanking)-1] = updateModel
			return sortRanking, updateModel, deleteModel
		}
	} else {
		updateModel := SetCommonRankingRoom(masterRankingId, roomId, userId, score, now)
		ranking := append(sortRanking, updateModel)
		return ranking, updateModel, nil
	}

	return sortRanking, nil, nil
}

// CheckRankingByUserId ランキングに対象のユーザーが存在するか確認する
func (s CommonRankingRooms) CheckRankingByUserId(userId string) bool {
	for _, commonRankingRoom := range s {
		if commonRankingRoom.UserId == userId {
			return true
		}
	}

	return false
}

// updateScoreByUserId ユーザーIDからスコアを更新する
func (s CommonRankingRooms) updateScoreByUserId(userId string, score int32, now time.Time) *CommonRankingRoom {
	for _, commonRankingRoom := range s {
		if commonRankingRoom.UserId == userId && commonRankingRoom.Score <= score {
			commonRankingRoom.Score = score
			commonRankingRoom.RankedAt = now
			return commonRankingRoom
		}
	}

	return nil
}

// aggregateByUserId ユーザーIDから集計する
func (s CommonRankingRooms) aggregateByUserId(commonRankingRoom *CommonRankingRoom) CommonRankingRooms {
	for i, model := range s {
		if model.UserId == commonRankingRoom.UserId && model.Score <= commonRankingRoom.Score {
			result := s
			result[i] = commonRankingRoom
			return result
		}
	}

	return s
}
