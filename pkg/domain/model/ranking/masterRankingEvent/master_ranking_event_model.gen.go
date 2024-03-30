// Package masterRankingEvent ランキングイベント
package masterRankingEvent

import (
	"time"
)

type MasterRankingEvents []*MasterRankingEvent

type MasterRankingEvent struct {
	Id            int64
	Name          string
	ResetHour     int32
	IntervalHour  int32
	RepeatSetting bool
	StartAt       time.Time
	EndAt         *time.Time
}

func NewMasterRankingEvent() *MasterRankingEvent {
	return &MasterRankingEvent{}
}

func NewMasterRankingEvents() MasterRankingEvents {
	return MasterRankingEvents{}
}

func SetMasterRankingEvent(id int64, name string, resetHour int32, intervalHour int32, repeatSetting bool, startAt time.Time, endAt *time.Time) *MasterRankingEvent {
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
