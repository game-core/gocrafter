package account

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	request "github.com/game-core/gocrafter/api/presentation/request/account"
	response "github.com/game-core/gocrafter/api/presentation/response/account"
	userRepository "github.com/game-core/gocrafter/domain/repository/user"
	accountRepository "github.com/game-core/gocrafter/domain/repository/user/account"
)

func TestAccountService_RegisterAccount(t *testing.T) {
	type fields struct {
		transactionRepository func(ctrl *gomock.Controller) userRepository.TransactionRepository
		accountRepository     func(ctrl *gomock.Controller) accountRepository.AccountRepository
	}
	type args struct {
		req *request.RegisterAccount
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.RegisterAccount
		wantErr error
	}{
		{
			name: "正常：登録できる",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					return m
				},
			},
			args: args{
				req: &request.RegisterAccount{
					Name: "name",
				},
			},
			want: &response.RegisterAccount{
				Status: 200,
				Item: response.Account{
					ID:       1,
					UUID:     "123",
					Name:     "name",
					Password: "plaintextpassword",
					Token:    "",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				transactionRepository: tt.fields.transactionRepository(ctrl),
				accountRepository:     tt.fields.accountRepository(ctrl),
			}

			got, err := s.RegisterAccount(tt.args.req)
			if err != nil {
				t.Errorf("RegisterAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
