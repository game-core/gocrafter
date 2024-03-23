package idleBonus

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusSchedule"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/userIdleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/item"
	"github.com/game-core/gocrafter/pkg/domain/model/item/masterItem"
	"github.com/game-core/gocrafter/pkg/domain/model/item/userItemBox"
)

func TestNewItemService_NewItemService(t *testing.T) {
	type args struct {
		itemService                            item.ItemService
		userIdleBonusMysqlRepository           userIdleBonus.UserIdleBonusMysqlRepository
		masterIdleBonusMysqlRepository         masterIdleBonus.MasterIdleBonusMysqlRepository
		masterIdleBonusEventMysqlRepository    masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository
		masterIdleBonusItemMysqlRepository     masterIdleBonusItem.MasterIdleBonusItemMysqlRepository
		masterIdleBonusScheduleMysqlRepository masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want IdleBonusService
	}{
		{
			name: "正常",
			args: args{
				itemService:                            nil,
				userIdleBonusMysqlRepository:           nil,
				masterIdleBonusMysqlRepository:         nil,
				masterIdleBonusEventMysqlRepository:    nil,
				masterIdleBonusItemMysqlRepository:     nil,
				masterIdleBonusScheduleMysqlRepository: nil,
			},
			want: &idleBonusService{
				itemService:                            nil,
				userIdleBonusMysqlRepository:           nil,
				masterIdleBonusMysqlRepository:         nil,
				masterIdleBonusEventMysqlRepository:    nil,
				masterIdleBonusItemMysqlRepository:     nil,
				masterIdleBonusScheduleMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIdleBonusService(
				tt.args.itemService,
				tt.args.userIdleBonusMysqlRepository,
				tt.args.masterIdleBonusMysqlRepository,
				tt.args.masterIdleBonusEventMysqlRepository,
				tt.args.masterIdleBonusItemMysqlRepository,
				tt.args.masterIdleBonusScheduleMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdleBonusService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewItemService_GetUser(t *testing.T) {
	type fields struct {
		itemService                            func(ctrl *gomock.Controller) item.ItemService
		userIdleBonusMysqlRepository           func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository
		masterIdleBonusMysqlRepository         func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository
		masterIdleBonusEventMysqlRepository    func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository
		masterIdleBonusItemMysqlRepository     func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository
		masterIdleBonusScheduleMysqlRepository func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *IdleBonusGetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IdleBonusGetUserResponse
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindListByUserId(
							gomock.Any(),
							"0:test",
						).
						Return(
							userIdleBonus.UserIdleBonuses{
								{
									UserId:            "0:test",
									MasterIdleBonusId: 1,
									ReceivedAt:        time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								},
								{
									UserId:            "0:test",
									MasterIdleBonusId: 2,
									ReceivedAt:        time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								},
								{
									UserId:            "0:test",
									MasterIdleBonusId: 3,
									ReceivedAt:        time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &IdleBonusGetUserRequest{
					UserId: "0:test",
				},
			},
			want: &IdleBonusGetUserResponse{
				UserIdleBonuses: userIdleBonus.UserIdleBonuses{
					{
						UserId:            "0:test",
						MasterIdleBonusId: 1,
						ReceivedAt:        time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					},
					{
						UserId:            "0:test",
						MasterIdleBonusId: 2,
						ReceivedAt:        time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					},
					{
						UserId:            "0:test",
						MasterIdleBonusId: 3,
						ReceivedAt:        time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userIdleBonusMysqlRepository.FindListByUserId",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindListByUserId(
							gomock.Any(),
							"0:test",
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &IdleBonusGetUserRequest{
					UserId: "0:test",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userIdleBonusMysqlRepository.FindListByUserId", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &idleBonusService{
				itemService:                            tt.fields.itemService(ctrl),
				userIdleBonusMysqlRepository:           tt.fields.userIdleBonusMysqlRepository(ctrl),
				masterIdleBonusMysqlRepository:         tt.fields.masterIdleBonusMysqlRepository(ctrl),
				masterIdleBonusEventMysqlRepository:    tt.fields.masterIdleBonusEventMysqlRepository(ctrl),
				masterIdleBonusItemMysqlRepository:     tt.fields.masterIdleBonusItemMysqlRepository(ctrl),
				masterIdleBonusScheduleMysqlRepository: tt.fields.masterIdleBonusScheduleMysqlRepository(ctrl),
			}

			got, err := s.GetUser(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewItemService_GetMaster(t *testing.T) {
	type fields struct {
		itemService                            func(ctrl *gomock.Controller) item.ItemService
		userIdleBonusMysqlRepository           func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository
		masterIdleBonusMysqlRepository         func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository
		masterIdleBonusEventMysqlRepository    func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository
		masterIdleBonusItemMysqlRepository     func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository
		masterIdleBonusScheduleMysqlRepository func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *IdleBonusGetMasterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IdleBonusGetMasterResponse
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テストログインボーナスイベント",
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
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusSchedule.MasterIdleBonusSchedules{
								{
									Id:                1,
									MasterIdleBonusId: 1,
									Step:              0,
									Name:              "ステップ0",
								},
								{
									Id:                2,
									MasterIdleBonusId: 1,
									Step:              1,
									Name:              "ステップ1",
								},
								{
									Id:                3,
									MasterIdleBonusId: 1,
									Step:              2,
									Name:              "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        1,
									MasterIdleBonusScheduleId: 1,
									MasterItemId:              1,
									Name:                      "テストログインボーナスアイテム1",
									Count:                     1,
								},
							},
							nil,
						)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        2,
									MasterIdleBonusScheduleId: 2,
									MasterItemId:              1,
									Name:                      "テストログインボーナスアイテム2",
									Count:                     1,
								},
							},
							nil,
						)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(3),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        3,
									MasterIdleBonusScheduleId: 3,
									MasterItemId:              1,
									Name:                      "テストログインボーナスアイテム3",
									Count:                     1,
								},
							},
							nil,
						)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &IdleBonusGetMasterRequest{
					MasterIdleBonusId: 1,
				},
			},
			want: &IdleBonusGetMasterResponse{
				MasterIdleBonus: &masterIdleBonus.MasterIdleBonus{
					Id:                     1,
					MasterIdleBonusEventId: 1,
					Name:                   "テストログインボーナス",
				},
				MasterIdleBonusEvent: &masterIdleBonusEvent.MasterIdleBonusEvent{
					Id:            1,
					Name:          "テストログインボーナスイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         nil,
				},
				MasterIdleBonusItems: masterIdleBonusItem.MasterIdleBonusItems{
					{
						Id:                        1,
						MasterIdleBonusScheduleId: 1,
						MasterItemId:              1,
						Name:                      "テストログインボーナスアイテム1",
						Count:                     1,
					},
					{
						Id:                        2,
						MasterIdleBonusScheduleId: 2,
						MasterItemId:              1,
						Name:                      "テストログインボーナスアイテム2",
						Count:                     1,
					},
					{
						Id:                        3,
						MasterIdleBonusScheduleId: 3,
						MasterItemId:              1,
						Name:                      "テストログインボーナスアイテム3",
						Count:                     1,
					},
				},
				MasterIdleBonusSchedules: masterIdleBonusSchedule.MasterIdleBonusSchedules{
					{
						Id:                1,
						MasterIdleBonusId: 1,
						Step:              0,
						Name:              "ステップ0",
					},
					{
						Id:                2,
						MasterIdleBonusId: 1,
						Step:              1,
						Name:              "ステップ1",
					},
					{
						Id:                3,
						MasterIdleBonusId: 1,
						Step:              2,
						Name:              "ステップ2",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterIdleBonusMysqlRepository.Find",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
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
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &IdleBonusGetMasterRequest{
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterIdleBonusMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.masterIdleBonusEventMysqlRepository.FindByMasterIdleBonusId",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
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
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &IdleBonusGetMasterRequest{
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterIdleBonusEventMysqlRepository.FindByMasterIdleBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.masterIdleBonusScheduleMysqlRepository.FindListByMasterIdleBonusId",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テストログインボーナスイベント",
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
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &IdleBonusGetMasterRequest{
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterIdleBonusScheduleMysqlRepository.FindListByMasterIdleBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.getItems",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テストログインボーナスイベント",
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
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusSchedule.MasterIdleBonusSchedules{
								{
									Id:                1,
									MasterIdleBonusId: 1,
									Step:              0,
									Name:              "ステップ0",
								},
								{
									Id:                2,
									MasterIdleBonusId: 1,
									Step:              1,
									Name:              "ステップ1",
								},
								{
									Id:                3,
									MasterIdleBonusId: 1,
									Step:              2,
									Name:              "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        1,
									MasterIdleBonusScheduleId: 1,
									MasterItemId:              1,
									Name:                      "テストログインボーナスアイテム1",
									Count:                     1,
								},
							},
							nil,
						)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &IdleBonusGetMasterRequest{
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getItems: failed to s.masterIdleBonusItemMysqlRepository.FindListByMasterIdleBonusScheduleId", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &idleBonusService{
				itemService:                            tt.fields.itemService(ctrl),
				userIdleBonusMysqlRepository:           tt.fields.userIdleBonusMysqlRepository(ctrl),
				masterIdleBonusMysqlRepository:         tt.fields.masterIdleBonusMysqlRepository(ctrl),
				masterIdleBonusEventMysqlRepository:    tt.fields.masterIdleBonusEventMysqlRepository(ctrl),
				masterIdleBonusItemMysqlRepository:     tt.fields.masterIdleBonusItemMysqlRepository(ctrl),
				masterIdleBonusScheduleMysqlRepository: tt.fields.masterIdleBonusScheduleMysqlRepository(ctrl),
			}

			got, err := s.GetMaster(tt.args.ctx, tt.args.req)
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

func TestNewItemService_Receive(t *testing.T) {
	type fields struct {
		itemService                            func(ctrl *gomock.Controller) item.ItemService
		userIdleBonusMysqlRepository           func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository
		masterIdleBonusMysqlRepository         func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository
		masterIdleBonusEventMysqlRepository    func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository
		masterIdleBonusItemMysqlRepository     func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository
		masterIdleBonusScheduleMysqlRepository func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		now time.Time
		req *IdleBonusReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IdleBonusReceiveResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusSchedule.MasterIdleBonusSchedules{
								{
									Id:                1,
									MasterIdleBonusId: 1,
									Step:              0,
									Name:              "ステップ0",
								},
								{
									Id:                2,
									MasterIdleBonusId: 1,
									Step:              1,
									Name:              "ステップ1",
								},
								{
									Id:                3,
									MasterIdleBonusId: 1,
									Step:              2,
									Name:              "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        1,
									MasterIdleBonusScheduleId: 1,
									MasterItemId:              1,
									Name:                      "テスト放置ボーナスアイテム1",
									Count:                     1,
								},
							},
							nil,
						)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        2,
									MasterIdleBonusScheduleId: 2,
									MasterItemId:              1,
									Name:                      "テスト放置ボーナスアイテム2",
									Count:                     1,
								},
							},
							nil,
						)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Receive(
							nil,
							nil,
							&item.ItemReceiveRequest{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
								Items: item.Items{
									{
										MasterItemId: 1,
										Count:        1,
									},
									{
										MasterItemId: 1,
										Count:        1,
									},
								},
							},
						).
						Return(
							&item.ItemReceiveResponse{
								UserItemBoxes: userItemBox.UserItemBoxes{
									{
										UserId:       "0:WntR-PyhOJeDiE5jodeR",
										MasterItemId: 1,
										Count:        1,
									},
									{
										UserId:       "0:WntR-PyhOJeDiE5jodeR",
										MasterItemId: 1,
										Count:        1,
									},
								},
								MasterItems: masterItem.MasterItems{
									{
										Id:           1,
										Name:         "テストアイテム",
										ResourceType: enum.ResourceType_Normal,
										RarityType:   enum.RarityType_N,
										Content:      "テストノーマルアイテム",
									},
									{
										Id:           1,
										Name:         "テストアイテム",
										ResourceType: enum.ResourceType_Normal,
										RarityType:   enum.RarityType_N,
										Content:      "テストノーマルアイテム",
									},
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want: &IdleBonusReceiveResponse{
				UserIdleBonus: &userIdleBonus.UserIdleBonus{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
					ReceivedAt:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				},
				MasterIdleBonus: &masterIdleBonus.MasterIdleBonus{
					Id:                     1,
					MasterIdleBonusEventId: 1,
					Name:                   "テスト放置ボーナス",
				},
				MasterIdleBonusEvent: &masterIdleBonusEvent.MasterIdleBonusEvent{
					Id:            1,
					Name:          "テスト放置ボーナスイベント",
					ResetHour:     9,
					IntervalHour:  1,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         nil,
				},
				MasterIdleBonusItems: masterIdleBonusItem.MasterIdleBonusItems{
					{
						Id:                        1,
						MasterIdleBonusScheduleId: 1,
						MasterItemId:              1,
						Name:                      "テスト放置ボーナスアイテム1",
						Count:                     1,
					},
					{
						Id:                        2,
						MasterIdleBonusScheduleId: 2,
						MasterItemId:              1,
						Name:                      "テスト放置ボーナスアイテム2",
						Count:                     1,
					},
				},
				MasterIdleBonusSchedules: masterIdleBonusSchedule.MasterIdleBonusSchedules{
					{
						Id:                1,
						MasterIdleBonusId: 1,
						Step:              0,
						Name:              "ステップ0",
					},
					{
						Id:                2,
						MasterIdleBonusId: 1,
						Step:              1,
						Name:              "ステップ1",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：取得できる（初回）",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want: &IdleBonusReceiveResponse{
				UserIdleBonus: &userIdleBonus.UserIdleBonus{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
					ReceivedAt:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				},
				MasterIdleBonus: &masterIdleBonus.MasterIdleBonus{
					Id:                     1,
					MasterIdleBonusEventId: 1,
					Name:                   "テスト放置ボーナス",
				},
				MasterIdleBonusEvent: &masterIdleBonusEvent.MasterIdleBonusEvent{
					Id:            1,
					Name:          "テスト放置ボーナスイベント",
					ResetHour:     9,
					IntervalHour:  1,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         nil,
				},
				MasterIdleBonusItems:     masterIdleBonusItem.MasterIdleBonusItems{},
				MasterIdleBonusSchedules: masterIdleBonusSchedule.MasterIdleBonusSchedules{},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterIdleBonusMysqlRepository.Find",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
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
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterIdleBonusMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：failed to s.getEvent: failed to s.masterIdleBonusEventMysqlRepository.FindByMasterIdleBonusId",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
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
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getEvent: failed to s.masterIdleBonusEventMysqlRepository.FindByMasterIdleBonusId", errors.NewTestError()),
		},
		{
			name: "異常：failed to s.getEvent: outside the event period",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 7, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getEvent: outside the event period"),
		},
		{
			name: "異常：s.userIdleBonusMysqlRepository.FindOrNil",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userIdleBonusMysqlRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：s.masterIdleBonusScheduleMysqlRepository.FindListByMasterIdleBonusId",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getSchedules: failed to s.masterIdleBonusScheduleMysqlRepository.FindListByMasterIdleBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.masterIdleBonusItemMysqlRepository.FindListByMasterIdleBonusScheduleId",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusSchedule.MasterIdleBonusSchedules{
								{
									Id:                1,
									MasterIdleBonusId: 1,
									Step:              0,
									Name:              "ステップ0",
								},
								{
									Id:                2,
									MasterIdleBonusId: 1,
									Step:              1,
									Name:              "ステップ1",
								},
								{
									Id:                3,
									MasterIdleBonusId: 1,
									Step:              2,
									Name:              "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getItems: failed to s.masterIdleBonusItemMysqlRepository.FindListByMasterIdleBonusScheduleId", errors.NewTestError()),
		},
		{
			name: "異常：s.itemService.Receive",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusSchedule.MasterIdleBonusSchedules{
								{
									Id:                1,
									MasterIdleBonusId: 1,
									Step:              0,
									Name:              "ステップ0",
								},
								{
									Id:                2,
									MasterIdleBonusId: 1,
									Step:              1,
									Name:              "ステップ1",
								},
								{
									Id:                3,
									MasterIdleBonusId: 1,
									Step:              2,
									Name:              "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        1,
									MasterIdleBonusScheduleId: 1,
									MasterItemId:              1,
									Name:                      "テスト放置ボーナスアイテム1",
									Count:                     1,
								},
							},
							nil,
						)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        2,
									MasterIdleBonusScheduleId: 2,
									MasterItemId:              1,
									Name:                      "テスト放置ボーナスアイテム2",
									Count:                     1,
								},
							},
							nil,
						)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Receive(
							nil,
							nil,
							&item.ItemReceiveRequest{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
								Items: item.Items{
									{
										MasterItemId: 1,
										Count:        1,
									},
									{
										MasterItemId: 1,
										Count:        1,
									},
								},
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
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.receive: failed to s.itemService.Receive", errors.NewTestError()),
		},
		{
			name: "異常：s.userIdleBonusMysqlRepository.Update",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusSchedule.MasterIdleBonusSchedules{
								{
									Id:                1,
									MasterIdleBonusId: 1,
									Step:              0,
									Name:              "ステップ0",
								},
								{
									Id:                2,
									MasterIdleBonusId: 1,
									Step:              1,
									Name:              "ステップ1",
								},
								{
									Id:                3,
									MasterIdleBonusId: 1,
									Step:              2,
									Name:              "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        1,
									MasterIdleBonusScheduleId: 1,
									MasterItemId:              1,
									Name:                      "テスト放置ボーナスアイテム1",
									Count:                     1,
								},
							},
							nil,
						)
					m.EXPECT().
						FindListByMasterIdleBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterIdleBonusItem.MasterIdleBonusItems{
								{
									Id:                        2,
									MasterIdleBonusScheduleId: 2,
									MasterItemId:              1,
									Name:                      "テスト放置ボーナスアイテム2",
									Count:                     1,
								},
							},
							nil,
						)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					m.EXPECT().
						Receive(
							nil,
							nil,
							&item.ItemReceiveRequest{
								UserId: "0:WntR-PyhOJeDiE5jodeR",
								Items: item.Items{
									{
										MasterItemId: 1,
										Count:        1,
									},
									{
										MasterItemId: 1,
										Count:        1,
									},
								},
							},
						).
						Return(
							&item.ItemReceiveResponse{
								UserItemBoxes: userItemBox.UserItemBoxes{
									{
										UserId:       "0:WntR-PyhOJeDiE5jodeR",
										MasterItemId: 1,
										Count:        1,
									},
									{
										UserId:       "0:WntR-PyhOJeDiE5jodeR",
										MasterItemId: 1,
										Count:        1,
									},
								},
								MasterItems: masterItem.MasterItems{
									{
										Id:           1,
										Name:         "テストアイテム",
										ResourceType: enum.ResourceType_Normal,
										RarityType:   enum.RarityType_N,
										Content:      "テストノーマルアイテム",
									},
									{
										Id:           1,
										Name:         "テストアイテム",
										ResourceType: enum.ResourceType_Normal,
										RarityType:   enum.RarityType_N,
										Content:      "テストノーマルアイテム",
									},
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.update: failed to s.userIdleBonusMysqlRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.userIdleBonusMysqlRepository.Create",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userIdleBonusMysqlRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：masterIdleBonusSchedules.GetStep",
			fields: fields{
				masterIdleBonusMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusMysqlRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonus.MasterIdleBonus{
								Id:                     1,
								MasterIdleBonusEventId: 1,
								Name:                   "テスト放置ボーナス",
							},
							nil,
						)
					return m
				},
				masterIdleBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventMysqlRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterIdleBonusEvent.MasterIdleBonusEvent{
								Id:            1,
								Name:          "テスト放置ボーナスイベント",
								ResetHour:     9,
								IntervalHour:  1,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterIdleBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleMysqlRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterIdleBonusId(
							nil,
							int64(1),
						).
						Return(
							masterIdleBonusSchedule.MasterIdleBonusSchedules{
								{
									Id:                1,
									MasterIdleBonusId: 1,
									Step:              0,
									Name:              "ステップ0",
								},
								{
									Id:                2,
									MasterIdleBonusId: 1,
									Step:              1,
									Name:              "ステップ1",
								},
								{
									Id:                3,
									MasterIdleBonusId: 1,
									Step:              2,
									Name:              "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterIdleBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemMysqlRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemMysqlRepository(ctrl)
					return m
				},
				userIdleBonusMysqlRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusMysqlRepository {
					m := userIdleBonus.NewMockUserIdleBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userIdleBonus.UserIdleBonus{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
								ReceivedAt:        time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 1, 9, 1, 0, 0, time.UTC),
				req: &IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getSchedules: failed to masterIdleBonusSchedules.GetStep: already received"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &idleBonusService{
				itemService:                            tt.fields.itemService(ctrl),
				userIdleBonusMysqlRepository:           tt.fields.userIdleBonusMysqlRepository(ctrl),
				masterIdleBonusMysqlRepository:         tt.fields.masterIdleBonusMysqlRepository(ctrl),
				masterIdleBonusEventMysqlRepository:    tt.fields.masterIdleBonusEventMysqlRepository(ctrl),
				masterIdleBonusItemMysqlRepository:     tt.fields.masterIdleBonusItemMysqlRepository(ctrl),
				masterIdleBonusScheduleMysqlRepository: tt.fields.masterIdleBonusScheduleMysqlRepository(ctrl),
			}

			got, err := s.Receive(tt.args.ctx, tt.args.tx, tt.args.now, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Receive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Receive() = %v, want %v", got, tt.want)
			}
		})
	}
}
