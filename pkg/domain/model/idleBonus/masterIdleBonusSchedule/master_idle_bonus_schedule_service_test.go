package masterIdleBonusSchedule

import (
	"reflect"
	"testing"
	"time"

	"github.com/game-core/gocrafter/internal/errors"
)

func TestMasterIdleBonusSchedules_GetSchedulesByStep(t *testing.T) {
	type fields struct {
		MasterIdleBonusSchedules MasterIdleBonusSchedules
	}
	type args struct {
		step int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MasterIdleBonusSchedules
		wantErr error
	}{
		{
			name: "正常：1回",
			fields: fields{
				MasterIdleBonusSchedules: MasterIdleBonusSchedules{
					{
						Id:                1,
						MasterIdleBonusId: 1,
						Step:              0,
						Name:              "スケジュール1",
					},
					{
						Id:                2,
						MasterIdleBonusId: 1,
						Step:              1,
						Name:              "スケジュール2",
					},
					{
						Id:                3,
						MasterIdleBonusId: 1,
						Step:              2,
						Name:              "スケジュール3",
					},
				},
			},
			args: args{
				step: 1,
			},
			want: MasterIdleBonusSchedules{
				{
					Id:                1,
					MasterIdleBonusId: 1,
					Step:              0,
					Name:              "スケジュール1",
				},
				{
					Id:                2,
					MasterIdleBonusId: 1,
					Step:              1,
					Name:              "スケジュール2",
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：0回",
			fields: fields{
				MasterIdleBonusSchedules: MasterIdleBonusSchedules{
					{
						Id:                1,
						MasterIdleBonusId: 1,
						Step:              0,
						Name:              "スケジュール1",
					},
					{
						Id:                2,
						MasterIdleBonusId: 1,
						Step:              1,
						Name:              "スケジュール2",
					},
					{
						Id:                3,
						MasterIdleBonusId: 1,
						Step:              2,
						Name:              "スケジュール3",
					},
				},
			},
			args: args{
				step: 0,
			},
			want: MasterIdleBonusSchedules{
				{
					Id:                1,
					MasterIdleBonusId: 1,
					Step:              0,
					Name:              "スケジュール1",
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：最大",
			fields: fields{
				MasterIdleBonusSchedules: MasterIdleBonusSchedules{
					{
						Id:                1,
						MasterIdleBonusId: 1,
						Step:              0,
						Name:              "スケジュール1",
					},
					{
						Id:                2,
						MasterIdleBonusId: 1,
						Step:              1,
						Name:              "スケジュール2",
					},
					{
						Id:                3,
						MasterIdleBonusId: 1,
						Step:              2,
						Name:              "スケジュール3",
					},
				},
			},
			args: args{
				step: 10,
			},
			want: MasterIdleBonusSchedules{
				{
					Id:                1,
					MasterIdleBonusId: 1,
					Step:              0,
					Name:              "スケジュール1",
				},
				{
					Id:                2,
					MasterIdleBonusId: 1,
					Step:              1,
					Name:              "スケジュール2",
				},
				{
					Id:                3,
					MasterIdleBonusId: 1,
					Step:              2,
					Name:              "スケジュール3",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.MasterIdleBonusSchedules.GetSchedulesByStep(tt.args.step)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSchedulesByStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasterIdleBonusSchedules_GetStep(t *testing.T) {
	type fields struct {
		MasterIdleBonusSchedules MasterIdleBonusSchedules
	}
	type args struct {
		intervalHour int32
		receivedAt   time.Time
		now          time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr error
	}{
		{
			name: "正常：0",
			fields: fields{
				MasterIdleBonusSchedules: MasterIdleBonusSchedules{
					{
						Id:                1,
						MasterIdleBonusId: 1,
						Step:              0,
						Name:              "スケジュール1",
					},
					{
						Id:                2,
						MasterIdleBonusId: 1,
						Step:              1,
						Name:              "スケジュール2",
					},
					{
						Id:                3,
						MasterIdleBonusId: 1,
						Step:              2,
						Name:              "スケジュール3",
					},
				},
			},
			args: args{
				intervalHour: 1,
				receivedAt:   time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				now:          time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "正常：1",
			fields: fields{
				MasterIdleBonusSchedules: MasterIdleBonusSchedules{
					{
						Id:                1,
						MasterIdleBonusId: 1,
						Step:              0,
						Name:              "スケジュール1",
					},
					{
						Id:                2,
						MasterIdleBonusId: 1,
						Step:              1,
						Name:              "スケジュール2",
					},
					{
						Id:                3,
						MasterIdleBonusId: 1,
						Step:              2,
						Name:              "スケジュール3",
					},
				},
			},
			args: args{
				intervalHour: 1,
				receivedAt:   time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				now:          time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "正常：最大",
			fields: fields{
				MasterIdleBonusSchedules: MasterIdleBonusSchedules{
					{
						Id:                1,
						MasterIdleBonusId: 1,
						Step:              0,
						Name:              "スケジュール1",
					},
					{
						Id:                2,
						MasterIdleBonusId: 1,
						Step:              1,
						Name:              "スケジュール2",
					},
					{
						Id:                3,
						MasterIdleBonusId: 1,
						Step:              2,
						Name:              "スケジュール3",
					},
				},
			},
			args: args{
				intervalHour: 1,
				receivedAt:   time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				now:          time.Date(2023, 1, 1, 20, 0, 0, 0, time.UTC),
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "異常：already received",
			fields: fields{
				MasterIdleBonusSchedules: MasterIdleBonusSchedules{
					{
						Id:                1,
						MasterIdleBonusId: 1,
						Step:              0,
						Name:              "スケジュール1",
					},
					{
						Id:                2,
						MasterIdleBonusId: 1,
						Step:              1,
						Name:              "スケジュール2",
					},
					{
						Id:                3,
						MasterIdleBonusId: 1,
						Step:              2,
						Name:              "スケジュール3",
					},
				},
			},
			args: args{
				intervalHour: 1,
				receivedAt:   time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				now:          time.Date(2023, 1, 1, 9, 59, 59, 59, time.UTC),
			},
			want:    0,
			wantErr: errors.NewError("already received"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.MasterIdleBonusSchedules.GetStep(tt.args.intervalHour, tt.args.receivedAt, tt.args.now)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetStep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
