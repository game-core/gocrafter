// Package action アクションマスター取得レスポンス
package action

import (
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
)

type ActionGetMasterResponses []*ActionGetMasterResponse

type ActionGetMasterResponse struct {
	MasterActions     masterAction.MasterActions
	MasterActionSteps masterAction.MasterActionSteps
}

func NewActionGetMasterResponse() *ActionGetMasterResponse {
	return &ActionGetMasterResponse{}
}

func NewActionGetMasterResponses() ActionGetMasterResponses {
	return ActionGetMasterResponses{}
}

func SetActionGetMasterResponse(masterActions masterAction.MasterActions, masterActionSteps masterAction.MasterActionSteps) *ActionGetMasterResponse {
	return &ActionGetMasterResponse{
		MasterActions:     masterActions,
		MasterActionSteps: masterActionSteps,
	}
}
