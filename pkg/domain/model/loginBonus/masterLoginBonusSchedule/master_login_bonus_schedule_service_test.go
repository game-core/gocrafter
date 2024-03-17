package masterLoginBonusSchedule

import (
	"reflect"
	"testing"
	"time"
)

func TestMasterLoginBonusSchedules_GetSchedulesByStep(t *testing.T) {
	type fields struct {
		MasterLoginBonusSchedules MasterLoginBonusSchedules
	}
	type args struct {
		step int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *MasterLoginBonusSchedule
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				MasterLoginBonusSchedules: MasterLoginBonusSchedules{
					{
						Id:                 1,
						MasterLoginBonusId: 1,
						Step:               0,
						Name:               "スケジュール1",
					},
					{
						Id:                 2,
						MasterLoginBonusId: 1,
						Step:               1,
						Name:               "スケジュール2",
					},
					{
						Id:                 3,
						MasterLoginBonusId: 1,
						Step:               2,
						Name:               "スケジュール3",
					},
				},
			},
			args: args{
				step: 0,
			},
			want: &MasterLoginBonusSchedule{
				Id:                 1,
				MasterLoginBonusId: 1,
				Step:               0,
				Name:               "スケジュール1",
			},
			wantErr: nil,
		},
		{
			name: "正常：取得できる",
			fields: fields{
				MasterLoginBonusSchedules: MasterLoginBonusSchedules{
					{
						Id:                 1,
						MasterLoginBonusId: 1,
						Step:               0,
						Name:               "スケジュール1",
					},
					{
						Id:                 2,
						MasterLoginBonusId: 1,
						Step:               1,
						Name:               "スケジュール2",
					},
					{
						Id:                 3,
						MasterLoginBonusId: 1,
						Step:               2,
						Name:               "スケジュール3",
					},
				},
			},
			args: args{
				step: 1,
			},
			want: &MasterLoginBonusSchedule{
				Id:                 2,
				MasterLoginBonusId: 1,
				Step:               1,
				Name:               "スケジュール2",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.MasterLoginBonusSchedules.GetScheduleByStep(tt.args.step)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetScheduleByStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasterLoginBonusSchedules_GetStep(t *testing.T) {
	type fields struct {
		MasterLoginBonusSchedules MasterLoginBonusSchedules
	}
	type args struct {
		intervalHour int32
		receivedAt   time.Time
		now          time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{
		{
			name: "正常：0",
			fields: fields{
				MasterLoginBonusSchedules: MasterLoginBonusSchedules{
					{
						Id:                 1,
						MasterLoginBonusId: 1,
						Step:               0,
						Name:               "スケジュール1",
					},
					{
						Id:                 2,
						MasterLoginBonusId: 1,
						Step:               1,
						Name:               "スケジュール2",
					},
					{
						Id:                 3,
						MasterLoginBonusId: 1,
						Step:               2,
						Name:               "スケジュール3",
					},
				},
			},
			args: args{
				intervalHour: 24,
				receivedAt:   time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				now:          time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "正常：1",
			fields: fields{
				MasterLoginBonusSchedules: MasterLoginBonusSchedules{
					{
						Id:                 1,
						MasterLoginBonusId: 1,
						Step:               0,
						Name:               "スケジュール1",
					},
					{
						Id:                 2,
						MasterLoginBonusId: 1,
						Step:               1,
						Name:               "スケジュール2",
					},
					{
						Id:                 3,
						MasterLoginBonusId: 1,
						Step:               2,
						Name:               "スケジュール3",
					},
				},
			},
			args: args{
				intervalHour: 24,
				receivedAt:   time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				now:          time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
			},
			want: 1,
		},
		{
			name: "正常：2",
			fields: fields{
				MasterLoginBonusSchedules: MasterLoginBonusSchedules{
					{
						Id:                 1,
						MasterLoginBonusId: 1,
						Step:               0,
						Name:               "スケジュール1",
					},
					{
						Id:                 2,
						MasterLoginBonusId: 1,
						Step:               1,
						Name:               "スケジュール2",
					},
					{
						Id:                 3,
						MasterLoginBonusId: 1,
						Step:               2,
						Name:               "スケジュール3",
					},
				},
			},
			args: args{
				intervalHour: 24,
				receivedAt:   time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				now:          time.Date(2023, 1, 3, 9, 0, 0, 0, time.UTC),
			},
			want: 2,
		},
		{
			name: "正常：3（２週目以降）",
			fields: fields{
				MasterLoginBonusSchedules: MasterLoginBonusSchedules{
					{
						Id:                 1,
						MasterLoginBonusId: 1,
						Step:               0,
						Name:               "スケジュール1",
					},
					{
						Id:                 2,
						MasterLoginBonusId: 1,
						Step:               1,
						Name:               "スケジュール2",
					},
					{
						Id:                 3,
						MasterLoginBonusId: 1,
						Step:               2,
						Name:               "スケジュール3",
					},
				},
			},
			args: args{
				intervalHour: 24,
				receivedAt:   time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				now:          time.Date(2023, 1, 4, 9, 0, 0, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "正常：4（２週目以降）",
			fields: fields{
				MasterLoginBonusSchedules: MasterLoginBonusSchedules{
					{
						Id:                 1,
						MasterLoginBonusId: 1,
						Step:               0,
						Name:               "スケジュール1",
					},
					{
						Id:                 2,
						MasterLoginBonusId: 1,
						Step:               1,
						Name:               "スケジュール2",
					},
					{
						Id:                 3,
						MasterLoginBonusId: 1,
						Step:               2,
						Name:               "スケジュール3",
					},
				},
			},
			args: args{
				intervalHour: 24,
				receivedAt:   time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				now:          time.Date(2023, 1, 5, 9, 0, 0, 0, time.UTC),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.MasterLoginBonusSchedules.GetStep(tt.args.intervalHour, tt.args.receivedAt, tt.args.now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
