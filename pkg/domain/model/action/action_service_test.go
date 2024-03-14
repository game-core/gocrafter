package action

import (
	"reflect"
	"testing"

	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
	"github.com/game-core/gocrafter/pkg/domain/model/action/userAction"
)

func TestNewActionService_NewActionService(t *testing.T) {
	type args struct {
		masterActionRepository masterAction.MasterActionRepository
		userActionRepository   userAction.UserActionRepository
	}
	tests := []struct {
		name string
		args args
		want ActionService
	}{
		{
			name: "正常",
			args: args{
				masterActionRepository: nil,
				userActionRepository:   nil,
			},
			want: &actionService{
				masterActionRepository: nil,
				userActionRepository:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewActionService(
				tt.args.masterActionRepository,
				tt.args.userActionRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActionService() = %v, want %v", got, tt.want)
			}
		})
	}
}
