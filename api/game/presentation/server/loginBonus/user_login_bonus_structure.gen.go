// Package loginBonus ユーザーログインボーナス
package loginBonus

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SetUserLoginBonus(userId string, masterLoginBonusId int64, receivedAt *timestamppb.Timestamp) *UserLoginBonus {
	return &UserLoginBonus{
		UserId:             userId,
		MasterLoginBonusId: masterLoginBonusId,
		ReceivedAt:         receivedAt,
	}
}
