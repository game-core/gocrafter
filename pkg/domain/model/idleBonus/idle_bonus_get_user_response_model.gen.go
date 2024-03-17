// Package idleBonus 放置ボーナスユーザー取得レスポンス
package idleBonus

import (
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/userIdleBonus"
)

type IdleBonusGetUserResponses []*IdleBonusGetUserResponse

type IdleBonusGetUserResponse struct {
	UserIdleBonuses userIdleBonus.UserIdleBonuses
}

func NewIdleBonusGetUserResponse() *IdleBonusGetUserResponse {
	return &IdleBonusGetUserResponse{}
}

func NewIdleBonusGetUserResponses() IdleBonusGetUserResponses {
	return IdleBonusGetUserResponses{}
}

func SetIdleBonusGetUserResponse(userIdleBonuses userIdleBonus.UserIdleBonuses) *IdleBonusGetUserResponse {
	return &IdleBonusGetUserResponse{
		UserIdleBonuses: userIdleBonuses,
	}
}
