package account

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccountToken"
	"github.com/game-core/gocrafter/pkg/domain/model/shard"
)

func TestNewAccountService_NewAccountService(t *testing.T) {
	type args struct {
		shardService                    shard.ShardService
		userAccountMysqlRepository      userAccount.UserAccountMysqlRepository
		userAccountRedisRepository      userAccount.UserAccountRedisRepository
		userAccountTokenRedisRepository userAccountToken.UserAccountTokenRedisRepository
	}
	tests := []struct {
		name string
		args args
		want AccountService
	}{
		{
			name: "正常",
			args: args{
				shardService:                    nil,
				userAccountMysqlRepository:      nil,
				userAccountRedisRepository:      nil,
				userAccountTokenRedisRepository: nil,
			},
			want: &accountService{
				shardService:                    nil,
				userAccountMysqlRepository:      nil,
				userAccountRedisRepository:      nil,
				userAccountTokenRedisRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAccountService(
				tt.args.shardService,
				tt.args.userAccountMysqlRepository,
				tt.args.userAccountRedisRepository,
				tt.args.userAccountTokenRedisRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_FindByUserId(t *testing.T) {
	type fields struct {
		userAccountMysqlRepository func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository
		shardService               func(ctrl *gomock.Controller) shard.ShardService
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
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
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
			name: "異常：s.userAccountMysqlRepository.Find",
			fields: fields{
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.userAccountMysqlRepository.Find", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountMysqlRepository: tt.fields.userAccountMysqlRepository(ctrl),
				shardService:               tt.fields.shardService(ctrl),
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
		userAccountMysqlRepository func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository
		shardService               func(ctrl *gomock.Controller) shard.ShardService
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
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
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
			name: "異常：s.userAccountMysqlRepository.Create",
			fields: fields{
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.userAccountMysqlRepository.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountMysqlRepository: tt.fields.userAccountMysqlRepository(ctrl),
				shardService:               tt.fields.shardService(ctrl),
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
		userAccountMysqlRepository      func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository
		userAccountRedisRepository      func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository
		userAccountTokenRedisRepository func(ctrl *gomock.Controller) userAccountToken.UserAccountTokenRedisRepository
		shardService                    func(ctrl *gomock.Controller) shard.ShardService
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		rtx redis.Pipeliner
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
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
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
				userAccountRedisRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository {
					m := userAccount.NewMockUserAccountRedisRepository(ctrl)
					m.EXPECT().
						Set(
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
				userAccountTokenRedisRepository: func(ctrl *gomock.Controller) userAccountToken.UserAccountTokenRedisRepository {
					m := userAccountToken.NewMockUserAccountTokenRedisRepository(ctrl)
					m.EXPECT().
						Set(
							nil,
							nil,
							gomock.Any(),
						).
						Return(
							&userAccountToken.UserAccountToken{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
								Token:  "token",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				rtx: nil,
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
			name: "異常：s.userAccountMysqlRepository.Find",
			fields: fields{
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
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
				userAccountRedisRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository {
					m := userAccount.NewMockUserAccountRedisRepository(ctrl)
					return m
				},
				userAccountTokenRedisRepository: func(ctrl *gomock.Controller) userAccountToken.UserAccountTokenRedisRepository {
					m := userAccountToken.NewMockUserAccountTokenRedisRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				rtx: nil,
				req: &AccountLoginRequest{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "M3sDUx-vXLGfYzhGz7Nl",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.userAccountMysqlRepository.Update",
			fields: fields{
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
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
				userAccountRedisRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository {
					m := userAccount.NewMockUserAccountRedisRepository(ctrl)
					return m
				},
				userAccountTokenRedisRepository: func(ctrl *gomock.Controller) userAccountToken.UserAccountTokenRedisRepository {
					m := userAccountToken.NewMockUserAccountTokenRedisRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				rtx: nil,
				req: &AccountLoginRequest{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "M3sDUx-vXLGfYzhGz7Nl",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountMysqlRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.userAccountRedisRepository.Set",
			fields: fields{
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
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
				userAccountRedisRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository {
					m := userAccount.NewMockUserAccountRedisRepository(ctrl)
					m.EXPECT().
						Set(
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
				userAccountTokenRedisRepository: func(ctrl *gomock.Controller) userAccountToken.UserAccountTokenRedisRepository {
					m := userAccountToken.NewMockUserAccountTokenRedisRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				rtx: nil,
				req: &AccountLoginRequest{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "M3sDUx-vXLGfYzhGz7Nl",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountRedisRepository.Set", errors.NewTestError()),
		},
		{
			name: "異常：s.userAccountTokenRedisRepository.Set",
			fields: fields{
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
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
				userAccountRedisRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository {
					m := userAccount.NewMockUserAccountRedisRepository(ctrl)
					m.EXPECT().
						Set(
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
				userAccountTokenRedisRepository: func(ctrl *gomock.Controller) userAccountToken.UserAccountTokenRedisRepository {
					m := userAccountToken.NewMockUserAccountTokenRedisRepository(ctrl)
					m.EXPECT().
						Set(
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
			},
			args: args{
				ctx: nil,
				tx:  nil,
				rtx: nil,
				req: &AccountLoginRequest{
					UserId:   "0:WntR-PyhOJeDiE5jodeR",
					Name:     "test_user_account",
					Password: "M3sDUx-vXLGfYzhGz7Nl",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountTokenRedisRepository.Set", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountMysqlRepository:      tt.fields.userAccountMysqlRepository(ctrl),
				userAccountRedisRepository:      tt.fields.userAccountRedisRepository(ctrl),
				userAccountTokenRedisRepository: tt.fields.userAccountTokenRedisRepository(ctrl),
				shardService:                    tt.fields.shardService(ctrl),
			}

			got, err := s.Login(tt.args.ctx, tt.args.tx, tt.args.rtx, tt.args.req)
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
		userAccountRedisRepository func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository
		shardService               func(ctrl *gomock.Controller) shard.ShardService
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
				userAccountRedisRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository {
					m := userAccount.NewMockUserAccountRedisRepository(ctrl)
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
			name: "異常：s.userAccountRedisRepository.Find",
			fields: fields{
				userAccountRedisRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository {
					m := userAccount.NewMockUserAccountRedisRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.userAccountRedisRepository.Find", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountRedisRepository: tt.fields.userAccountRedisRepository(ctrl),
				shardService:               tt.fields.shardService(ctrl),
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

func TestAccountService_CheckToken(t *testing.T) {
	type fields struct {
		userAccountMysqlRepository      func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository
		userAccountRedisRepository      func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository
		userAccountTokenRedisRepository func(ctrl *gomock.Controller) userAccountToken.UserAccountTokenRedisRepository
		shardService                    func(ctrl *gomock.Controller) shard.ShardService
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		rtx redis.Pipeliner
		req *AccountCheckTokenRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountCheckTokenResponse
		wantErr error
	}{
		{
			name: "正常：確認できる場合",
			fields: fields{
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
				userAccountRedisRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository {
					m := userAccount.NewMockUserAccountRedisRepository(ctrl)
					return m
				},
				userAccountTokenRedisRepository: func(ctrl *gomock.Controller) userAccountToken.UserAccountTokenRedisRepository {
					m := userAccountToken.NewMockUserAccountTokenRedisRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
						).
						Return(
							&userAccountToken.UserAccountToken{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
								Token:  "token",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				rtx: nil,
				req: &AccountCheckTokenRequest{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
				},
			},
			want: &AccountCheckTokenResponse{
				UserAccountToken: &userAccountToken.UserAccountToken{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
					Token:  "token",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userAccountTokenRedisRepository.Find",
			fields: fields{
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					return m
				},
				userAccountRedisRepository: func(ctrl *gomock.Controller) userAccount.UserAccountRedisRepository {
					m := userAccount.NewMockUserAccountRedisRepository(ctrl)
					return m
				},
				userAccountTokenRedisRepository: func(ctrl *gomock.Controller) userAccountToken.UserAccountTokenRedisRepository {
					m := userAccountToken.NewMockUserAccountTokenRedisRepository(ctrl)
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
			},
			args: args{
				ctx: nil,
				tx:  nil,
				rtx: nil,
				req: &AccountCheckTokenRequest{
					UserId: "0:WntR-PyhOJeDiE5jodeR",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userAccountTokenRedisRepository.Find", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &accountService{
				userAccountMysqlRepository:      tt.fields.userAccountMysqlRepository(ctrl),
				userAccountRedisRepository:      tt.fields.userAccountRedisRepository(ctrl),
				userAccountTokenRedisRepository: tt.fields.userAccountTokenRedisRepository(ctrl),
				shardService:                    tt.fields.shardService(ctrl),
			}

			got, err := s.CheckToken(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("CheckToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_GenerateUserId(t *testing.T) {
	type fields struct {
		userAccountMysqlRepository func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository
		shardService               func(ctrl *gomock.Controller) shard.ShardService
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
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					m.EXPECT().
						GetShardKey(
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
				userAccountMysqlRepository: func(ctrl *gomock.Controller) userAccount.UserAccountMysqlRepository {
					m := userAccount.NewMockUserAccountMysqlRepository(ctrl)
					return m
				},
				shardService: func(ctrl *gomock.Controller) shard.ShardService {
					m := shard.NewMockShardService(ctrl)
					m.EXPECT().
						GetShardKey(
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
				userAccountMysqlRepository: tt.fields.userAccountMysqlRepository(ctrl),
				shardService:               tt.fields.shardService(ctrl),
			}

			got, err := s.GenerateUserId(tt.args.ctx)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GenerateUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				got = "0:WntR-PyhOJeDiE5jodeR"
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}
