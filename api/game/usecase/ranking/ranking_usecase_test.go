package ranking

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	rankingServer "github.com/game-core/gocrafter/api/game/presentation/server/ranking"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/times"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	rankingService "github.com/game-core/gocrafter/pkg/domain/model/ranking"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingRoom"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingWorld"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRanking"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/masterRankingScope"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

func TestRankingUsecase_NewRankingUsecase(t *testing.T) {
	type args struct {
		rankingService     rankingService.RankingService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want RankingUsecase
	}{
		{
			name: "正常",
			args: args{
				rankingService:     nil,
				transactionService: nil,
			},
			want: &rankingUsecase{
				rankingService:     nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRankingUsecase(tt.args.rankingService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRankingUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankingUsecase_GetMaster(t *testing.T) {
	type fields struct {
		rankingService     func(ctrl *gomock.Controller) rankingService.RankingService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *rankingServer.RankingGetMasterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rankingServer.RankingGetMasterResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				rankingService: func(ctrl *gomock.Controller) rankingService.RankingService {
					m := rankingService.NewMockRankingService(ctrl)
					m.EXPECT().
						GetMaster(
							gomock.Any(),
							&rankingService.RankingGetMasterRequest{
								MasterRankingId: 1,
							},
						).
						Return(
							&rankingService.RankingGetMasterResponse{
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
				req: &rankingServer.RankingGetMasterRequest{
					UserId:          "0:test",
					MasterRankingId: 1,
				},
			},
			want: &rankingServer.RankingGetMasterResponse{
				MasterRanking: &rankingServer.MasterRanking{
					Id:                   1,
					MasterRankingEventId: 1,
					Name:                 "テストランキング",
					RankingScopeType:     rankingServer.RankingScopeType_World,
					RankingLimit:         10,
				},
				MasterRankingEvent: &rankingServer.MasterRankingEvent{
					Id:            1,
					Name:          "テストイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC))),
					EndAt:         nil,
				},
				MasterRankingScope: &rankingServer.MasterRankingScope{
					Id:               1,
					Name:             "テストイベント",
					RankingScopeType: rankingServer.RankingScopeType_World,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.rankingService.GetMaster",
			fields: fields{
				rankingService: func(ctrl *gomock.Controller) rankingService.RankingService {
					m := rankingService.NewMockRankingService(ctrl)
					m.EXPECT().
						GetMaster(
							gomock.Any(),
							&rankingService.RankingGetMasterRequest{
								MasterRankingId: 1,
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
				req: &rankingServer.RankingGetMasterRequest{
					UserId:          "0:test",
					MasterRankingId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.rankingService.GetMaster", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &rankingUsecase{
				rankingService:     tt.fields.rankingService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.GetMaster(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetMaster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankingUsecase_Get(t *testing.T) {
	type fields struct {
		rankingService     func(ctrl *gomock.Controller) rankingService.RankingService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *rankingServer.RankingGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rankingServer.RankingGetResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				rankingService: func(ctrl *gomock.Controller) rankingService.RankingService {
					m := rankingService.NewMockRankingService(ctrl)
					m.EXPECT().
						Get(
							gomock.Any(),
							gomock.Any(),
							&rankingService.RankingGetRequest{
								UserId:          "0:test1",
								MasterRankingId: 1,
								RoomId:          "room",
							},
						).
						Return(
							&rankingService.RankingGetResponse{
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
				req: &rankingServer.RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want: &rankingServer.RankingGetResponse{
				CommonRankingRooms: []*rankingServer.CommonRankingRoom{
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test2",
						Score:           2,
						RankedAt:        times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC))),
					},
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test1",
						Score:           1,
						RankedAt:        times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC))),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.rankingService.Get",
			fields: fields{
				rankingService: func(ctrl *gomock.Controller) rankingService.RankingService {
					m := rankingService.NewMockRankingService(ctrl)
					m.EXPECT().
						Get(
							gomock.Any(),
							gomock.Any(),
							&rankingService.RankingGetRequest{
								UserId:          "0:test1",
								MasterRankingId: 1,
								RoomId:          "room",
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
				req: &rankingServer.RankingGetRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.rankingService.Get", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &rankingUsecase{
				rankingService:     tt.fields.rankingService(ctrl),
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

func TestRankingUsecase_Update(t *testing.T) {
	type fields struct {
		rankingService     func(ctrl *gomock.Controller) rankingService.RankingService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *rankingServer.RankingUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rankingServer.RankingUpdateResponse
		wantErr error
	}{
		{
			name: "正常：更新できる",
			fields: fields{
				rankingService: func(ctrl *gomock.Controller) rankingService.RankingService {
					m := rankingService.NewMockRankingService(ctrl)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&rankingService.RankingUpdateRequest{
								UserId:          "0:test1",
								MasterRankingId: 1,
								RoomId:          "room",
								Score:           3,
							},
						).
						Return(
							&rankingService.RankingUpdateResponse{
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
										Score:           3,
										RankedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
									},
								},
								CommonRankingWorlds: commonRankingWorld.CommonRankingWorlds{},
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
				req: &rankingServer.RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           3,
				},
			},
			want: &rankingServer.RankingUpdateResponse{
				CommonRankingRooms: []*rankingServer.CommonRankingRoom{
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test2",
						Score:           2,
						RankedAt:        times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC))),
					},
					{
						MasterRankingId: 1,
						RoomId:          "room",
						UserId:          "0:test1",
						Score:           3,
						RankedAt:        times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC))),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.CommonMysqlBegin",
			fields: fields{
				rankingService: func(ctrl *gomock.Controller) rankingService.RankingService {
					m := rankingService.NewMockRankingService(ctrl)
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
				req: &rankingServer.RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           3,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.CommonMysqlBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.rankingService.Update",
			fields: fields{
				rankingService: func(ctrl *gomock.Controller) rankingService.RankingService {
					m := rankingService.NewMockRankingService(ctrl)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&rankingService.RankingUpdateRequest{
								UserId:          "0:test1",
								MasterRankingId: 1,
								RoomId:          "room",
								Score:           3,
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
				req: &rankingServer.RankingUpdateRequest{
					UserId:          "0:test1",
					MasterRankingId: 1,
					RoomId:          "room",
					Score:           3,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.rankingService.Update", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &rankingUsecase{
				rankingService:     tt.fields.rankingService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Update(tt.args.ctx, tt.args.req)
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
