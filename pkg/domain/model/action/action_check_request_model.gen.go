// Package action アクション確認リクエスト
package action

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type ActionCheckRequests []*ActionCheckRequest

type ActionCheckRequest struct {
	UserId         string
	ActionStepType enum.ActionStepType
}

func NewActionCheckRequest() *ActionCheckRequest {
	return &ActionCheckRequest{}
}

func NewActionCheckRequests() ActionCheckRequests {
	return ActionCheckRequests{}
}

func SetActionCheckRequest(userId string, actionStepType enum.ActionStepType) *ActionCheckRequest {
	return &ActionCheckRequest{
		UserId:         userId,
		ActionStepType: actionStepType,
	}
}
