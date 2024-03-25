package room

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	roomServer "github.com/game-core/gocrafter/api/game/presentation/server/room"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	roomService "github.com/game-core/gocrafter/pkg/domain/model/room"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
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

func TestRoomUsecase_Search(t *testing.T) {
	type fields struct {
		roomService        func(ctrl *gomock.Controller) roomService.RoomService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *roomServer.RoomSearchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *roomServer.RoomSearchResponse
		wantErr error
	}{
		{
			name: "正常：検索できる場合",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Search(
							gomock.Any(),
							&roomService.RoomSearchRequest{
								UserId: "0:test",
								Name:   "test_user_room",
							},
						).
						Return(
							&roomService.RoomSearchResponse{
								CommonRooms: commonRoom.CommonRooms{
									{
										RoomId:          "roomId",
										HostUserId:      "0:test",
										RoomReleaseType: enum.RoomReleaseType_Public,
										Name:            "test_user_room",
										UserCount:       1,
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
				req: &roomServer.RoomSearchRequest{
					UserId: "0:test",
					Name:   "test_user_room",
				},
			},
			want: &roomServer.RoomSearchResponse{
				CommonRooms: []*roomServer.CommonRoom{
					{
						RoomId:          "roomId",
						HostUserId:      "0:test",
						RoomReleaseType: roomServer.RoomReleaseType_Public,
						Name:            "test_user_room",
						UserCount:       1,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.roomService.Search",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Search(
							gomock.Any(),
							&roomService.RoomSearchRequest{
								UserId: "0:test",
								Name:   "test_user_room",
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
				req: &roomServer.RoomSearchRequest{
					UserId: "0:test",
					Name:   "test_user_room",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.roomService.Search", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &roomUsecase{
				roomService:        tt.fields.roomService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Search(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomUsecase_Create(t *testing.T) {
	type fields struct {
		roomService        func(ctrl *gomock.Controller) roomService.RoomService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *roomServer.RoomCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *roomServer.RoomCreateResponse
		wantErr error
	}{
		{
			name: "正常：作成できる場合",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							&roomService.RoomCreateRequest{
								UserId:          "0:test",
								Name:            "test_user_room",
								RoomReleaseType: enum.RoomReleaseType_Public,
							},
						).
						Return(
							&roomService.RoomCreateResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "roomId",
									HostUserId:      "0:test",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "test_user_room",
									UserCount:       1,
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						CommonMysqlBegin(
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommonMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &roomServer.RoomCreateRequest{
					UserId:          "0:test",
					Name:            "test_user_room",
					RoomReleaseType: roomServer.RoomReleaseType_Public,
				},
			},
			want: &roomServer.RoomCreateResponse{
				CommonRoom: &roomServer.CommonRoom{
					RoomId:          "roomId",
					HostUserId:      "0:test",
					RoomReleaseType: roomServer.RoomReleaseType_Public,
					Name:            "test_user_room",
					UserCount:       1,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.CommonMysqlBegin",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						CommonMysqlBegin(
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
				req: &roomServer.RoomCreateRequest{
					UserId:          "0:test",
					Name:            "test_user_room",
					RoomReleaseType: roomServer.RoomReleaseType_Public,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.CommonMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.roomService.Create",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							&roomService.RoomCreateRequest{
								UserId:          "0:test",
								Name:            "test_user_room",
								RoomReleaseType: enum.RoomReleaseType_Public,
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
						CommonMysqlBegin(
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommonMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &roomServer.RoomCreateRequest{
					UserId:          "0:test",
					Name:            "test_user_room",
					RoomReleaseType: roomServer.RoomReleaseType_Public,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.roomService.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &roomUsecase{
				roomService:        tt.fields.roomService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Create(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomUsecase_Delete(t *testing.T) {
	type fields struct {
		roomService        func(ctrl *gomock.Controller) roomService.RoomService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *roomServer.RoomDeleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *roomServer.RoomDeleteResponse
		wantErr error
	}{
		{
			name: "正常：削除できる場合",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&roomService.RoomDeleteRequest{
								UserId: "0:test",
								RoomId: "roomId",
							},
						).
						Return(
							&roomService.RoomDeleteResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "roomId",
									HostUserId:      "0:test",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "test_user_room",
									UserCount:       1,
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						CommonMysqlBegin(
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommonMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &roomServer.RoomDeleteRequest{
					UserId: "0:test",
					RoomId: "roomId",
				},
			},
			want: &roomServer.RoomDeleteResponse{
				CommonRoom: &roomServer.CommonRoom{
					RoomId:          "roomId",
					HostUserId:      "0:test",
					RoomReleaseType: roomServer.RoomReleaseType_Public,
					Name:            "test_user_room",
					UserCount:       1,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.CommonMysqlBegin",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						CommonMysqlBegin(
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
				req: &roomServer.RoomDeleteRequest{
					UserId: "0:test",
					RoomId: "roomId",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.CommonMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.roomService.Delete",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&roomService.RoomDeleteRequest{
								UserId: "0:test",
								RoomId: "roomId",
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
						CommonMysqlBegin(
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommonMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &roomServer.RoomDeleteRequest{
					UserId: "0:test",
					RoomId: "roomId",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.roomService.Delete", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &roomUsecase{
				roomService:        tt.fields.roomService(ctrl),
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

func TestRoomUsecase_CheckIn(t *testing.T) {
	type fields struct {
		roomService        func(ctrl *gomock.Controller) roomService.RoomService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *roomServer.RoomCheckInRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *roomServer.RoomCheckInResponse
		wantErr error
	}{
		{
			name: "正常：入室できる場合",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						CheckIn(
							gomock.Any(),
							gomock.Any(),
							&roomService.RoomCheckInRequest{
								UserId: "1:test",
								RoomId: "roomId",
							},
						).
						Return(
							&roomService.RoomCheckInResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "roomId",
									HostUserId:      "0:test",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "test_user_room",
									UserCount:       2,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "roomId",
									UserId:               "1:test",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						CommonMysqlBegin(
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommonMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &roomServer.RoomCheckInRequest{
					UserId: "1:test",
					RoomId: "roomId",
				},
			},
			want: &roomServer.RoomCheckInResponse{
				CommonRoom: &roomServer.CommonRoom{
					RoomId:          "roomId",
					HostUserId:      "0:test",
					RoomReleaseType: roomServer.RoomReleaseType_Public,
					Name:            "test_user_room",
					UserCount:       2,
				},
				CommonRoomUser: &roomServer.CommonRoomUser{
					RoomId:               "roomId",
					UserId:               "1:test",
					RoomUserPositionType: roomServer.RoomUserPositionType_General,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.CommonMysqlBegin",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						CommonMysqlBegin(
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
				req: &roomServer.RoomCheckInRequest{
					UserId: "0:test",
					RoomId: "roomId",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.CommonMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.roomService.CheckIn",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						CheckIn(
							gomock.Any(),
							gomock.Any(),
							&roomService.RoomCheckInRequest{
								UserId: "0:test",
								RoomId: "roomId",
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
						CommonMysqlBegin(
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommonMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &roomServer.RoomCheckInRequest{
					UserId: "0:test",
					RoomId: "roomId",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.roomService.CheckIn", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &roomUsecase{
				roomService:        tt.fields.roomService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.CheckIn(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("CheckIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomUsecase_CheckOut(t *testing.T) {
	type fields struct {
		roomService        func(ctrl *gomock.Controller) roomService.RoomService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *roomServer.RoomCheckOutRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *roomServer.RoomCheckOutResponse
		wantErr error
	}{
		{
			name: "正常：退出できる場合",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						CheckOut(
							gomock.Any(),
							gomock.Any(),
							&roomService.RoomCheckOutRequest{
								UserId: "1:test",
								RoomId: "roomId",
							},
						).
						Return(
							&roomService.RoomCheckOutResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "roomId",
									HostUserId:      "0:test",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "test_user_room",
									UserCount:       2,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "roomId",
									UserId:               "1:test",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						CommonMysqlBegin(
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommonMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &roomServer.RoomCheckOutRequest{
					UserId: "1:test",
					RoomId: "roomId",
				},
			},
			want: &roomServer.RoomCheckOutResponse{
				CommonRoom: &roomServer.CommonRoom{
					RoomId:          "roomId",
					HostUserId:      "0:test",
					RoomReleaseType: roomServer.RoomReleaseType_Public,
					Name:            "test_user_room",
					UserCount:       2,
				},
				CommonRoomUser: &roomServer.CommonRoomUser{
					RoomId:               "roomId",
					UserId:               "1:test",
					RoomUserPositionType: roomServer.RoomUserPositionType_General,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.CommonMysqlBegin",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						CommonMysqlBegin(
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
				req: &roomServer.RoomCheckOutRequest{
					UserId: "0:test",
					RoomId: "roomId",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.CommonMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.roomService.CheckOut",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						CheckOut(
							gomock.Any(),
							gomock.Any(),
							&roomService.RoomCheckOutRequest{
								UserId: "0:test",
								RoomId: "roomId",
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
						CommonMysqlBegin(
							gomock.Any(),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						CommonMysqlEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &roomServer.RoomCheckOutRequest{
					UserId: "0:test",
					RoomId: "roomId",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.roomService.CheckOut", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &roomUsecase{
				roomService:        tt.fields.roomService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.CheckOut(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("CheckOut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckOut() = %v, want %v", got, tt.want)
			}
		})
	}
}
