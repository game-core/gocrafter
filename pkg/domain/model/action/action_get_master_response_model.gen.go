// Package action アクションマスター取得レスポンス
package action

import (
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionStep"
)

type ActionGetMasterResponses []*ActionGetMasterResponse

type ActionGetMasterResponse struct {
	MasterActions     masterAction.MasterActions
	MasterActionSteps masterActionStep.MasterActionSteps
}

func NewActionGetMasterResponse() *ActionGetMasterResponse {
	return &ActionGetMasterResponse{}
}

func NewActionGetMasterResponses() ActionGetMasterResponses {
	return ActionGetMasterResponses{}
}

func SetActionGetMasterResponse(masterActions masterAction.MasterActions, masterActionSteps masterActionStep.MasterActionSteps) *ActionGetMasterResponse {
	return &ActionGetMasterResponse{
		MasterActions:     masterActions,
		MasterActionSteps: masterActionSteps,
	}
}
