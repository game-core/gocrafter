package loginBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusSchedule"
)

func SetMasterLoginBonusSchedules(masterLoginBonusScheduleModels masterLoginBonusSchedule.MasterLoginBonusSchedules) []*MasterLoginBonusSchedule {
	var results []*MasterLoginBonusSchedule

	for _, masterLoginBonusScheduleModel := range masterLoginBonusScheduleModels {
		results = append(
			results,
			SetMasterLoginBonusSchedule(
				masterLoginBonusScheduleModel.Id,
				masterLoginBonusScheduleModel.MasterLoginBonusId,
				masterLoginBonusScheduleModel.Step,
				masterLoginBonusScheduleModel.Name,
			),
		)
	}

	return results
}
