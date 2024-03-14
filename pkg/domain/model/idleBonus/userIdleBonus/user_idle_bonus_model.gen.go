// Package userIdleBonus ユーザー放置ボーナス
package userIdleBonus

import (
	"time"
)

type UserIdleBonuses []*UserIdleBonus

type UserIdleBonus struct {
	UserId            string
	MasterIdleBonusId int64
	ReceivedAt        time.Time
}

func NewUserIdleBonus() *UserIdleBonus {
	return &UserIdleBonus{}
}

func NewUserIdleBonuses() UserIdleBonuses {
	return UserIdleBonuses{}
}

func SetUserIdleBonus(userId string, masterIdleBonusId int64, receivedAt time.Time) *UserIdleBonus {
	return &UserIdleBonus{
		UserId:            userId,
		MasterIdleBonusId: masterIdleBonusId,
		ReceivedAt:        receivedAt,
	}
}
