// Package loginBonus ログインボーナスマスター取得レスポンス
package loginBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusSchedule"
)

type LoginBonusGetMasterResponses []*LoginBonusGetMasterResponse

type LoginBonusGetMasterResponse struct {
	MasterLoginBonus          *masterLoginBonus.MasterLoginBonus
	MasterLoginBonusEvent     *masterLoginBonusEvent.MasterLoginBonusEvent
	MasterLoginBonusItems     masterLoginBonusItem.MasterLoginBonusItems
	MasterLoginBonusSchedules masterLoginBonusSchedule.MasterLoginBonusSchedules
}

func NewLoginBonusGetMasterResponse() *LoginBonusGetMasterResponse {
	return &LoginBonusGetMasterResponse{}
}

func NewLoginBonusGetMasterResponses() LoginBonusGetMasterResponses {
	return LoginBonusGetMasterResponses{}
}

func SetLoginBonusGetMasterResponse(masterLoginBonus *masterLoginBonus.MasterLoginBonus, masterLoginBonusEvent *masterLoginBonusEvent.MasterLoginBonusEvent, masterLoginBonusItems masterLoginBonusItem.MasterLoginBonusItems, masterLoginBonusSchedules masterLoginBonusSchedule.MasterLoginBonusSchedules) *LoginBonusGetMasterResponse {
	return &LoginBonusGetMasterResponse{
		MasterLoginBonus:          masterLoginBonus,
		MasterLoginBonusEvent:     masterLoginBonusEvent,
		MasterLoginBonusItems:     masterLoginBonusItems,
		MasterLoginBonusSchedules: masterLoginBonusSchedules,
	}
}
