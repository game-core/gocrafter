// Package idleBonus 放置ボーナス受け取りレスポンス
package idleBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusSchedule"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/userIdleBonus"
)

type IdleBonusReceiveResponses []*IdleBonusReceiveResponse

type IdleBonusReceiveResponse struct {
	UserIdleBonus            *userIdleBonus.UserIdleBonus
	MasterIdleBonus          *masterIdleBonus.MasterIdleBonus
	MasterIdleBonusEvent     *masterIdleBonusEvent.MasterIdleBonusEvent
	MasterIdleBonusItems     masterIdleBonusItem.MasterIdleBonusItems
	MasterIdleBonusSchedules masterIdleBonusSchedule.MasterIdleBonusSchedules
}

func NewIdleBonusReceiveResponse() *IdleBonusReceiveResponse {
	return &IdleBonusReceiveResponse{}
}

func NewIdleBonusReceiveResponses() IdleBonusReceiveResponses {
	return IdleBonusReceiveResponses{}
}

func SetIdleBonusReceiveResponse(userIdleBonus *userIdleBonus.UserIdleBonus, masterIdleBonus *masterIdleBonus.MasterIdleBonus, masterIdleBonusEvent *masterIdleBonusEvent.MasterIdleBonusEvent, masterIdleBonusItems masterIdleBonusItem.MasterIdleBonusItems, masterIdleBonusSchedules masterIdleBonusSchedule.MasterIdleBonusSchedules) *IdleBonusReceiveResponse {
	return &IdleBonusReceiveResponse{
		UserIdleBonus:            userIdleBonus,
		MasterIdleBonus:          masterIdleBonus,
		MasterIdleBonusEvent:     masterIdleBonusEvent,
		MasterIdleBonusItems:     masterIdleBonusItems,
		MasterIdleBonusSchedules: masterIdleBonusSchedules,
	}
}
