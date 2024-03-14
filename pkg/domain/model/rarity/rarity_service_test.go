package rarity

import (
	"reflect"
	"testing"

	"github.com/game-core/gocrafter/pkg/domain/model/rarity/masterRarity"
)

func TestNewRarityService_NewRarityService(t *testing.T) {
	type args struct {
		masterRarityRepository masterRarity.MasterRarityRepository
	}
	tests := []struct {
		name string
		args args
		want RarityService
	}{
		{
			name: "正常",
			args: args{
				masterRarityRepository: nil,
			},
			want: &rarityService{
				masterRarityRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRarityService(
				tt.args.masterRarityRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRarityService() = %v, want %v", got, tt.want)
			}
		})
	}
}
