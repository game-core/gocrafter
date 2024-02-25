// Package masterLoginBonusEvent ログインボーナスイベント
package masterLoginBonusEvent

import (
	"time"
)

type MasterLoginBonusEvents []*MasterLoginBonusEvent

type MasterLoginBonusEvent struct {
	Id            int64
	Name          string
	ResetHour     int32
	IntervalHour  int32
	RepeatSetting bool
	StartAt       time.Time
	EndAt         *time.Time
}

func NewMasterLoginBonusEvent() *MasterLoginBonusEvent {
	return &MasterLoginBonusEvent{}
}

func NewMasterLoginBonusEvents() MasterLoginBonusEvents {
	return MasterLoginBonusEvents{}
}

func SetMasterLoginBonusEvent(id int64, name string, resetHour int32, intervalHour int32, repeatSetting bool, startAt time.Time, endAt *time.Time) *MasterLoginBonusEvent {
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

func (t *MasterLoginBonusEvent) TableName() string {
	return "master_login_bonus_event"
}
