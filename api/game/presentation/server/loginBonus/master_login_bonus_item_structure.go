package loginBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusItem"
)

func SetMasterLoginBonusItems(masterLoginBonusItemModels masterLoginBonusItem.MasterLoginBonusItems) []*MasterLoginBonusItem {
	var results []*MasterLoginBonusItem

	for _, masterLoginBonusItemModel := range masterLoginBonusItemModels {
		results = append(
			results,
			SetMasterLoginBonusItem(
				masterLoginBonusItemModel.Id,
				masterLoginBonusItemModel.MasterLoginBonusScheduleId,
				masterLoginBonusItemModel.MasterItemId,
				masterLoginBonusItemModel.Name,
				masterLoginBonusItemModel.Count,
			),
		)
	}

	return results
}
