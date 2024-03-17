package idleBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusItem"
)

func SetMasterIdleBonusItems(masterIdleBonusItemModels masterIdleBonusItem.MasterIdleBonusItems) []*MasterIdleBonusItem {
	var results []*MasterIdleBonusItem

	for _, masterIdleBonusItemModel := range masterIdleBonusItemModels {
		results = append(
			results,
			SetMasterIdleBonusItem(
				masterIdleBonusItemModel.Id,
				masterIdleBonusItemModel.MasterIdleBonusScheduleId,
				masterIdleBonusItemModel.MasterItemId,
				masterIdleBonusItemModel.Name,
				masterIdleBonusItemModel.Count,
			),
		)
	}

	return results
}
