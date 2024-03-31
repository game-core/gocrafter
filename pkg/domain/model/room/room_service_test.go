package room

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	configService "github.com/game-core/gocrafter/pkg/domain/model/config"
	friendService "github.com/game-core/gocrafter/pkg/domain/model/friend"
	"github.com/game-core/gocrafter/pkg/domain/model/friend/userFriend"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
)

func TestNewRoomService_NewRoomService(t *testing.T) {
	type args struct {
		configService                 configService.ConfigService
		friendService                 friendService.FriendService
		commonRoomMysqlRepository     commonRoom.CommonRoomMysqlRepository
		commonRoomUserMysqlRepository commonRoomUser.CommonRoomUserMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want RoomService
	}{
		{
			name: "正常",
			args: args{
				configService:                 nil,
				friendService:                 nil,
				commonRoomMysqlRepository:     nil,
				commonRoomUserMysqlRepository: nil,
			},
			want: &roomService{
				configService:                 nil,
				friendService:                 nil,
				commonRoomMysqlRepository:     nil,
				commonRoomUserMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRoomService(
				tt.args.configService,
				tt.args.friendService,
				tt.args.commonRoomMysqlRepository,
				tt.args.commonRoomUserMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoomService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomService_Search(t *testing.T) {
	type fields struct {
		configService                 func(ctrl *gomock.Controller) configService.ConfigService
		friendService                 func(ctrl *gomock.Controller) friendService.FriendService
		commonRoomMysqlRepository     func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository
		commonRoomUserMysqlRepository func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *RoomSearchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RoomSearchResponse
		wantErr error
	}{
		{
			name: "正常：検索できる場合",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByName(
							gomock.Any(),
							"テストルーム",
						).
						Return(
							commonRoom.CommonRooms{
								{
									RoomId:          "roomId",
									HostUserId:      "0:test",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
							},
							nil,
						)

					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomSearchRequest{
					UserId: "0:test",
					Name:   "テストルーム",
				},
			},
			want: &RoomSearchResponse{
				CommonRooms: commonRoom.CommonRooms{
					{
						RoomId:          "roomId",
						HostUserId:      "0:test",
						RoomReleaseType: enum.RoomReleaseType_Public,
						Name:            "テストルーム",
						UserCount:       1,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.commonRoomMysqlRepository.FindListByName",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByName(
							gomock.Any(),
							"テストルーム",
						).
						Return(
							nil,
							errors.NewTestError(),
						)

					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomSearchRequest{
					UserId: "0:test",
					Name:   "テストルーム",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.FindListByName", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &roomService{
				configService:                 tt.fields.configService(ctrl),
				friendService:                 tt.fields.friendService(ctrl),
				commonRoomMysqlRepository:     tt.fields.commonRoomMysqlRepository(ctrl),
				commonRoomUserMysqlRepository: tt.fields.commonRoomUserMysqlRepository(ctrl),
			}

			got, err := s.Search(tt.args.ctx, tt.args.req)
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

func TestRoomService_Create(t *testing.T) {
	type fields struct {
		configService                 func(ctrl *gomock.Controller) configService.ConfigService
		friendService                 func(ctrl *gomock.Controller) friendService.FriendService
		commonRoomMysqlRepository     func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository
		commonRoomUserMysqlRepository func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *RoomCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RoomCreateResponse
		wantErr error
	}{
		{
			name: "正常：作成できる場合",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					m.EXPECT().
						GetInt32(
							gomock.Any(),
							enum.ConfigType_MaxRoomNumber,
						).
						Return(
							int32(3),
							nil,
						)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByHostUserId(
							gomock.Any(),
							"0:test",
						).
						Return(
							commonRoom.CommonRooms{},
							nil,
						)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "roomId",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							&commonRoomUser.CommonRoomUser{
								RoomId:               "roomId",
								UserId:               "0:test",
								RoomUserPositionType: enum.RoomUserPositionType_Leader,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCreateRequest{
					UserId:          "0:test",
					Name:            "テストルーム",
					RoomReleaseType: enum.RoomReleaseType_Public,
				},
			},
			want: &RoomCreateResponse{
				CommonRoom: &commonRoom.CommonRoom{
					RoomId:          "roomId",
					HostUserId:      "0:test",
					RoomReleaseType: enum.RoomReleaseType_Public,
					Name:            "テストルーム",
					UserCount:       1,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.generateRoomId: failed to s.commonRoomMysqlRepository.FindListByHostUserId",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByHostUserId(
							gomock.Any(),
							"0:test",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCreateRequest{
					UserId:          "0:test",
					Name:            "テストルーム",
					RoomReleaseType: enum.RoomReleaseType_Public,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.generateRoomId: failed to s.commonRoomMysqlRepository.FindListByHostUserId", errors.NewTestError()),
		},
		{
			name: "異常：s.generateRoomId: failed to s.configService.GetInt32",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					m.EXPECT().
						GetInt32(
							gomock.Any(),
							enum.ConfigType_MaxRoomNumber,
						).
						Return(
							int32(0),
							errors.NewTestError(),
						)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByHostUserId(
							gomock.Any(),
							"0:test",
						).
						Return(
							commonRoom.CommonRooms{},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCreateRequest{
					UserId:          "0:test",
					Name:            "テストルーム",
					RoomReleaseType: enum.RoomReleaseType_Public,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.generateRoomId: failed to s.configService.GetInt32", errors.NewTestError()),
		},
		{
			name: "異常：s.generateRoomId: room number exceeded",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					m.EXPECT().
						GetInt32(
							gomock.Any(),
							enum.ConfigType_MaxRoomNumber,
						).
						Return(
							int32(2),
							nil,
						)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByHostUserId(
							gomock.Any(),
							"0:test",
						).
						Return(
							commonRoom.CommonRooms{
								{
									RoomId:          "roomId1",
									HostUserId:      "0:test",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム1",
									UserCount:       1,
								},
								{
									RoomId:          "roomId2",
									HostUserId:      "0:test",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム2",
									UserCount:       1,
								},
								{
									RoomId:          "roomId3",
									HostUserId:      "0:test",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム3",
									UserCount:       1,
								},
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCreateRequest{
					UserId:          "0:test",
					Name:            "テストルーム",
					RoomReleaseType: enum.RoomReleaseType_Public,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.generateRoomId: room number exceeded"),
		},
		{
			name: "異常：s.commonRoomMysqlRepository.Create",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					m.EXPECT().
						GetInt32(
							gomock.Any(),
							enum.ConfigType_MaxRoomNumber,
						).
						Return(
							int32(3),
							nil,
						)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByHostUserId(
							gomock.Any(),
							"0:test",
						).
						Return(
							commonRoom.CommonRooms{},
							nil,
						)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCreateRequest{
					UserId:          "0:test",
					Name:            "テストルーム",
					RoomReleaseType: enum.RoomReleaseType_Public,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：s.commonRoomUserMysqlRepository.Create",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					m.EXPECT().
						GetInt32(
							gomock.Any(),
							enum.ConfigType_MaxRoomNumber,
						).
						Return(
							int32(3),
							nil,
						)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByHostUserId(
							gomock.Any(),
							"0:test",
						).
						Return(
							commonRoom.CommonRooms{},
							nil,
						)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "roomId",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
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
				req: &RoomCreateRequest{
					UserId:          "0:test",
					Name:            "テストルーム",
					RoomReleaseType: enum.RoomReleaseType_Public,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomUserMysqlRepository.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &roomService{
				configService:                 tt.fields.configService(ctrl),
				friendService:                 tt.fields.friendService(ctrl),
				commonRoomMysqlRepository:     tt.fields.commonRoomMysqlRepository(ctrl),
				commonRoomUserMysqlRepository: tt.fields.commonRoomUserMysqlRepository(ctrl),
			}

			got, err := s.Create(tt.args.ctx, tt.args.tx, tt.args.req)
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

func TestRoomService_Delete(t *testing.T) {
	type fields struct {
		configService                 func(ctrl *gomock.Controller) configService.ConfigService
		friendService                 func(ctrl *gomock.Controller) friendService.FriendService
		commonRoomMysqlRepository     func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository
		commonRoomUserMysqlRepository func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *RoomDeleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RoomDeleteResponse
		wantErr error
	}{
		{
			name: "正常：削除できる場合",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindByRoomIdAndHostUserId(
							gomock.Any(),
							"room1",
							"0:test",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム1",
								UserCount:       1,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム1",
								UserCount:       1,
							},
						).
						Return(
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						FindListByRoomId(
							gomock.Any(),
							"room1",
						).
						Return(
							commonRoomUser.CommonRoomUsers{
								{
									RoomId:               "room1",
									UserId:               "0:test",
									RoomUserPositionType: enum.RoomUserPositionType_Leader,
								},
								{
									RoomId:               "room1",
									UserId:               "1:test",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "0:test",
								RoomUserPositionType: enum.RoomUserPositionType_Leader,
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "1:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
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
				req: &RoomDeleteRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want: &RoomDeleteResponse{
				CommonRoom: &commonRoom.CommonRoom{
					RoomId:          "room1",
					HostUserId:      "0:test",
					RoomReleaseType: enum.RoomReleaseType_Public,
					Name:            "テストルーム1",
					UserCount:       1,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.commonRoomMysqlRepository.FindByRoomIdAndHostUserId",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindByRoomIdAndHostUserId(
							gomock.Any(),
							"room1",
							"0:test",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomDeleteRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.FindByRoomIdAndHostUserId", errors.NewTestError()),
		},
		{
			name: "異常：s.commonRoomMysqlRepository.Delete",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindByRoomIdAndHostUserId(
							gomock.Any(),
							"room1",
							"0:test",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム1",
								UserCount:       1,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム1",
								UserCount:       1,
							},
						).
						Return(
							errors.NewTestError(),
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomDeleteRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.Delete", errors.NewTestError()),
		},
		{
			name: "異常：s.commonRoomUserMysqlRepository.FindListByRoomId",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindByRoomIdAndHostUserId(
							gomock.Any(),
							"room1",
							"0:test",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム1",
								UserCount:       1,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム1",
								UserCount:       1,
							},
						).
						Return(
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						FindListByRoomId(
							gomock.Any(),
							"room1",
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
				req: &RoomDeleteRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomUserMysqlRepository.FindListByRoomId", errors.NewTestError()),
		},
		{
			name: "異常：s.commonRoomUserMysqlRepository.Delete",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindByRoomIdAndHostUserId(
							gomock.Any(),
							"room1",
							"0:test",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム1",
								UserCount:       1,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム1",
								UserCount:       1,
							},
						).
						Return(
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						FindListByRoomId(
							gomock.Any(),
							"room1",
						).
						Return(
							commonRoomUser.CommonRoomUsers{
								{
									RoomId:               "room1",
									UserId:               "0:test",
									RoomUserPositionType: enum.RoomUserPositionType_Leader,
								},
								{
									RoomId:               "room1",
									UserId:               "1:test",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "0:test",
								RoomUserPositionType: enum.RoomUserPositionType_Leader,
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "1:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
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
				req: &RoomDeleteRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomUserMysqlRepository.Delete", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &roomService{
				configService:                 tt.fields.configService(ctrl),
				friendService:                 tt.fields.friendService(ctrl),
				commonRoomMysqlRepository:     tt.fields.commonRoomMysqlRepository(ctrl),
				commonRoomUserMysqlRepository: tt.fields.commonRoomUserMysqlRepository(ctrl),
			}

			got, err := s.Delete(tt.args.ctx, tt.args.tx, tt.args.req)
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

func TestRoomService_Check(t *testing.T) {
	type fields struct {
		configService                 func(ctrl *gomock.Controller) configService.ConfigService
		friendService                 func(ctrl *gomock.Controller) friendService.FriendService
		commonRoomMysqlRepository     func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository
		commonRoomUserMysqlRepository func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *RoomCheckRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RoomCheckResponse
		wantErr error
	}{
		{
			name: "正常：確認できる場合",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
							"0:test",
						).
						Return(
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "0:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want: &RoomCheckResponse{
				CommonRoom: &commonRoom.CommonRoom{
					RoomId:          "room1",
					HostUserId:      "0:test",
					RoomReleaseType: enum.RoomReleaseType_Public,
					Name:            "テストルーム",
					UserCount:       1,
				},
				CommonRoomUser: &commonRoomUser.CommonRoomUser{
					RoomId:               "room1",
					UserId:               "0:test",
					RoomUserPositionType: enum.RoomUserPositionType_General,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.commonRoomMysqlRepository.Find",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.commonRoomUserMysqlRepository.Find",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
							"0:test",
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
				req: &RoomCheckRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomUserMysqlRepository.Find", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &roomService{
				configService:                 tt.fields.configService(ctrl),
				friendService:                 tt.fields.friendService(ctrl),
				commonRoomMysqlRepository:     tt.fields.commonRoomMysqlRepository(ctrl),
				commonRoomUserMysqlRepository: tt.fields.commonRoomUserMysqlRepository(ctrl),
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

func TestRoomService_CheckIn(t *testing.T) {
	type fields struct {
		configService                 func(ctrl *gomock.Controller) configService.ConfigService
		friendService                 func(ctrl *gomock.Controller) friendService.FriendService
		commonRoomMysqlRepository     func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository
		commonRoomUserMysqlRepository func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *RoomCheckInRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RoomCheckInResponse
		wantErr error
	}{
		{
			name: "正常：入室できる場合（Publicルーム）",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       2,
							},
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       2,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "0:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
							},
						).
						Return(
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "0:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckInRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want: &RoomCheckInResponse{
				CommonRoom: &commonRoom.CommonRoom{
					RoomId:          "room1",
					HostUserId:      "0:test",
					RoomReleaseType: enum.RoomReleaseType_Public,
					Name:            "テストルーム",
					UserCount:       2,
				},
				CommonRoomUser: &commonRoomUser.CommonRoomUser{
					RoomId:               "room1",
					UserId:               "0:test",
					RoomUserPositionType: enum.RoomUserPositionType_General,
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：入室できる場合（Privateルーム）",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Check(
							gomock.Any(),
							&friendService.FriendCheckRequest{
								UserId:       "0:test",
								FriendUserId: "1:test",
							},
						).
						Return(
							&friendService.FriendCheckResponse{
								UserFriend: &userFriend.UserFriend{
									UserId:       "0:test",
									FriendUserId: "1:test",
									FriendType:   enum.FriendType_Approved,
								},
							},
							nil,
						)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       2,
							},
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       2,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "1:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
							},
						).
						Return(
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "1:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckInRequest{
					UserId: "1:test",
					RoomId: "room1",
				},
			},
			want: &RoomCheckInResponse{
				CommonRoom: &commonRoom.CommonRoom{
					RoomId:          "room1",
					HostUserId:      "0:test",
					RoomReleaseType: enum.RoomReleaseType_Private,
					Name:            "テストルーム",
					UserCount:       2,
				},
				CommonRoomUser: &commonRoomUser.CommonRoomUser{
					RoomId:               "room1",
					UserId:               "1:test",
					RoomUserPositionType: enum.RoomUserPositionType_General,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.commonRoomMysqlRepository.Find",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckInRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.commonRoomMysqlRepository.Update",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       2,
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckInRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.commonRoomMysqlRepository.Create",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       2,
							},
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Public,
								Name:            "テストルーム",
								UserCount:       2,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "0:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
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
				req: &RoomCheckInRequest{
					UserId: "0:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：s.friendService.Check",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					m.EXPECT().
						Check(
							gomock.Any(),
							&friendService.FriendCheckRequest{
								UserId:       "0:test",
								FriendUserId: "1:test",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckInRequest{
					UserId: "1:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.friendService.Check", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &roomService{
				configService:                 tt.fields.configService(ctrl),
				friendService:                 tt.fields.friendService(ctrl),
				commonRoomMysqlRepository:     tt.fields.commonRoomMysqlRepository(ctrl),
				commonRoomUserMysqlRepository: tt.fields.commonRoomUserMysqlRepository(ctrl),
			}

			got, err := s.CheckIn(tt.args.ctx, tt.args.tx, tt.args.req)
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

func TestRoomService_CheckOut(t *testing.T) {
	type fields struct {
		configService                 func(ctrl *gomock.Controller) configService.ConfigService
		friendService                 func(ctrl *gomock.Controller) friendService.FriendService
		commonRoomMysqlRepository     func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository
		commonRoomUserMysqlRepository func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *RoomCheckOutRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RoomCheckOutResponse
		wantErr error
	}{
		{
			name: "正常：退出できる場合",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       2,
							},
							nil,
						)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       1,
							},
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
							"1:test",
						).
						Return(
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "1:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "1:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
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
				req: &RoomCheckOutRequest{
					UserId: "1:test",
					RoomId: "room1",
				},
			},
			want: &RoomCheckOutResponse{
				CommonRoom: &commonRoom.CommonRoom{
					RoomId:          "room1",
					HostUserId:      "0:test",
					RoomReleaseType: enum.RoomReleaseType_Private,
					Name:            "テストルーム",
					UserCount:       1,
				},
				CommonRoomUser: &commonRoomUser.CommonRoomUser{
					RoomId:               "room1",
					UserId:               "1:test",
					RoomUserPositionType: enum.RoomUserPositionType_General,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.commonRoomMysqlRepository.Find",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckOutRequest{
					UserId: "1:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：host user cannot check out",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "1:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       2,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckOutRequest{
					UserId: "1:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewError("host user cannot check out"),
		},
		{
			name: "異常：s.commonRoomUserMysqlRepository.Find",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       2,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
							"1:test",
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
				req: &RoomCheckOutRequest{
					UserId: "1:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomUserMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.commonRoomMysqlRepository.Update",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       2,
							},
							nil,
						)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       1,
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
							"1:test",
						).
						Return(
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "1:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RoomCheckOutRequest{
					UserId: "1:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomMysqlRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.commonRoomUserMysqlRepository.Delete",
			fields: fields{
				configService: func(ctrl *gomock.Controller) configService.ConfigService {
					m := configService.NewMockConfigService(ctrl)
					return m
				},
				friendService: func(ctrl *gomock.Controller) friendService.FriendService {
					m := friendService.NewMockFriendService(ctrl)
					return m
				},
				commonRoomMysqlRepository: func(ctrl *gomock.Controller) commonRoom.CommonRoomMysqlRepository {
					m := commonRoom.NewMockCommonRoomMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       2,
							},
							nil,
						)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       1,
							},
						).
						Return(
							&commonRoom.CommonRoom{
								RoomId:          "room1",
								HostUserId:      "0:test",
								RoomReleaseType: enum.RoomReleaseType_Private,
								Name:            "テストルーム",
								UserCount:       1,
							},
							nil,
						)
					return m
				},
				commonRoomUserMysqlRepository: func(ctrl *gomock.Controller) commonRoomUser.CommonRoomUserMysqlRepository {
					m := commonRoomUser.NewMockCommonRoomUserMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							gomock.Any(),
							"room1",
							"1:test",
						).
						Return(
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "1:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
							},
							nil,
						)
					m.EXPECT().
						Delete(
							gomock.Any(),
							gomock.Any(),
							&commonRoomUser.CommonRoomUser{
								RoomId:               "room1",
								UserId:               "1:test",
								RoomUserPositionType: enum.RoomUserPositionType_General,
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
				req: &RoomCheckOutRequest{
					UserId: "1:test",
					RoomId: "room1",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.commonRoomUserMysqlRepository.Delete", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &roomService{
				configService:                 tt.fields.configService(ctrl),
				friendService:                 tt.fields.friendService(ctrl),
				commonRoomMysqlRepository:     tt.fields.commonRoomMysqlRepository(ctrl),
				commonRoomUserMysqlRepository: tt.fields.commonRoomUserMysqlRepository(ctrl),
			}

			got, err := s.CheckOut(tt.args.ctx, tt.args.tx, tt.args.req)
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
