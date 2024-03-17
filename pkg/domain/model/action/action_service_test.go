package action

import (
	"reflect"
	"testing"

	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionRun"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionStep"
	"github.com/game-core/gocrafter/pkg/domain/model/action/userAction"
)

func TestNewActionService_NewActionService(t *testing.T) {
	type args struct {
		masterActionRepository     masterAction.MasterActionRepository
		masterActionRunRepository  masterActionRun.MasterActionRunRepository
		masterActionStepRepository masterActionStep.MasterActionStepRepository
		userActionRepository       userAction.UserActionRepository
	}
	tests := []struct {
		name string
		args args
		want ActionService
	}{
		{
			name: "正常",
			args: args{
				masterActionRepository:     nil,
				masterActionRunRepository:  nil,
				masterActionStepRepository: nil,
				userActionRepository:       nil,
			},
			want: &actionService{
				masterActionRepository:     nil,
				masterActionRunRepository:  nil,
				masterActionStepRepository: nil,
				userActionRepository:       nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewActionService(
				tt.args.masterActionRepository,
				tt.args.masterActionRunRepository,
				tt.args.masterActionStepRepository,
				tt.args.userActionRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActionService() = %v, want %v", got, tt.want)
			}
		})
	}
}
