package room

import (
	"reflect"
	"testing"

	roomService "github.com/game-core/gocrafter/pkg/domain/model/room"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

func TestRoomUsecase_NewRoomUsecase(t *testing.T) {
	type args struct {
		roomService        roomService.RoomService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want RoomUsecase
	}{
		{
			name: "正常",
			args: args{
				roomService:        nil,
				transactionService: nil,
			},
			want: &roomUsecase{
				roomService:        nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRoomUsecase(tt.args.roomService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoomUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
