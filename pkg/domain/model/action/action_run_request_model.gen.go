// Package action アクション実行リクエスト
package action

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type ActionRunRequests []*ActionRunRequest

type ActionRunRequest struct {
	UserId         string
	ActionStepType enum.ActionStepType
}

func NewActionRunRequest() *ActionRunRequest {
	return &ActionRunRequest{}
}

func NewActionRunRequests() ActionRunRequests {
	return ActionRunRequests{}
}

func SetActionRunRequest(userId string, actionStepType enum.ActionStepType) *ActionRunRequest {
	return &ActionRunRequest{
		UserId:         userId,
		ActionStepType: actionStepType,
	}
}
