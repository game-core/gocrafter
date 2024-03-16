package idleBonus

import (
	"github.com/game-core/gocrafter/internal/times"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/userIdleBonus"
)

func SetUserIdleBonuses(userIdleBonusModels userIdleBonus.UserIdleBonuses) []*UserIdleBonus {
	var results []*UserIdleBonus

	for _, model := range userIdleBonusModels {
		results = append(
			results,
			SetUserIdleBonus(
				model.UserId,
				model.MasterIdleBonusId,
				times.TimeToPb(&model.ReceivedAt),
			),
		)
	}

	return results
}
