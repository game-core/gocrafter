package ranking

import (
	"github.com/game-core/gocrafter/internal/times"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingWorld"
)

func SetCommonRankingWorlds(commonRankingWorldModels commonRankingWorld.CommonRankingWorlds) []*CommonRankingWorld {
	var results []*CommonRankingWorld

	for _, model := range commonRankingWorldModels {
		results = append(
			results,
			SetCommonRankingWorld(
				model.MasterRankingId,
				model.UserId,
				model.Score,
				times.TimeToPb(&model.RankedAt),
			),
		)
	}

	return results
}
