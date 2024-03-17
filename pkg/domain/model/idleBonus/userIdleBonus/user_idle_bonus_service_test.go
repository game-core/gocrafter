package userIdleBonus

import (
	"reflect"
	"testing"
	"time"
)

func TestUserIdleBonus_GetReceivedAt(t *testing.T) {
	type fields struct {
		UserIdleBonus *UserIdleBonus
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    time.Time
		wantErr error
	}{
		{
			name: "正常：初回",
			fields: fields{
				UserIdleBonus: nil,
			},
			args: args{
				now: time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
			},
			want:    time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
			wantErr: nil,
		},
		{
			name: "正常：２回目以降",
			fields: fields{
				UserIdleBonus: &UserIdleBonus{
					ReceivedAt: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now: time.Date(2023, 1, 2, 9, 0, 0, 0, time.UTC),
			},
			want:    time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.UserIdleBonus.GetReceivedAt(tt.args.now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReceivedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
