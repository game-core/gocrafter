// Package action アクションマスター取得レスポンス
package action

import (
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionRun"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionStep"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionTrigger"
)

type ActionGetMasterResponses []*ActionGetMasterResponse

type ActionGetMasterResponse struct {
	MasterActions        masterAction.MasterActions
	MasterActionRuns     masterActionRun.MasterActionRuns
	MasterActionSteps    masterActionStep.MasterActionSteps
	MasterActionTriggers masterActionTrigger.MasterActionTriggers
}

func NewActionGetMasterResponse() *ActionGetMasterResponse {
	return &ActionGetMasterResponse{}
}

func NewActionGetMasterResponses() ActionGetMasterResponses {
	return ActionGetMasterResponses{}
}

func SetActionGetMasterResponse(masterActions masterAction.MasterActions, masterActionRuns masterActionRun.MasterActionRuns, masterActionSteps masterActionStep.MasterActionSteps, masterActionTriggers masterActionTrigger.MasterActionTriggers) *ActionGetMasterResponse {
	return &ActionGetMasterResponse{
		MasterActions:        masterActions,
		MasterActionRuns:     masterActionRuns,
		MasterActionSteps:    masterActionSteps,
		MasterActionTriggers: masterActionTriggers,
	}
}
