// Package loginBonus ログインボーナスイベント
package loginBonus

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SetMasterLoginBonusEvent(id int64, name string, resetHour int32, intervalHour int32, repeatSetting bool, startAt *timestamppb.Timestamp, endAt *timestamppb.Timestamp) *MasterLoginBonusEvent {
	return &MasterLoginBonusEvent{
		Id:            id,
		Name:          name,
		ResetHour:     resetHour,
		IntervalHour:  intervalHour,
		RepeatSetting: repeatSetting,
		StartAt:       startAt,
		EndAt:         endAt,
	}
}
