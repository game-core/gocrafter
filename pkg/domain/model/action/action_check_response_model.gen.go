// Package action アクション確認レスポンス
package action

import (
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionStep"
)

type ActionCheckResponses []*ActionCheckResponse

type ActionCheckResponse struct {
	Executable       bool
	MasterAction     *masterAction.MasterAction
	MasterActionStep *masterActionStep.MasterActionStep
}

func NewActionCheckResponse() *ActionCheckResponse {
	return &ActionCheckResponse{}
}

func NewActionCheckResponses() ActionCheckResponses {
	return ActionCheckResponses{}
}

func SetActionCheckResponse(executable bool, masterAction *masterAction.MasterAction, masterActionStep *masterActionStep.MasterActionStep) *ActionCheckResponse {
	return &ActionCheckResponse{
		Executable:       executable,
		MasterAction:     masterAction,
		MasterActionStep: masterActionStep,
	}
}
