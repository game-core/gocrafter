// Package masterAction アクション
package masterAction

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterActions []*MasterAction

type MasterAction struct {
	Id              int64
	Name            string
	ActionStepType  enum.ActionStepType
	AnyId           *int64
	TriggerActionId *int64
	Expiration      *int32
}

func NewMasterAction() *MasterAction {
	return &MasterAction{}
}

func NewMasterActions() MasterActions {
	return MasterActions{}
}

func SetMasterAction(id int64, name string, actionStepType enum.ActionStepType, anyId *int64, triggerActionId *int64, expiration *int32) *MasterAction {
	return &MasterAction{
		Id:              id,
		Name:            name,
		ActionStepType:  actionStepType,
		AnyId:           anyId,
		TriggerActionId: triggerActionId,
		Expiration:      expiration,
	}
}
