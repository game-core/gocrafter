// Package masterIdleBonusEvent 放置ボーナスイベント
package masterIdleBonusEvent

import (
	"time"
)

type MasterIdleBonusEvents []*MasterIdleBonusEvent

type MasterIdleBonusEvent struct {
	Id            int64
	Name          string
	ResetHour     int32
	IntervalHour  int32
	RepeatSetting bool
	StartAt       time.Time
	EndAt         *time.Time
}

func NewMasterIdleBonusEvent() *MasterIdleBonusEvent {
	return &MasterIdleBonusEvent{}
}

func NewMasterIdleBonusEvents() MasterIdleBonusEvents {
	return MasterIdleBonusEvents{}
}

func SetMasterIdleBonusEvent(id int64, name string, resetHour int32, intervalHour int32, repeatSetting bool, startAt time.Time, endAt *time.Time) *MasterIdleBonusEvent {
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

func (t *MasterIdleBonusEvent) TableName() string {
	return "master_idle_bonus_event"
}
