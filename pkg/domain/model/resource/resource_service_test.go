package resource

import (
	"reflect"
	"testing"

	"github.com/game-core/gocrafter/pkg/domain/model/resource/masterResource"
)

func TestNewResourceService_NewResourceService(t *testing.T) {
	type args struct {
		masterResourceRepository masterResource.MasterResourceRepository
	}
	tests := []struct {
		name string
		args args
		want ResourceService
	}{
		{
			name: "正常",
			args: args{
				masterResourceRepository: nil,
			},
			want: &resourceService{
				masterResourceRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewResourceService(
				tt.args.masterResourceRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourceService() = %v, want %v", got, tt.want)
			}
		})
	}
}
