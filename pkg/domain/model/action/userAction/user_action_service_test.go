package userAction

import (
	"reflect"
	"testing"
	"time"

	"github.com/game-core/gocrafter/internal/pointers"
)

func TestUserAction_CheckExpiration(t *testing.T) {
	type fields struct {
		UserAction *UserAction
	}
	type args struct {
		now        time.Time
		expiration *int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "正常：無期限",
			fields: fields{
				UserAction: &UserAction{
					UserId:         "0:0000",
					MasterActionId: 1,
					StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				expiration: nil,
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "正常：期限内",
			fields: fields{
				UserAction: &UserAction{
					UserId:         "0:0000",
					MasterActionId: 1,
					StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				expiration: pointers.Int32ToPointer(24),
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "正常：期限外",
			fields: fields{
				UserAction: &UserAction{
					UserId:         "0:0000",
					MasterActionId: 1,
					StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				now:        time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC),
				expiration: pointers.Int32ToPointer(24),
			},
			want:    false,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.UserAction.CheckExpiration(tt.args.now, tt.args.expiration)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckExpiration() = %v, want %v", got, tt.want)
			}
		})
	}
}
