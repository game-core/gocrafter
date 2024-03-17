// Package idleBonus 放置ボーナスマスター取得レスポンス
package idleBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusSchedule"
)

type IdleBonusGetMasterResponses []*IdleBonusGetMasterResponse

type IdleBonusGetMasterResponse struct {
	MasterIdleBonus          *masterIdleBonus.MasterIdleBonus
	MasterIdleBonusEvent     *masterIdleBonusEvent.MasterIdleBonusEvent
	MasterIdleBonusItems     masterIdleBonusItem.MasterIdleBonusItems
	MasterIdleBonusSchedules masterIdleBonusSchedule.MasterIdleBonusSchedules
}

func NewIdleBonusGetMasterResponse() *IdleBonusGetMasterResponse {
	return &IdleBonusGetMasterResponse{}
}

func NewIdleBonusGetMasterResponses() IdleBonusGetMasterResponses {
	return IdleBonusGetMasterResponses{}
}

func SetIdleBonusGetMasterResponse(masterIdleBonus *masterIdleBonus.MasterIdleBonus, masterIdleBonusEvent *masterIdleBonusEvent.MasterIdleBonusEvent, masterIdleBonusItems masterIdleBonusItem.MasterIdleBonusItems, masterIdleBonusSchedules masterIdleBonusSchedule.MasterIdleBonusSchedules) *IdleBonusGetMasterResponse {
	return &IdleBonusGetMasterResponse{
		MasterIdleBonus:          masterIdleBonus,
		MasterIdleBonusEvent:     masterIdleBonusEvent,
		MasterIdleBonusItems:     masterIdleBonusItems,
		MasterIdleBonusSchedules: masterIdleBonusSchedules,
	}
}
