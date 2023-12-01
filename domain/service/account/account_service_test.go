package account

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	request "github.com/game-core/gocrafter/api/presentation/request/account"
	response "github.com/game-core/gocrafter/api/presentation/response/account"
	shardResponse "github.com/game-core/gocrafter/api/presentation/response/shard"
	accountEntity "github.com/game-core/gocrafter/domain/entity/user/account"
	userRepository "github.com/game-core/gocrafter/domain/repository/user"
	accountRepository "github.com/game-core/gocrafter/domain/repository/user/account"
	shardService "github.com/game-core/gocrafter/domain/service/shard"
)

func TestAccountService_RegisterAccount(t *testing.T) {
	type fields struct {
		shardService          func(ctrl *gomock.Controller) shardService.ShardService
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
				shardService: func(ctrl *gomock.Controller) shardService.ShardService {
					m := shardService.NewMockShardService(ctrl)
					m.EXPECT().
						GetShard().
						Return(
							&shardResponse.GetShard{
								Status:       200,
								NextShardKey: "SHARD_1",
								Shards: &shardResponse.Shards{
									{
										ID:       1,
										ShardKey: "SHARD_1",
										Name:     "name1",
										Count:    1,
									},
									{
										ID:       2,
										ShardKey: "SHARD_2",
										Name:     "name2",
										Count:    2,
									},
								},
							},
							nil,
						)
					return m
				},
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin("SHARD_1").
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommitOrRollback(
							gomock.Any(),
							nil,
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
							"SHARD_1",
							nil,
						).
						Return(
							&accountEntity.Account{
								UUID:      "uuid",
								ShardKey:  "SHARD_1",
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
				Account: response.Account{
					ID:       0,
					ShardKey: "SHARD_1",
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
				shardService: func(ctrl *gomock.Controller) shardService.ShardService {
					m := shardService.NewMockShardService(ctrl)
					m.EXPECT().
						GetShard().
						Return(
							&shardResponse.GetShard{
								Status:       200,
								NextShardKey: "SHARD_1",
								Shards: &shardResponse.Shards{
									{
										ID:       1,
										ShardKey: "SHARD_1",
										Name:     "name1",
										Count:    1,
									},
									{
										ID:       2,
										ShardKey: "SHARD_2",
										Name:     "name2",
										Count:    2,
									},
								},
							},
							nil,
						)
					return m
				},
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin(
							"SHARD_1",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommitOrRollback(
							gomock.Any(),
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
							"SHARD_1",
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
				shardService: func(ctrl *gomock.Controller) shardService.ShardService {
					m := shardService.NewMockShardService(ctrl)
					m.EXPECT().
						GetShard().
						Return(
							&shardResponse.GetShard{
								Status:       200,
								NextShardKey: "SHARD_1",
								Shards: &shardResponse.Shards{
									{
										ID:       1,
										ShardKey: "SHARD_1",
										Name:     "name1",
										Count:    1,
									},
									{
										ID:       2,
										ShardKey: "SHARD_2",
										Name:     "name2",
										Count:    2,
									},
								},
							},
							nil,
						)
					return m
				},
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin("SHARD_1").
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
		{
			name: "異常：エラー（shardService.GetShard）",
			fields: fields{
				shardService: func(ctrl *gomock.Controller) shardService.ShardService {
					m := shardService.NewMockShardService(ctrl)
					m.EXPECT().
						GetShard().
						Return(
							nil,
							errors.New("shardService.GetShard"),
						)
					return m
				},
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
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
			wantErr: errors.New("shardService.GetShard"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				shardService:          tt.fields.shardService(ctrl),
				transactionRepository: tt.fields.transactionRepository(ctrl),
				accountRepository:     tt.fields.accountRepository(ctrl),
			}

			got, err := s.RegisterAccount(tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("RegisterAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				got.Account.Password = "password"
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
							"SHARD_1",
						).
						Return(
							&accountEntity.Account{
								UUID:      "uuid",
								ShardKey:  "SHARD_1",
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
					ShardKey: "SHARD_1",
					UUID:     "uuid",
					Password: "txbw-V8xmREN12sx88zo",
				},
			},
			want: &response.LoginAccount{
				Status: 200,
				Account: response.Account{
					ID:       0,
					ShardKey: "SHARD_1",
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
							"SHARD_1",
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
					ShardKey: "SHARD_1",
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
				t.Errorf("LoginAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				got.Account.Token = "token"
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginAccount() = %v, want %v", got, tt.want)
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
							"SHARD_1",
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
					ShardKey: "SHARD_1",
					UUID:     "uuid",
				},
			},
			want: &response.CheckAccount{
				Status: 200,
				Account: response.Account{
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
							"SHARD_1",
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
					ShardKey: "SHARD_1",
					UUID:     "uuid",
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
				t.Errorf("CheckAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
