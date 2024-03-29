// Package masterEvent イベント
package masterEvent

import (
	"time"
)

type MasterEvents []*MasterEvent

type MasterEvent struct {
	Id            int64
	Name          string
	ResetHour     int32
	IntervalHour  int32
	RepeatSetting bool
	StartAt       time.Time
	EndAt         *time.Time
}

func NewMasterEvent() *MasterEvent {
	return &MasterEvent{}
}

func NewMasterEvents() MasterEvents {
	return MasterEvents{}
}

func SetMasterEvent(id int64, name string, resetHour int32, intervalHour int32, repeatSetting bool, startAt time.Time, endAt *time.Time) *MasterEvent {
	return &MasterEvent{
		Id:            id,
		Name:          name,
		ResetHour:     resetHour,
		IntervalHour:  intervalHour,
		RepeatSetting: repeatSetting,
		StartAt:       startAt,
		EndAt:         endAt,
	}
}

func (t *MasterEvent) TableName() string {
	return "master_event"
}
