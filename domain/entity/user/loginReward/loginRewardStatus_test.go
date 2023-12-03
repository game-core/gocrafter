package loginReward

import (
	"reflect"
	"testing"
	"time"
)

func TestLoginRewardStatusEntity_HasReceived(t *testing.T) {
	type fields struct {
		status *LoginRewardStatus
	}
	type args struct {
		now       time.Time
		resetHour int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "正常：受け取っていない",
			fields: fields{
				status: &LoginRewardStatus{
					LastReceivedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:       time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				resetHour: 9,
			},
			want: true,
		},
		{
			name: "正常：受け取っていない",
			fields: fields{
				status: &LoginRewardStatus{
					LastReceivedAt: time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:       time.Date(2023, 1, 2, 8, 0, 0, 0, time.UTC),
				resetHour: 9,
			},
			want: false,
		},
		{
			name: "正常：受け取り済み",
			fields: fields{
				status: &LoginRewardStatus{
					LastReceivedAt: time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:       time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				resetHour: 9,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.status.HasReceived(tt.args.now, tt.args.resetHour)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasReceived() = %v, want %v", got, tt.want)
			}
		})
	}
}
