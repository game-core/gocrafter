// Package idleBonus 放置ボーナスイベント
package idleBonus

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SetMasterIdleBonusEvent(id int64, name string, resetHour int32, intervalHour int32, repeatSetting bool, startAt *timestamppb.Timestamp, endAt *timestamppb.Timestamp) *MasterIdleBonusEvent {
	return &MasterIdleBonusEvent{
		Id:            id,
		Name:          name,
		ResetHour:     resetHour,
		IntervalHour:  intervalHour,
		RepeatSetting: repeatSetting,
		StartAt:       startAt,
		EndAt:         endAt,
	}
}
