// Package loginBonus ログインボーナス受け取りレスポンス
package loginBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusSchedule"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/userLoginBonus"
)

type LoginBonusReceiveResponses []*LoginBonusReceiveResponse

type LoginBonusReceiveResponse struct {
	UserLoginBonus           *userLoginBonus.UserLoginBonus
	MasterLoginBonus         *masterLoginBonus.MasterLoginBonus
	MasterLoginBonusEvent    *masterLoginBonusEvent.MasterLoginBonusEvent
	MasterLoginBonusItems    *masterLoginBonusItem.MasterLoginBonusItems
	MasterLoginBonusSchedule *masterLoginBonusSchedule.MasterLoginBonusSchedule
}

func NewLoginBonusReceiveResponse() *LoginBonusReceiveResponse {
	return &LoginBonusReceiveResponse{}
}

func NewLoginBonusReceiveResponses() LoginBonusReceiveResponses {
	return LoginBonusReceiveResponses{}
}

func SetLoginBonusReceiveResponse(userLoginBonus *userLoginBonus.UserLoginBonus, masterLoginBonus *masterLoginBonus.MasterLoginBonus, masterLoginBonusEvent *masterLoginBonusEvent.MasterLoginBonusEvent, masterLoginBonusItems *masterLoginBonusItem.MasterLoginBonusItems, masterLoginBonusSchedule *masterLoginBonusSchedule.MasterLoginBonusSchedule) *LoginBonusReceiveResponse {
	return &LoginBonusReceiveResponse{
		UserLoginBonus:           userLoginBonus,
		MasterLoginBonus:         masterLoginBonus,
		MasterLoginBonusEvent:    masterLoginBonusEvent,
		MasterLoginBonusItems:    masterLoginBonusItems,
		MasterLoginBonusSchedule: masterLoginBonusSchedule,
	}
}
