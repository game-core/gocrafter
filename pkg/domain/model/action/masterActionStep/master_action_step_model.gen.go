// Package masterActionStep アクションステップ
package masterActionStep

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterActionSteps []*MasterActionStep

type MasterActionStep struct {
	Id             int64
	Name           string
	ActionStepType enum.ActionStepType
}

func NewMasterActionStep() *MasterActionStep {
	return &MasterActionStep{}
}

func NewMasterActionSteps() MasterActionSteps {
	return MasterActionSteps{}
}

func SetMasterActionStep(id int64, name string, actionStepType enum.ActionStepType) *MasterActionStep {
	return &MasterActionStep{
		Id:             id,
		Name:           name,
		ActionStepType: actionStepType,
	}
}
