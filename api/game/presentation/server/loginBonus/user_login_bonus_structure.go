package loginBonus

import (
	"github.com/game-core/gocrafter/internal/times"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/userLoginBonus"
)

func SetUserLoginBonuses(userLoginBonusModels userLoginBonus.UserLoginBonuses) []*UserLoginBonus {
	var results []*UserLoginBonus

	for _, model := range userLoginBonusModels {
		results = append(
			results,
			SetUserLoginBonus(
				model.UserId,
				model.MasterLoginBonusId,
				times.TimeToPb(&model.ReceivedAt),
			),
		)
	}

	return results
}
