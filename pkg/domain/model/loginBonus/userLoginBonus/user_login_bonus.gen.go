// Package userLoginBonus ユーザーログインボーナス
package userLoginBonus

import (
	"time"
)

type UserLoginBonuses []*UserLoginBonus

type UserLoginBonus struct {
	UserId             string
	MasterLoginBonusId int64
	ReceivedAt         time.Time
}

func NewUserLoginBonus() *UserLoginBonus {
	return &UserLoginBonus{}
}

func NewUserLoginBonuses() UserLoginBonuses {
	return UserLoginBonuses{}
}

func SetUserLoginBonus(userId string, masterLoginBonusId int64, receivedAt time.Time) *UserLoginBonus {
	return &UserLoginBonus{
		UserId:             userId,
		MasterLoginBonusId: masterLoginBonusId,
		ReceivedAt:         receivedAt,
	}
}
