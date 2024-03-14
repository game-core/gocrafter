package idleBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusSchedule"
)

func SetMasterIdleBonusSchedules(masterIdleBonusScheduleModels masterIdleBonusSchedule.MasterIdleBonusSchedules) []*MasterIdleBonusSchedule {
	var results []*MasterIdleBonusSchedule

	for _, masterIdleBonusScheduleModel := range masterIdleBonusScheduleModels {
		results = append(
			results,
			SetMasterIdleBonusSchedule(
				masterIdleBonusScheduleModel.Id,
				masterIdleBonusScheduleModel.MasterIdleBonusId,
				masterIdleBonusScheduleModel.Step,
				masterIdleBonusScheduleModel.Name,
			),
		)
	}

	return results
}
