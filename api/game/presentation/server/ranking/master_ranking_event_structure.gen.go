// Package ranking ランキングイベント
package ranking

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SetMasterRankingEvent(id int64, name string, resetHour int32, intervalHour int32, repeatSetting bool, startAt *timestamppb.Timestamp, endAt *timestamppb.Timestamp) *MasterRankingEvent {
	return &MasterRankingEvent{
		Id:            id,
		Name:          name,
		ResetHour:     resetHour,
		IntervalHour:  intervalHour,
		RepeatSetting: repeatSetting,
		StartAt:       startAt,
		EndAt:         endAt,
	}
}
