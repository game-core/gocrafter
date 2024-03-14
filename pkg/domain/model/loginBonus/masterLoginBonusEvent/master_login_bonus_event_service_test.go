package masterLoginBonusEvent

import (
	"reflect"
	"testing"
	"time"

	"github.com/game-core/gocrafter/internal/times"
)

func TestMasterLoginBonusEvent_CheckEventPeriod(t *testing.T) {
	type fields struct {
		MasterLoginBonusEvent *MasterLoginBonusEvent
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "正常：イベント期間内の場合",
			fields: fields{
				MasterLoginBonusEvent: &MasterLoginBonusEvent{
					Id:            1,
					Name:          "テストイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         nil,
				},
			},
			args: args{
				now: time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "正常：イベント期間外の場合（開始前）",
			fields: fields{
				MasterLoginBonusEvent: &MasterLoginBonusEvent{
					Id:            1,
					Name:          "テストイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         nil,
				},
			},
			args: args{
				now: time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC),
			},
			want:    false,
			wantErr: nil,
		},
		{
			name: "正常：イベント期間外の場合（終了後）",
			fields: fields{
				MasterLoginBonusEvent: &MasterLoginBonusEvent{
					Id:            1,
					Name:          "テストイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         times.TimeToPointer(time.Date(2023, 1, 10, 8, 59, 59, 59, time.UTC)),
				},
			},
			args: args{
				now: time.Date(2023, 1, 11, 9, 0, 0, 0, time.UTC),
			},
			want:    false,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.MasterLoginBonusEvent.CheckEventPeriod(tt.args.now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckEventPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}
