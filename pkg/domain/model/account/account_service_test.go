package account

import (
	"context"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
	"github.com/game-core/gocrafter/pkg/domain/model/shard"
)

func TestNewAccountService_NewAccountService(t *testing.T) {
	type args struct {
		shardService          shard.ShardService
		userAccountRepository userAccount.UserAccountRepository
	}
	tests := []struct {
		name string
		args args
		want AccountService
	}{
		{
			name: "正常",
			args: args{
				shardService:          nil,
				userAccountRepository: nil,
			},
			want: &accountService{
				shardService:          nil,
				userAccountRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAccountService(
				tt.args.shardService,
				tt.args.userAccountRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_FindByUserId(t *testing.T) {
	type fields struct {
		userAccountRepository func(ctrl *gomock.Controller) userAccount.UserAccountRepository
		shardService          func(ctrl *gomock.Controller) shard.ShardService
	}
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *userAccount.UserAccount
		wantErr error
	}{
		{
			name: "正常：アカウントが存在する場合",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
						).
						Return(
							&userAccount.UserAccount{
								UserId:   "0:WntR-PyhOJeDiE5jodeR",
								Name:     "test_user_account",
								Password: "iaOI3vb2Ea8JadVjZtVH",
								LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
			},
			args: args{
				ctx:    nil,
				userId: "0:WntR-PyhOJeDiE5jodeR",
			},
			want: &userAccount.UserAccount{
				UserId:   "0:WntR-PyhOJeDiE5jodeR",
				Name:     "test_user_account",
				Password: "iaOI3vb2Ea8JadVjZtVH",
				LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userAccountRepository.Find",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
			},
			args: args{
				ctx:    nil,
				userId: "0:WntR-PyhOJeDiE5jodeR",
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountRepository.Find", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountRepository: tt.fields.userAccountRepository(ctrl),
				shardService:          tt.fields.shardService(ctrl),
			}

			got, err := s.FindByUserId(tt.args.ctx, tt.args.userId)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("FindByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_Create(t *testing.T) {
	type fields struct {
		userAccountRepository func(ctrl *gomock.Controller) userAccount.UserAccountRepository
		shardService          func(ctrl *gomock.Controller) shard.ShardService
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *AccountCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountCreateResponse
		wantErr error
	}{
		{
			name: "正常：作成できる場合",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					m.EXPECT().
						Create(
							nil,
							nil,
							gomock.Any(),
						).
						Return(
							&userAccount.UserAccount{
								UserId:   "0:WntR-PyhOJeDiE5jodeR",
								Name:     "test_user_account",
								Password: "iaOI3vb2Ea8JadVjZtVH",
								LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &AccountCreateRequest{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
					Name:   "test_user_account",
				},
			},
			want: &AccountCreateResponse{
				UserAccount: &userAccount.UserAccount{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "iaOI3vb2Ea8JadVjZtVH",
					LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userAccountRepository.Create",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					m.EXPECT().
						Create(
							nil,
							nil,
							gomock.Any(),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &AccountCreateRequest{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
					Name:   "test_user_account",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountRepository.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountRepository: tt.fields.userAccountRepository(ctrl),
				shardService:          tt.fields.shardService(ctrl),
			}

			got, err := s.Create(tt.args.ctx, tt.args.tx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				got.UserAccount.Password = "iaOI3vb2Ea8JadVjZtVH"
				got.UserAccount.LoginAt = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
				got.UserAccount.LogoutAt = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_Login(t *testing.T) {
	type fields struct {
		userAccountRepository func(ctrl *gomock.Controller) userAccount.UserAccountRepository
		shardService          func(ctrl *gomock.Controller) shard.ShardService
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *AccountLoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountLoginResponse
		wantErr error
	}{
		{
			name: "正常：ログインできる場合",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
						).
						Return(
							&userAccount.UserAccount{
								UserId:   "0:WntR-PyhOJeDiE5jodeR",
								Name:     "test_user_account",
								Password: "$2a$10$J/phCDt8nXe02rhhPcDx1.9LEX3jw4mXxQq2ulKhuWcFmllaWSSqm",
								LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							gomock.Any(),
						).
						Return(
							&userAccount.UserAccount{
								UserId:   "0:WntR-PyhOJeDiE5jodeR",
								Name:     "test_user_account",
								Password: "$2a$10$J/phCDt8nXe02rhhPcDx1.9LEX3jw4mXxQq2ulKhuWcFmllaWSSqm",
								LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &AccountLoginRequest{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "M3sDUx-vXLGfYzhGz7Nl",
				},
			},
			want: &AccountLoginResponse{
				Token: "token",
				UserAccount: &userAccount.UserAccount{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "$2a$10$J/phCDt8nXe02rhhPcDx1.9LEX3jw4mXxQq2ulKhuWcFmllaWSSqm",
					LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userAccountRepository.Find",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &AccountLoginRequest{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "M3sDUx-vXLGfYzhGz7Nl",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.userAccountRepository.Update",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
						).
						Return(
							&userAccount.UserAccount{
								UserId:   "0:WntR-PyhOJeDiE5jodeR",
								Name:     "test_user_account",
								Password: "$2a$10$J/phCDt8nXe02rhhPcDx1.9LEX3jw4mXxQq2ulKhuWcFmllaWSSqm",
								LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							gomock.Any(),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &AccountLoginRequest{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "M3sDUx-vXLGfYzhGz7Nl",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountRepository.Update", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountRepository: tt.fields.userAccountRepository(ctrl),
				shardService:          tt.fields.shardService(ctrl),
			}

			got, err := s.Login(tt.args.ctx, tt.args.tx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				got.Token = "token"
				got.UserAccount.Password = "$2a$10$J/phCDt8nXe02rhhPcDx1.9LEX3jw4mXxQq2ulKhuWcFmllaWSSqm"
				got.UserAccount.LoginAt = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
				got.UserAccount.LogoutAt = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_Check(t *testing.T) {
	type fields struct {
		userAccountRepository func(ctrl *gomock.Controller) userAccount.UserAccountRepository
		shardService          func(ctrl *gomock.Controller) shard.ShardService
	}
	type args struct {
		ctx context.Context
		req *AccountCheckRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountCheckResponse
		wantErr error
	}{
		{
			name: "正常：アカウントが存在する場合",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
						).
						Return(
							&userAccount.UserAccount{
								UserId:   "0:WntR-PyhOJeDiE5jodeR",
								Name:     "test_user_account",
								Password: "iaOI3vb2Ea8JadVjZtVH",
								LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &AccountCheckRequest{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
				},
			},
			want: &AccountCheckResponse{
				&userAccount.UserAccount{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "iaOI3vb2Ea8JadVjZtVH",
					LoginAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					LogoutAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userAccountRepository.Find",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &AccountCheckRequest{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountRepository.Find", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountRepository: tt.fields.userAccountRepository(ctrl),
				shardService:          tt.fields.shardService(ctrl),
			}

			got, err := s.Check(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_GenerateUserID(t *testing.T) {
	type fields struct {
		userAccountRepository func(ctrl *gomock.Controller) userAccount.UserAccountRepository
		shardService          func(ctrl *gomock.Controller) shard.ShardService
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{
			name: "正常：アカウントが存在する場合",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					m.EXPECT().
						GetShardKey(
							nil,
							nil,
						).
						Return(
							"0",
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    "0:WntR-PyhOJeDiE5jodeR",
			wantErr: nil,
		},
		{
			name: "異常：s.shardService.GetShardKey",
			fields: fields{
				userAccountRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRepository {
					m := userAccount.NewMockUserAccountRepository(ctrl)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					m.EXPECT().
						GetShardKey(
							nil,
							nil,
						).
						Return(
							"",
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    "",
			wantErr: errors.NewMethodError("s.shardService.GetShardKey", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountRepository: tt.fields.userAccountRepository(ctrl),
				shardService:          tt.fields.shardService(ctrl),
			}

			got, err := s.GenerateUserID(tt.args.ctx)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GenerateUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				got = "0:WntR-PyhOJeDiE5jodeR"
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
