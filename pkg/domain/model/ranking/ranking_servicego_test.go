package ranking

import (
	"context"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/golang/mock/gomock"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingWorld"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRanking"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingScope"
	roomService "github.com/game-core/gocrafter/pkg/domain/model/room"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/room/commonRoomUser"
)

func TestNewRankingService_NewRankingService(t *testing.T) {
	type args struct {
		roomService                       roomService.RoomService
		commonRankingRoomMysqlRepository  commonRankingRoom.CommonRankingRoomMysqlRepository
		commonRankingWorldMysqlRepository commonRankingWorld.CommonRankingWorldMysqlRepository
		masterRankingMysqlRepository      masterRanking.MasterRankingMysqlRepository
		masterRankingEventMysqlRepository masterRankingEvent.MasterRankingEventMysqlRepository
		masterRankingScopeMysqlRepository masterRankingScope.MasterRankingScopeMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want RankingService
	}{
		{
			name: "正常",
			args: args{
				roomService:                       nil,
				commonRankingRoomMysqlRepository:  nil,
				commonRankingWorldMysqlRepository: nil,
				masterRankingMysqlRepository:      nil,
				masterRankingEventMysqlRepository: nil,
				masterRankingScopeMysqlRepository: nil,
			},
			want: &rankingService{
				roomService:                       nil,
				commonRankingRoomMysqlRepository:  nil,
				commonRankingWorldMysqlRepository: nil,
				masterRankingMysqlRepository:      nil,
				masterRankingEventMysqlRepository: nil,
				masterRankingScopeMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRankingService(
				tt.args.roomService,
				tt.args.commonRankingRoomMysqlRepository,
				tt.args.commonRankingWorldMysqlRepository,
				tt.args.masterRankingMysqlRepository,
				tt.args.masterRankingEventMysqlRepository,
				tt.args.masterRankingScopeMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRankingService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankingService_GetMaster(t *testing.T) {
	type fields struct {
		roomService                       func(ctrl *gomock.Controller) roomService.RoomService
		commonRankingRoomMysqlRepository  func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository
		commonRankingWorldMysqlRepository func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository
		masterRankingMysqlRepository      func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository
		masterRankingEventMysqlRepository func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository
		masterRankingScopeMysqlRepository func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *RankingGetMasterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RankingGetMasterResponse
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					m.EXPECT().
						FindByRankingScopeType(
							nil,
							enum.RankingScopeType_World,
						).
						Return(
							&masterRankingScope.MasterRankingScope{
								Id:               1,
								Name:             "テストイベント",
								RankingScopeType: enum.RankingScopeType_World,
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RankingGetMasterRequest{
					MasterRankingId: 1,
				},
			},
			want: &RankingGetMasterResponse{
				MasterRanking: &masterRanking.MasterRanking{
					Id:                   1,
					MasterRankingEventId: 1,
					Name:                 "テストランキング",
					RankingScopeType:     enum.RankingScopeType_World,
					RankingLimit:         10,
				},
				MasterRankingEvent: &masterRankingEvent.MasterRankingEvent{
					Id:            1,
					Name:          "テストイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         nil,
				},
				MasterRankingScope: &masterRankingScope.MasterRankingScope{
					Id:               1,
					Name:             "テストイベント",
					RankingScopeType: enum.RankingScopeType_World,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterRankingMysqlRepository.Find",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RankingGetMasterRequest{
					MasterRankingId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterRankingMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.masterRankingEventMysqlRepository.Find",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &RankingGetMasterRequest{
					MasterRankingId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterRankingEventMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.masterRankingScopeMysqlRepository.FindByRankingScopeType",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					m.EXPECT().
						FindByRankingScopeType(
							nil,
							enum.RankingScopeType_World,
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
				req: &RankingGetMasterRequest{
					MasterRankingId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterRankingScopeMysqlRepository.FindByRankingScopeType", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &rankingService{
				roomService:                       tt.fields.roomService(ctrl),
				commonRankingRoomMysqlRepository:  tt.fields.commonRankingRoomMysqlRepository(ctrl),
				commonRankingWorldMysqlRepository: tt.fields.commonRankingWorldMysqlRepository(ctrl),
				masterRankingMysqlRepository:      tt.fields.masterRankingMysqlRepository(ctrl),
				masterRankingEventMysqlRepository: tt.fields.masterRankingEventMysqlRepository(ctrl),
				masterRankingScopeMysqlRepository: tt.fields.masterRankingScopeMysqlRepository(ctrl),
			}

			got, err := s.GetMaster(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Getmaster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankingService_Get(t *testing.T) {
	type fields struct {
		roomService                       func(ctrl *gomock.Controller) roomService.RoomService
		commonRankingRoomMysqlRepository  func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository
		commonRankingWorldMysqlRepository func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository
		masterRankingMysqlRepository      func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository
		masterRankingEventMysqlRepository func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository
		masterRankingScopeMysqlRepository func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository
	}
	type args struct {
		ctx context.Context
		now time.Time
		req *RankingGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RankingGetResponse
		wantErr error
	}{
		{
			name: "正常：取得できる場合（ルームランキング）",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							&roomService.RoomCheckResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "room",
									HostUserId:      "0:test1",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "room",
									UserId:               "0:test1",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingIdAndRoomId(
							nil,
							int64(1),
							"room",
						).
						Return(
							commonRankingRoom.CommonRankingRooms{
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test1",
									Score:           1,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test2",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want: &RankingGetResponse{
				CommonRankingRooms: commonRankingRoom.CommonRankingRooms{
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test2",
						Score:           2,
						RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
					},
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test1",
						Score:           1,
						RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
					},
				},
				CommonRankingWorlds: commonRankingWorld.CommonRankingWorlds{},
			},
			wantErr: nil,
		},
		{
			name: "正常：取得できる場合（ワールドランキング）",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingId(
							nil,
							int64(1),
						).
						Return(
							commonRankingWorld.CommonRankingWorlds{
								{
									MasterRankingId: 1,
									UserId:          "0:test1",
									Score:           1,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									UserId:          "0:test2",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want: &RankingGetResponse{
				CommonRankingRooms: commonRankingRoom.CommonRankingRooms{},
				CommonRankingWorlds: commonRankingWorld.CommonRankingWorlds{
					{
						MasterRankingId: 1,
						UserId:          "0:test2",
						Score:           2,
						RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
					},
					{
						MasterRankingId: 1,
						UserId:          "0:test1",
						Score:           1,
						RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterRankingMysqlRepository.Find",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterRankingMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.getEvent: failed to s.masterRankingEventMysqlRepository.Find",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getEvent: failed to s.masterRankingEventMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：failed to s.getEvent: outside the event period",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getEvent: outside the event period"),
		},
		{
			name: "異常：s.getRoomRankings: failed to s.roomService.Check",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getRoomRankings: failed to s.roomService.Check", errors.NewTestError()),
		},
		{
			name: "異常：s.getRoomRankings: failed to s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							&roomService.RoomCheckResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "room",
									HostUserId:      "0:test1",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "room",
									UserId:               "0:test1",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingIdAndRoomId(
							nil,
							int64(1),
							"room",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getRoomRankings: failed to s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId", errors.NewTestError()),
		},
		{
			name: "異常：s.getWorldRankings: failed to s.commonRankingWorldMysqlRepository.FindListByMasterRankingId",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getWorldRankings: failed to s.commonRankingWorldMysqlRepository.FindListByMasterRankingId", errors.NewTestError()),
		},
		{
			name: "異常：RankingScopeType that does not exist",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     999,
								RankingLimit:         10,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want:    nil,
			wantErr: errors.NewError("RankingScopeType that does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &rankingService{
				roomService:                       tt.fields.roomService(ctrl),
				commonRankingRoomMysqlRepository:  tt.fields.commonRankingRoomMysqlRepository(ctrl),
				commonRankingWorldMysqlRepository: tt.fields.commonRankingWorldMysqlRepository(ctrl),
				masterRankingMysqlRepository:      tt.fields.masterRankingMysqlRepository(ctrl),
				masterRankingEventMysqlRepository: tt.fields.masterRankingEventMysqlRepository(ctrl),
				masterRankingScopeMysqlRepository: tt.fields.masterRankingScopeMysqlRepository(ctrl),
			}

			got, err := s.Get(tt.args.ctx, tt.args.now, tt.args.req)
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

func TestRankingService_Update(t *testing.T) {
	type fields struct {
		roomService                       func(ctrl *gomock.Controller) roomService.RoomService
		commonRankingRoomMysqlRepository  func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository
		commonRankingWorldMysqlRepository func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository
		masterRankingMysqlRepository      func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository
		masterRankingEventMysqlRepository func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository
		masterRankingScopeMysqlRepository func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		now time.Time
		req *RankingUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RankingUpdateResponse
		wantErr error
	}{
		{
			name: "正常：追加できる場合（ルームランキング）",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							&roomService.RoomCheckResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "room",
									HostUserId:      "0:test1",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "room",
									UserId:               "0:test1",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingIdAndRoomId(
							nil,
							int64(1),
							"room",
						).
						Return(
							commonRankingRoom.CommonRankingRooms{
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test2",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test2",
								Score:           2,
								RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want: &RankingUpdateResponse{
				CommonRankingRooms: commonRankingRoom.CommonRankingRooms{
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test3",
						Score:           3,
						RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
					},
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test1",
						Score:           4,
						RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
					},
				},
				CommonRankingWorlds: commonRankingWorld.CommonRankingWorlds{},
			},
			wantErr: nil,
		},
		{
			name: "正常：更新できる場合（ルームランキング）",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							&roomService.RoomCheckResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "room",
									HostUserId:      "0:test1",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "room",
									UserId:               "0:test1",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingIdAndRoomId(
							nil,
							int64(1),
							"room",
						).
						Return(
							commonRankingRoom.CommonRankingRooms{
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test1",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want: &RankingUpdateResponse{
				CommonRankingRooms: commonRankingRoom.CommonRankingRooms{
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test3",
						Score:           3,
						RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
					},
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test1",
						Score:           4,
						RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
					},
				},
				CommonRankingWorlds: commonRankingWorld.CommonRankingWorlds{},
			},
			wantErr: nil,
		},
		{
			name: "正常：追加できる場合（ワールドランキング）",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingId(
							nil,
							int64(1),
						).
						Return(
							commonRankingWorld.CommonRankingWorlds{
								{
									MasterRankingId: 1,
									UserId:          "0:test2",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&commonRankingWorld.CommonRankingWorld{
								MasterRankingId: 1,
								UserId:          "0:test2",
								Score:           2,
								RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&commonRankingWorld.CommonRankingWorld{
								MasterRankingId: 1,
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&commonRankingWorld.CommonRankingWorld{
								MasterRankingId: 1,
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want: &RankingUpdateResponse{
				CommonRankingRooms: commonRankingRoom.CommonRankingRooms{},
				CommonRankingWorlds: commonRankingWorld.CommonRankingWorlds{
					{
						MasterRankingId: 1,
						UserId:          "0:test3",
						Score:           3,
						RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
					},
					{
						MasterRankingId: 1,
						UserId:          "0:test1",
						Score:           4,
						RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：更新できる場合（ワールドランキング）",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingId(
							nil,
							int64(1),
						).
						Return(
							commonRankingWorld.CommonRankingWorlds{
								{
									MasterRankingId: 1,
									UserId:          "0:test1",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&commonRankingWorld.CommonRankingWorld{
								MasterRankingId: 1,
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&commonRankingWorld.CommonRankingWorld{
								MasterRankingId: 1,
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want: &RankingUpdateResponse{
				CommonRankingRooms: commonRankingRoom.CommonRankingRooms{},
				CommonRankingWorlds: commonRankingWorld.CommonRankingWorlds{
					{
						MasterRankingId: 1,
						UserId:          "0:test3",
						Score:           3,
						RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
					},
					{
						MasterRankingId: 1,
						UserId:          "0:test1",
						Score:           4,
						RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterRankingMysqlRepository.Find",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterRankingMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.getEvent: failed to s.masterRankingEventMysqlRepository.Find",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getEvent: failed to s.masterRankingEventMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：failed to s.getEvent: outside the event period",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getEvent: outside the event period"),
		},
		{
			name: "異常：s.updateRoomRankings: failed to s.roomService.Check",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.updateRoomRankings: failed to s.roomService.Check", errors.NewTestError()),
		},
		{
			name: "異常：s.updateRoomRankings: failed to s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							&roomService.RoomCheckResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "room",
									HostUserId:      "0:test1",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "room",
									UserId:               "0:test1",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingIdAndRoomId(
							nil,
							int64(1),
							"room",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.updateRoomRankings: failed to s.commonRankingRoomMysqlRepository.FindListByMasterRankingIdAndRoomId", errors.NewTestError()),
		},
		{
			name: "異常：s.updateRoomRankings: failed to s.commonRankingRoomMysqlRepository.Delete",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							&roomService.RoomCheckResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "room",
									HostUserId:      "0:test1",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "room",
									UserId:               "0:test1",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingIdAndRoomId(
							nil,
							int64(1),
							"room",
						).
						Return(
							commonRankingRoom.CommonRankingRooms{
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test2",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test2",
								Score:           2,
								RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
							},
						).
						Return(
							errors.NewTestError(),
						)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.updateRoomRankings: failed to s.commonRankingRoomMysqlRepository.Delete", errors.NewTestError()),
		},
		{
			name: "異常：s.updateRoomRankings: failed to s.updateRoomRanking: failed to s.commonRankingRoomMysqlRepository.Create",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							&roomService.RoomCheckResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "room",
									HostUserId:      "0:test1",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "room",
									UserId:               "0:test1",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingIdAndRoomId(
							nil,
							int64(1),
							"room",
						).
						Return(
							commonRankingRoom.CommonRankingRooms{
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test2",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test2",
								Score:           2,
								RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.updateRoomRankings: failed to s.updateRoomRanking: failed to s.commonRankingRoomMysqlRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：s.updateRoomRankings: failed to s.updateRoomRanking: failed to s.commonRankingRoomMysqlRepository.Update",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							&roomService.RoomCheckResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "room",
									HostUserId:      "0:test1",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "room",
									UserId:               "0:test1",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingIdAndRoomId(
							nil,
							int64(1),
							"room",
						).
						Return(
							commonRankingRoom.CommonRankingRooms{
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test1",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.updateRoomRankings: failed to s.updateRoomRanking: failed to s.commonRankingRoomMysqlRepository.Update", errors.NewTestError()),
		},
		{
			name: "正常：更新できる場合（ルームランキング）",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					m.EXPECT().
						Check(
							nil,
							&roomService.RoomCheckRequest{
								UserId: "0:test1",
								RoomId: "room",
							},
						).
						Return(
							&roomService.RoomCheckResponse{
								CommonRoom: &commonRoom.CommonRoom{
									RoomId:          "room",
									HostUserId:      "0:test1",
									RoomReleaseType: enum.RoomReleaseType_Public,
									Name:            "テストルーム",
									UserCount:       1,
								},
								CommonRoomUser: &commonRoomUser.CommonRoomUser{
									RoomId:               "room",
									UserId:               "0:test1",
									RoomUserPositionType: enum.RoomUserPositionType_General,
								},
							},
							nil,
						)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingIdAndRoomId(
							nil,
							int64(1),
							"room",
						).
						Return(
							commonRankingRoom.CommonRankingRooms{
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test1",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									RoomId:          "room",
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&commonRankingRoom.CommonRankingRoom{
								MasterRankingId: 1,
								RoomId:          "room",
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_Room,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want: &RankingUpdateResponse{
				CommonRankingRooms: commonRankingRoom.CommonRankingRooms{
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test3",
						Score:           3,
						RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
					},
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test1",
						Score:           4,
						RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
					},
				},
				CommonRankingWorlds: commonRankingWorld.CommonRankingWorlds{},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.updateWorldRankings: failed to s.commonRankingWorldMysqlRepository.FindListByMasterRankingId",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.updateWorldRankings: failed to s.commonRankingWorldMysqlRepository.FindListByMasterRankingId", errors.NewTestError()),
		},
		{
			name: "異常：s.updateWorldRankings: failed to s.commonRankingWorldMysqlRepository.Delete",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingId(
							nil,
							int64(1),
						).
						Return(
							commonRankingWorld.CommonRankingWorlds{
								{
									MasterRankingId: 1,
									UserId:          "0:test2",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&commonRankingWorld.CommonRankingWorld{
								MasterRankingId: 1,
								UserId:          "0:test2",
								Score:           2,
								RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
							},
						).
						Return(
							errors.NewTestError(),
						)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.updateWorldRankings: failed to s.commonRankingWorldMysqlRepository.Delete", errors.NewTestError()),
		},
		{
			name: "異常：s.updateWorldRankings: failed to s.updateWorldRanking: failed to s.commonRankingWorldMysqlRepository.Create",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingId(
							nil,
							int64(1),
						).
						Return(
							commonRankingWorld.CommonRankingWorlds{
								{
									MasterRankingId: 1,
									UserId:          "0:test2",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&commonRankingWorld.CommonRankingWorld{
								MasterRankingId: 1,
								UserId:          "0:test2",
								Score:           2,
								RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&commonRankingWorld.CommonRankingWorld{
								MasterRankingId: 1,
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.updateWorldRankings: failed to s.updateWorldRanking: failed to s.commonRankingWorldMysqlRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：s.updateWorldRankings: failed to s.updateWorldRanking: failed to s.commonRankingWorldMysqlRepository.Update",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterRankingId(
							nil,
							int64(1),
						).
						Return(
							commonRankingWorld.CommonRankingWorlds{
								{
									MasterRankingId: 1,
									UserId:          "0:test1",
									Score:           2,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
								{
									MasterRankingId: 1,
									UserId:          "0:test3",
									Score:           3,
									RankedAt:        time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
								},
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&commonRankingWorld.CommonRankingWorld{
								MasterRankingId: 1,
								UserId:          "0:test1",
								Score:           4,
								RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     enum.RankingScopeType_World,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.updateWorldRankings: failed to s.updateWorldRanking: failed to s.commonRankingWorldMysqlRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：RankingScopeType that does not exist",
			fields: fields{
				roomService: func(ctrl *gomock.Controller) roomService.RoomService {
					m := roomService.NewMockRoomService(ctrl)
					return m
				},
				commonRankingRoomMysqlRepository: func(ctrl *gomock.Controller) commonRankingRoom.CommonRankingRoomMysqlRepository {
					m := commonRankingRoom.NewMockCommonRankingRoomMysqlRepository(ctrl)
					return m
				},
				commonRankingWorldMysqlRepository: func(ctrl *gomock.Controller) commonRankingWorld.CommonRankingWorldMysqlRepository {
					m := commonRankingWorld.NewMockCommonRankingWorldMysqlRepository(ctrl)
					return m
				},
				masterRankingMysqlRepository: func(ctrl *gomock.Controller) masterRanking.MasterRankingMysqlRepository {
					m := masterRanking.NewMockMasterRankingMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRanking.MasterRanking{
								Id:                   1,
								MasterRankingEventId: 1,
								Name:                 "テストランキング",
								RankingScopeType:     999,
								RankingLimit:         2,
							},
							nil,
						)
					return m
				},
				masterRankingEventMysqlRepository: func(ctrl *gomock.Controller) masterRankingEvent.MasterRankingEventMysqlRepository {
					m := masterRankingEvent.NewMockMasterRankingEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterRankingEvent.MasterRankingEvent{
								Id:            1,
								Name:          "テストイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterRankingScopeMysqlRepository: func(ctrl *gomock.Controller) masterRankingScope.MasterRankingScopeMysqlRepository {
					m := masterRankingScope.NewMockMasterRankingScopeMysqlRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				req: &RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           4,
				},
			},
			want:    nil,
			wantErr: errors.NewError("RankingScopeType that does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &rankingService{
				roomService:                       tt.fields.roomService(ctrl),
				commonRankingRoomMysqlRepository:  tt.fields.commonRankingRoomMysqlRepository(ctrl),
				commonRankingWorldMysqlRepository: tt.fields.commonRankingWorldMysqlRepository(ctrl),
				masterRankingMysqlRepository:      tt.fields.masterRankingMysqlRepository(ctrl),
				masterRankingEventMysqlRepository: tt.fields.masterRankingEventMysqlRepository(ctrl),
				masterRankingScopeMysqlRepository: tt.fields.masterRankingScopeMysqlRepository(ctrl),
			}

			got, err := s.Update(tt.args.ctx, tt.args.tx, tt.args.now, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
