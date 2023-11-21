package account

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	request "github.com/game-core/gocrafter/api/presentation/request/account"
	response "github.com/game-core/gocrafter/api/presentation/response/account"
	accountEntity "github.com/game-core/gocrafter/domain/entity/user/account"
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
					m.EXPECT().
						Begin().
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Commit(
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				accountRepository: func(ctrl *gomock.Controller) accountRepository.AccountRepository {
					m := accountRepository.NewMockAccountRepository(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							nil,
						).
						Return(
							&accountEntity.Account{
								UUID:      "uuid",
								Name:      "name",
								Password:  "password",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
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
					ID:       0,
					UUID:     "uuid",
					Name:     "name",
					Password: "password",
					Token:    "",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー(accountRepository.Create)",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin().
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Rollback(
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				accountRepository: func(ctrl *gomock.Controller) accountRepository.AccountRepository {
					m := accountRepository.NewMockAccountRepository(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							nil,
						).
						Return(
							nil,
							errors.New("accountRepository.Create"),
						)
					return m
				},
			},
			args: args{
				req: &request.RegisterAccount{
					Name: "name",
				},
			},
			want:    nil,
			wantErr: errors.New("accountRepository.Create"),
		},
		{
			name: "異常：エラー（transactionRepository.Begin）",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin().
						Return(
							nil,
							errors.New("transactionRepository.Begin"),
						)
					return m
				},
				accountRepository: func(ctrl *gomock.Controller) accountRepository.AccountRepository {
					m := accountRepository.NewMockAccountRepository(ctrl)
					return m
				},
			},
			args: args{
				req: &request.RegisterAccount{
					Name: "name",
				},
			},
			want:    nil,
			wantErr: errors.New("transactionRepository.Begin"),
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
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("RegisterAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				got.Item.Password = "password"
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_LoginAccount(t *testing.T) {
	type fields struct {
		accountRepository func(ctrl *gomock.Controller) accountRepository.AccountRepository
	}
	type args struct {
		req *request.LoginAccount
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.LoginAccount
		wantErr error
	}{
		{
			name: "正常：ログインできる",
			fields: fields{
				accountRepository: func(ctrl *gomock.Controller) accountRepository.AccountRepository {
					m := accountRepository.NewMockAccountRepository(ctrl)
					m.EXPECT().
						FindByUUID(
							"uuid",
						).
						Return(
							&accountEntity.Account{
								UUID:      "uuid",
								Name:      "name",
								Password:  "$2a$10$DHKndG0mMDkIy2G0p4H2f.YsrxX5TZdwqlB9eWO8xEwvhdlErS3Kq",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				req: &request.LoginAccount{
					UUID:     "uuid",
					Password: "txbw-V8xmREN12sx88zo",
				},
			},
			want: &response.LoginAccount{
				Status: 200,
				Item: response.Account{
					ID:       0,
					UUID:     "uuid",
					Name:     "name",
					Password: "txbw-V8xmREN12sx88zo",
					Token:    "token",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー（accountRepository.FindByUUID）",
			fields: fields{
				accountRepository: func(ctrl *gomock.Controller) accountRepository.AccountRepository {
					m := accountRepository.NewMockAccountRepository(ctrl)
					m.EXPECT().
						FindByUUID(
							"uuid",
						).
						Return(
							nil,
							errors.New("accountRepository.FindByUUID"),
						)
					return m
				},
			},
			args: args{
				req: &request.LoginAccount{
					UUID:     "uuid",
					Password: "txbw-V8xmREN12sx88zo",
				},
			},
			want:    nil,
			wantErr: errors.New("accountRepository.FindByUUID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				accountRepository: tt.fields.accountRepository(ctrl),
			}

			got, err := s.LoginAccount(tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("RegisterAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				got.Item.Token = "token"
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_CheckAccount(t *testing.T) {
	type fields struct {
		accountRepository func(ctrl *gomock.Controller) accountRepository.AccountRepository
	}
	type args struct {
		req *request.CheckAccount
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.CheckAccount
		wantErr error
	}{
		{
			name: "正常：確認できる",
			fields: fields{
				accountRepository: func(ctrl *gomock.Controller) accountRepository.AccountRepository {
					m := accountRepository.NewMockAccountRepository(ctrl)
					m.EXPECT().
						FindByUUID(
							"uuid",
						).
						Return(
							&accountEntity.Account{
								UUID:      "uuid",
								Name:      "name",
								Password:  "$2a$10$DHKndG0mMDkIy2G0p4H2f.YsrxX5TZdwqlB9eWO8xEwvhdlErS3Kq",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				req: &request.CheckAccount{
					UUID: "uuid",
				},
			},
			want: &response.CheckAccount{
				Status: 200,
				Item: response.Account{
					ID:       0,
					UUID:     "uuid",
					Name:     "name",
					Password: "",
					Token:    "",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー（accountRepository.FindByUUID）",
			fields: fields{
				accountRepository: func(ctrl *gomock.Controller) accountRepository.AccountRepository {
					m := accountRepository.NewMockAccountRepository(ctrl)
					m.EXPECT().
						FindByUUID(
							"uuid",
						).
						Return(
							nil,
							errors.New("accountRepository.FindByUUID"),
						)
					return m
				},
			},
			args: args{
				req: &request.CheckAccount{
					UUID: "uuid",
				},
			},
			want:    nil,
			wantErr: errors.New("accountRepository.FindByUUID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				accountRepository: tt.fields.accountRepository(ctrl),
			}

			got, err := s.CheckAccount(tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("RegisterAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
