// Package action アクション実行レスポンス
package action

import (
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionStep"
	"github.com/game-core/gocrafter/pkg/domain/model/action/userAction"
)

type ActionRunResponses []*ActionRunResponse

type ActionRunResponse struct {
	UserAction       *userAction.UserAction
	MasterAction     *masterAction.MasterAction
	MasterActionStep *masterActionStep.MasterActionStep
}

func NewActionRunResponse() *ActionRunResponse {
	return &ActionRunResponse{}
}

func NewActionRunResponses() ActionRunResponses {
	return ActionRunResponses{}
}

func SetActionRunResponse(userAction *userAction.UserAction, masterAction *masterAction.MasterAction, masterActionStep *masterActionStep.MasterActionStep) *ActionRunResponse {
	return &ActionRunResponse{
		UserAction:       userAction,
		MasterAction:     masterAction,
		MasterActionStep: masterActionStep,
	}
}
