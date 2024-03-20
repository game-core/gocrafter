// Package action アクション確認リクエスト
package action

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type ActionCheckRequests []*ActionCheckRequest

type ActionCheckRequest struct {
	UserId         string
	ActionStepType enum.ActionStepType
	AnyId          *int64
}

func NewActionCheckRequest() *ActionCheckRequest {
	return &ActionCheckRequest{}
}

func NewActionCheckRequests() ActionCheckRequests {
	return ActionCheckRequests{}
}

func SetActionCheckRequest(userId string, actionStepType enum.ActionStepType, anyId *int64) *ActionCheckRequest {
	return &ActionCheckRequest{
		UserId:         userId,
		ActionStepType: actionStepType,
		AnyId:          anyId,
	}
}
