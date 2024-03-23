package friend

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	friendServer "github.com/game-core/gocrafter/api/game/presentation/server/friend"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	friendService "github.com/game-core/gocrafter/pkg/domain/model/friend"
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

func TestFriendUsecase_NewFriendUsecase(t *testing.T) {
	type args struct {
		friendService      friendService.FriendService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want FriendUsecase
	}{
		{
			name: "正常",
			args: args{
				friendService:      nil,
				transactionService: nil,
			},
			want: &friendUsecase{
				friendService:      nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFriendUsecase(tt.args.friendService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFriendUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendUsecase_Get(t *testing.T) {
	type fields struct {
		friendService      func(ctrl *gomock.Controller) friendService.FriendService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *friendServer.FriendGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *friendServer.FriendGetResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Get(
							gomock.Any(),
							&friendService.FriendGetRequest{
								UserId: "0:0000",
							},
						).
						Return(
							&friendService.FriendGetResponse{
								UserFriends: userFriend.UserFriends{
									{
										UserId:       "0:0000",
										FriendUserId: "1:1111",
										FriendType:   enum.FriendType_Applying,
									},
									{
										UserId:       "0:0000",
										FriendUserId: "1:2222",
										FriendType:   enum.FriendType_Approved,
									},
									{
										UserId:       "0:0000",
										FriendUserId: "1:3333",
										FriendType:   enum.FriendType_Approved,
									},
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendGetRequest{
					UserId: "0:0000",
				},
			},
			want: &friendServer.FriendGetResponse{
				UserFriends: []*friendServer.UserFriend{
					{
						UserId:       "0:0000",
						FriendUserId: "1:1111",
						FriendType:   friendServer.FriendType_Applying,
					},
					{
						UserId:       "0:0000",
						FriendUserId: "1:2222",
						FriendType:   friendServer.FriendType_Approved,
					},
					{
						UserId:       "0:0000",
						FriendUserId: "1:3333",
						FriendType:   friendServer.FriendType_Approved,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.friendService.Get",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Get(
							gomock.Any(),
							&friendService.FriendGetRequest{
								UserId: "0:0000",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendGetRequest{
					UserId: "0:0000",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.friendService.Get", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &friendUsecase{
				friendService:      tt.fields.friendService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Get(tt.args.ctx, tt.args.req)
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

func TestFriendUsecase_Send(t *testing.T) {
	type fields struct {
		friendService      func(ctrl *gomock.Controller) friendService.FriendService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *friendServer.FriendSendRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *friendServer.FriendSendResponse
		wantErr error
	}{
		{
			name: "正常：送信できる",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Send(
							gomock.Any(),
							gomock.Any(),
							&friendService.FriendSendRequest{
								UserId: "0:0000",
							},
						).
						Return(
							&friendService.FriendSendResponse{
								UserFriend: &userFriend.UserFriend{
									UserId:       "0:0000",
									FriendUserId: "1:1111",
									FriendType:   enum.FriendType_Applying,
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						MultiUserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						).
						Return()
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendSendRequest{
					UserId: "0:0000",
				},
			},
			want: &friendServer.FriendSendResponse{
				UserFriend: &friendServer.UserFriend{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   friendServer.FriendType_Applying,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.MultiUserMysqlBegin",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
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
				req: &friendServer.FriendSendRequest{
					UserId: "0:0000",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.MultiUserMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.friendService.Send",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Send(
							gomock.Any(),
							gomock.Any(),
							&friendService.FriendSendRequest{
								UserId: "0:0000",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						MultiUserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						).
						Return()
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendSendRequest{
					UserId: "0:0000",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.friendService.Send", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &friendUsecase{
				friendService:      tt.fields.friendService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Send(tt.args.ctx, tt.args.req)
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

func TestFriendUsecase_Approve(t *testing.T) {
	type fields struct {
		friendService      func(ctrl *gomock.Controller) friendService.FriendService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *friendServer.FriendApproveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *friendServer.FriendApproveResponse
		wantErr error
	}{
		{
			name: "正常：承認できる",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Approve(
							gomock.Any(),
							gomock.Any(),
							&friendService.FriendApproveRequest{
								UserId: "0:0000",
							},
						).
						Return(
							&friendService.FriendApproveResponse{
								UserFriend: &userFriend.UserFriend{
									UserId:       "0:0000",
									FriendUserId: "1:1111",
									FriendType:   enum.FriendType_Approved,
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						MultiUserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						).
						Return()
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendApproveRequest{
					UserId: "0:0000",
				},
			},
			want: &friendServer.FriendApproveResponse{
				UserFriend: &friendServer.UserFriend{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   friendServer.FriendType_Approved,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.MultiUserMysqlBegin",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
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
				req: &friendServer.FriendApproveRequest{
					UserId: "0:0000",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.MultiUserMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.friendService.Approve",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Approve(
							gomock.Any(),
							gomock.Any(),
							&friendService.FriendApproveRequest{
								UserId: "0:0000",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						MultiUserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						).
						Return()
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendApproveRequest{
					UserId: "0:0000",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.friendService.Approve", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &friendUsecase{
				friendService:      tt.fields.friendService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Approve(tt.args.ctx, tt.args.req)
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

func TestFriendUsecase_Disapprove(t *testing.T) {
	type fields struct {
		friendService      func(ctrl *gomock.Controller) friendService.FriendService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *friendServer.FriendDisapproveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *friendServer.FriendDisapproveResponse
		wantErr error
	}{
		{
			name: "正常：拒否できる",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Disapprove(
							gomock.Any(),
							gomock.Any(),
							&friendService.FriendDisapproveRequest{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
							},
						).
						Return(
							&friendService.FriendDisapproveResponse{
								UserFriend: &userFriend.UserFriend{
									UserId:       "0:0000",
									FriendUserId: "1:1111",
									FriendType:   enum.FriendType_Disapproved,
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						MultiUserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						).
						Return()
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendDisapproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want: &friendServer.FriendDisapproveResponse{
				UserFriend: &friendServer.UserFriend{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   friendServer.FriendType_Disapproved,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.MultiUserMysqlBegin",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
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
				req: &friendServer.FriendDisapproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.MultiUserMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.friendService.Disapprove",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Disapprove(
							gomock.Any(),
							gomock.Any(),
							&friendService.FriendDisapproveRequest{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						MultiUserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						).
						Return()
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendDisapproveRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.friendService.Disapprove", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &friendUsecase{
				friendService:      tt.fields.friendService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Disapprove(tt.args.ctx, tt.args.req)
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

func TestFriendUsecase_Delete(t *testing.T) {
	type fields struct {
		friendService      func(ctrl *gomock.Controller) friendService.FriendService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *friendServer.FriendDeleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *friendServer.FriendDeleteResponse
		wantErr error
	}{
		{
			name: "正常：拒否できる",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&friendService.FriendDeleteRequest{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
							},
						).
						Return(
							&friendService.FriendDeleteResponse{
								UserFriend: &userFriend.UserFriend{
									UserId:       "0:0000",
									FriendUserId: "1:1111",
									FriendType:   enum.FriendType_NotFriend,
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						MultiUserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						).
						Return()
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendDeleteRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want: &friendServer.FriendDeleteResponse{
				UserFriend: &friendServer.UserFriend{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
					FriendType:   friendServer.FriendType_NotFriend,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.MultiUserMysqlBegin",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
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
				req: &friendServer.FriendDeleteRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.MultiUserMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.friendService.Delete",
			fields: fields{
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&friendService.FriendDeleteRequest{
								UserId:       "0:0000",
								FriendUserId: "1:1111",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						MultiUserMysqlBegin(
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						MultiUserMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						).
						Return()
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &friendServer.FriendDeleteRequest{
					UserId:       "0:0000",
					FriendUserId: "1:1111",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.friendService.Delete", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &friendUsecase{
				friendService:      tt.fields.friendService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Delete(tt.args.ctx, tt.args.req)
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
