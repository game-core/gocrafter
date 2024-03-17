package userLoginBonus

import (
	"reflect"
	"testing"
	"time"
)

func TestUserLoginBonus_CheckReceived(t *testing.T) {
	type fields struct {
		UserLoginBonus *UserLoginBonus
	}
	type args struct {
		resetHour int32
		now       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "正常：未受け取りの場合",
			fields: fields{
				UserLoginBonus: &UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				resetHour: 9,
				now:       time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
			},
			want:    false,
			wantErr: nil,
		},
		{
			name: "正常：未受け取りの場合",
			fields: fields{
				UserLoginBonus: &UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         time.Date(2023, 1, 2, 8, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				resetHour: 9,
				now:       time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
			},
			want:    false,
			wantErr: nil,
		},
		{
			name: "正常：受け取り済みの場合",
			fields: fields{
				UserLoginBonus: &UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				resetHour: 9,
				now:       time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "正常：受け取り済みの場合",
			fields: fields{
				UserLoginBonus: &UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				resetHour: 9,
				now:       time.Date(2023, 1, 3, 8, 0, 0, 0, time.UTC),
			},
			want:    true,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.UserLoginBonus.CheckReceived(tt.args.resetHour, tt.args.now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckReceived() = %v, want %v", got, tt.want)
			}
		})
	}
}
