package friend

import (
	"context"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/account"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
)

func TestNewFriendService_NewFriendService(t *testing.T) {
	type args struct {
		accountService            account.AccountService
		userFriendMysqlRepository userFriend.UserFriendMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want FriendService
	}{
		{
			name: "正常",
			args: args{
				accountService:            nil,
				userFriendMysqlRepository: nil,
			},
			want: &friendService{
				userFriendMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFriendService(
				tt.args.accountService,
				tt.args.userFriendMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFriendService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendService_Get(t *testing.T) {
	type fields struct {
		userFriendMysqlRepository func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *FriendGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FriendGetResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
							"0:0000",
						).
						Return(
							userFriend.UserFriends{
								{
									UserId:       "0:0000",
									FriendUserId: "1:1111",
									FriendType:   enum.FriendType_Approved,
								},
								{
									UserId:       "0:0000",
									FriendUserId: "1:2222",
									FriendType:   enum.FriendType_Applying,
								},
								{
									UserId:       "0:0000",
									FriendUserId: "1:3333",
									FriendType:   enum.FriendType_Approved,
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendGetRequest{
					UserId: "0:0000",
				},
			},
			want: &FriendGetResponse{
				userFriend.UserFriends{
					{
						UserId:       "0:0000",
						FriendUserId: "1:1111",
						FriendType:   enum.FriendType_Approved,
					},
					{
						UserId:       "0:0000",
						FriendUserId: "1:2222",
						FriendType:   enum.FriendType_Applying,
					},
					{
						UserId:       "0:0000",
						FriendUserId: "1:3333",
						FriendType:   enum.FriendType_Approved,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userFriendMysqlRepository.FindList",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
							"0:0000",
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
				req: &FriendGetRequest{
					UserId: "0:0000",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.FindList", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &friendService{
				userFriendMysqlRepository: tt.fields.userFriendMysqlRepository(ctrl),
			}

			got, err := s.Get(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendService_Send(t *testing.T) {
	type fields struct {
		accountService            func(ctrl *gomock.Controller) account.AccountService
		userFriendMysqlRepository func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository
	}
	type args struct {
		ctx context.Context
		txs map[string]*gorm.DB
		req *FriendSendRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FriendSendResponse
		wantErr error
	}{
		{
			name: "正常：送信できる",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) account.AccountService {
					m := account.NewMockAccountService(ctrl)
					m.EXPECT().
						Check(
							nil,
							account.SetAccountCheckRequest(
								"1:1111",
							),
						).
						Return(
							&account.AccountCheckResponse{
								UserAccount: &userAccount.UserAccount{
									UserId:   "1:1111",
									Name:     "テストアカウント",
									Password: "パスワード",
									LoginAt:  time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
									LogoutAt: time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_NotApproved,
							},
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_NotApproved,
							},
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Applying,
							},
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Applying,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendSendRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want: &FriendSendResponse{
				UserFriend: &userFriend.UserFriend{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   enum.FriendType_Applying,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.accountService.Check",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) account.AccountService {
					m := account.NewMockAccountService(ctrl)
					m.EXPECT().
						Check(
							nil,
							account.SetAccountCheckRequest(
								"1:1111",
							),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendSendRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.accountService.Check", errors.NewTestError()),
		},
		{
			name: "異常：s.userFriendMysqlRepository.FindOrNil",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) account.AccountService {
					m := account.NewMockAccountService(ctrl)
					m.EXPECT().
						Check(
							nil,
							account.SetAccountCheckRequest(
								"1:1111",
							),
						).
						Return(
							&account.AccountCheckResponse{
								UserAccount: &userAccount.UserAccount{
									UserId:   "1:1111",
									Name:     "テストアカウント",
									Password: "パスワード",
									LoginAt:  time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
									LogoutAt: time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
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
				req: &FriendSendRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：already applied",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) account.AccountService {
					m := account.NewMockAccountService(ctrl)
					m.EXPECT().
						Check(
							nil,
							account.SetAccountCheckRequest(
								"1:1111",
							),
						).
						Return(
							&account.AccountCheckResponse{
								UserAccount: &userAccount.UserAccount{
									UserId:   "1:1111",
									Name:     "テストアカウント",
									Password: "パスワード",
									LoginAt:  time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
									LogoutAt: time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_NotApproved,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendSendRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewError("already applied"),
		},
		{
			name: "異常：s.userFriendMysqlRepository.Create",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) account.AccountService {
					m := account.NewMockAccountService(ctrl)
					m.EXPECT().
						Check(
							nil,
							account.SetAccountCheckRequest(
								"1:1111",
							),
						).
						Return(
							&account.AccountCheckResponse{
								UserAccount: &userAccount.UserAccount{
									UserId:   "1:1111",
									Name:     "テストアカウント",
									Password: "パスワード",
									LoginAt:  time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
									LogoutAt: time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_NotApproved,
							},
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
				req: &FriendSendRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：s.userFriendMysqlRepository.Create",
			fields: fields{
				accountService: func(ctrl *gomock.Controller) account.AccountService {
					m := account.NewMockAccountService(ctrl)
					m.EXPECT().
						Check(
							nil,
							account.SetAccountCheckRequest(
								"1:1111",
							),
						).
						Return(
							&account.AccountCheckResponse{
								UserAccount: &userAccount.UserAccount{
									UserId:   "1:1111",
									Name:     "テストアカウント",
									Password: "パスワード",
									LoginAt:  time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
									LogoutAt: time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_NotApproved,
							},
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Applying,
							},
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Applying,
							},
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
				req: &FriendSendRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &friendService{
				accountService:            tt.fields.accountService(ctrl),
				userFriendMysqlRepository: tt.fields.userFriendMysqlRepository(ctrl),
			}

			got, err := s.Send(tt.args.ctx, tt.args.txs, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendService_Approve(t *testing.T) {
	type fields struct {
		userFriendMysqlRepository func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository
	}
	type args struct {
		ctx context.Context
		txs map[string]*gorm.DB
		req *FriendApproveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FriendApproveResponse
		wantErr error
	}{
		{
			name: "正常：承認できる",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_NotApproved,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_Approved,
							},
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_Approved,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Approved,
							},
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Approved,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendApproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want: &FriendApproveResponse{
				UserFriend: &userFriend.UserFriend{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   enum.FriendType_Approved,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userFriendMysqlRepository.FindOrNil",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
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
				req: &FriendApproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：not applied",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							nil,
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendApproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewError("not applied"),
		},
		{
			name: "異常：s.userFriendMysqlRepository.Update",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_NotApproved,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_Approved,
							},
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
				req: &FriendApproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.userFriendMysqlRepository.Update",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_NotApproved,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_Approved,
							},
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_Approved,
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Approved,
							},
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
				req: &FriendApproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.Update", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &friendService{
				userFriendMysqlRepository: tt.fields.userFriendMysqlRepository(ctrl),
			}

			got, err := s.Approve(tt.args.ctx, tt.args.txs, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Approve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Approve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendService_Disapprove(t *testing.T) {
	type fields struct {
		userFriendMysqlRepository func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository
	}
	type args struct {
		ctx context.Context
		txs map[string]*gorm.DB
		req *FriendDisapproveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FriendDisapproveResponse
		wantErr error
	}{
		{
			name: "正常：拒否できる",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_NotApproved,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_NotApproved,
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_NotApproved,
							},
						).
						Return(
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendDisapproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want: &FriendDisapproveResponse{
				UserFriend: &userFriend.UserFriend{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   enum.FriendType_Disapproved,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userFriendMysqlRepository.FindOrNil",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
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
				req: &FriendDisapproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：not applied",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							nil,
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendDisapproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewError("not applied"),
		},
		{
			name: "異常：s.userFriendMysqlRepository.Delete",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_NotApproved,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_NotApproved,
							},
						).
						Return(
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendDisapproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.Delete", errors.NewTestError()),
		},
		{
			name: "異常：s.userFriendMysqlRepository.Delete",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_NotApproved,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_NotApproved,
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_NotApproved,
							},
						).
						Return(
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendDisapproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.Delete", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &friendService{
				userFriendMysqlRepository: tt.fields.userFriendMysqlRepository(ctrl),
			}

			got, err := s.Disapprove(tt.args.ctx, tt.args.txs, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Disapprove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Disapprove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendService_Delete(t *testing.T) {
	type fields struct {
		userFriendMysqlRepository func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository
	}
	type args struct {
		ctx context.Context
		txs map[string]*gorm.DB
		req *FriendDeleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FriendDeleteResponse
		wantErr error
	}{
		{
			name: "正常：削除できる",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Approved,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_Approved,
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Approved,
							},
						).
						Return(
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendDeleteRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want: &FriendDeleteResponse{
				UserFriend: &userFriend.UserFriend{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   enum.FriendType_NotFriend,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userFriendMysqlRepository.FindOrNil",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
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
				req: &FriendDeleteRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：not friend",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							nil,
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendDeleteRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewError("not friend"),
		},
		{
			name: "異常：s.userFriendMysqlRepository.Delete",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Approved,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_Approved,
							},
						).
						Return(
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendDeleteRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.Delete", errors.NewTestError()),
		},
		{
			name: "異常：s.userFriendMysqlRepository.Delete",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Approved,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "1:1111",
								FriendUserId: "0:0000",
								FriendType:   enum.FriendType_Approved,
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Approved,
							},
						).
						Return(
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendDeleteRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.Delete", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &friendService{
				userFriendMysqlRepository: tt.fields.userFriendMysqlRepository(ctrl),
			}

			got, err := s.Delete(tt.args.ctx, tt.args.txs, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendService_Check(t *testing.T) {
	type fields struct {
		userFriendMysqlRepository func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *FriendCheckRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FriendCheckResponse
		wantErr error
	}{
		{
			name: "正常：確認できる",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Approved,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendCheckRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want: &FriendCheckResponse{
				UserFriend: &userFriend.UserFriend{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   enum.FriendType_Approved,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userFriendMysqlRepository.FindOrNil",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:0000",
							"1:1111",
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
				req: &FriendCheckRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userFriendMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：not a friend",
			fields: fields{
				userFriendMysqlRepository: func(ctrl *gomock.Controller) userFriend.UserFriendMysqlRepository {
					m := userFriend.NewMockUserFriendMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:0000",
							"1:1111",
						).
						Return(
							&userFriend.UserFriend{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
								FriendType:   enum.FriendType_Applying,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &FriendCheckRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewError("not a friend"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &friendService{
				userFriendMysqlRepository: tt.fields.userFriendMysqlRepository(ctrl),
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
