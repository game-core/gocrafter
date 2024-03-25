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
			name: "正常：プロフィールが作成できる",
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
