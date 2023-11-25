package event

import (
	"github.com/game-core/gocrafter/config/pointer"
	"reflect"
	"testing"
	"time"
)

func TestExampleService_GetDayCount(t *testing.T) {
	type fields struct {
		RepeatSetting bool
		RepeatStartAt *time.Time
		StartAt       *time.Time
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "正常：取得できる（リピートあり）",
			fields: fields{
				RepeatSetting: true,
				RepeatStartAt: pointer.TimeToPointer(time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC)),
				StartAt:       nil,
			},
			args: args{
				now: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			want: 3,
		},
		{
			name: "正常：取得できる（リピートなし）",
			fields: fields{
				RepeatSetting: false,
				RepeatStartAt: nil,
				StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC)),
			},
			args: args{
				now: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := &Event{
				RepeatSetting: tt.fields.RepeatSetting,
				RepeatStartAt: tt.fields.RepeatStartAt,
				StartAt:       tt.fields.StartAt,
			}

			got := event.GetDayCount(tt.args.now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleService_IsEventPeriod(t *testing.T) {
	type fields struct {
		RepeatSetting bool
		RepeatStartAt *time.Time
		StartAt       *time.Time
		EndAt         *time.Time
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "正常：取得できる（リピートあり）",
			fields: fields{
				RepeatSetting: true,
				RepeatStartAt: pointer.TimeToPointer(time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC)),
				StartAt:       nil,
				EndAt:         nil,
			},
			args: args{
				now: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "正常：取得できる（リピートなし）",
			fields: fields{
				RepeatSetting: false,
				RepeatStartAt: nil,
				StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
				EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
			},
			args: args{
				now: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "異常：イベント期間外（リピートあり）",
			fields: fields{
				RepeatSetting: true,
				RepeatStartAt: pointer.TimeToPointer(time.Date(2023, 1, 6, 6, 0, 0, 0, time.UTC)),
				StartAt:       nil,
				EndAt:         nil,
			},
			args: args{
				now: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "異常：イベント開始前（リピートなし）",
			fields: fields{
				RepeatSetting: false,
				RepeatStartAt: nil,
				StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 6, 9, 0, 0, 0, time.UTC)),
				EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
			},
			args: args{
				now: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "異常：イベント終了後（リピートなし）",
			fields: fields{
				RepeatSetting: false,
				RepeatStartAt: nil,
				StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 6, 9, 0, 0, 0, time.UTC)),
				EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
			},
			args: args{
				now: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := &Event{
				RepeatSetting: tt.fields.RepeatSetting,
				RepeatStartAt: tt.fields.RepeatStartAt,
				StartAt:       tt.fields.StartAt,
				EndAt:         tt.fields.EndAt,
			}

			got := event.IsEventPeriod(tt.args.now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
