// Package masterActionTrigger アクショントリガー
package masterActionTrigger

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterActionTriggers []*MasterActionTrigger

type MasterActionTrigger struct {
	Id                int64
	Name              string
	ActionTriggerType enum.ActionTriggerType
}

func NewMasterActionTrigger() *MasterActionTrigger {
	return &MasterActionTrigger{}
}

func NewMasterActionTriggers() MasterActionTriggers {
	return MasterActionTriggers{}
}

func SetMasterActionTrigger(id int64, name string, actionTriggerType enum.ActionTriggerType) *MasterActionTrigger {
	return &MasterActionTrigger{
		Id:                id,
		Name:              name,
		ActionTriggerType: actionTriggerType,
	}
}
