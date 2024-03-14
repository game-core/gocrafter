package idleBonus

import (
	"reflect"
	"testing"

	idleBonusService "github.com/game-core/gocrafter/pkg/domain/model/idleBonus"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

func TestIdleBonusUsecase_NewIdleBonusUsecase(t *testing.T) {
	type args struct {
		idleBonusService   idleBonusService.IdleBonusService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want IdleBonusUsecase
	}{
		{
			name: "正常",
			args: args{
				idleBonusService:   nil,
				transactionService: nil,
			},
			want: &idleBonusUsecase{
				idleBonusService:   nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIdleBonusUsecase(tt.args.idleBonusService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdleBonusUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
