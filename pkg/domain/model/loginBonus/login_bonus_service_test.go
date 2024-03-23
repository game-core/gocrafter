package loginBonus

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/item"
	"github.com/game-core/gocrafter/pkg/domain/model/item/masterItem"
	"github.com/game-core/gocrafter/pkg/domain/model/item/userItemBox"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusSchedule"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/userLoginBonus"
)

func TestNewItemService_NewItemService(t *testing.T) {
	type args struct {
		itemService                             item.ItemService
		userLoginBonusMysqlRepository           userLoginBonus.UserLoginBonusMysqlRepository
		masterLoginBonusMysqlRepository         masterLoginBonus.MasterLoginBonusMysqlRepository
		masterLoginBonusEventMysqlRepository    masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository
		masterLoginBonusItemMysqlRepository     masterLoginBonusItem.MasterLoginBonusItemMysqlRepository
		masterLoginBonusScheduleMysqlRepository masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository
	}
	tests := []struct {
		name string
		args args
		want LoginBonusService
	}{
		{
			name: "正常",
			args: args{
				itemService:                             nil,
				userLoginBonusMysqlRepository:           nil,
				masterLoginBonusMysqlRepository:         nil,
				masterLoginBonusEventMysqlRepository:    nil,
				masterLoginBonusItemMysqlRepository:     nil,
				masterLoginBonusScheduleMysqlRepository: nil,
			},
			want: &loginBonusService{
				itemService:                             nil,
				userLoginBonusMysqlRepository:           nil,
				masterLoginBonusMysqlRepository:         nil,
				masterLoginBonusEventMysqlRepository:    nil,
				masterLoginBonusItemMysqlRepository:     nil,
				masterLoginBonusScheduleMysqlRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLoginBonusService(
				tt.args.itemService,
				tt.args.userLoginBonusMysqlRepository,
				tt.args.masterLoginBonusMysqlRepository,
				tt.args.masterLoginBonusEventMysqlRepository,
				tt.args.masterLoginBonusItemMysqlRepository,
				tt.args.masterLoginBonusScheduleMysqlRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginBonusService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewItemService_GetUser(t *testing.T) {
	type fields struct {
		itemService                             func(ctrl *gomock.Controller) item.ItemService
		userLoginBonusMysqlRepository           func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository
		masterLoginBonusMysqlRepository         func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository
		masterLoginBonusEventMysqlRepository    func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository
		masterLoginBonusItemMysqlRepository     func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository
		masterLoginBonusScheduleMysqlRepository func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *LoginBonusGetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LoginBonusGetUserResponse
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindListByUserId(
							gomock.Any(),
							"0:test",
						).
						Return(
							userLoginBonus.UserLoginBonuses{
								{
									UserId:             "0:test",
									MasterLoginBonusId: 1,
									ReceivedAt:         time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								},
								{
									UserId:             "0:test",
									MasterLoginBonusId: 2,
									ReceivedAt:         time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
								},
								{
									UserId:             "0:test",
									MasterLoginBonusId: 3,
									ReceivedAt:         time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
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
				req: &LoginBonusGetUserRequest{
					UserId: "0:test",
				},
			},
			want: &LoginBonusGetUserResponse{
				UserLoginBonuses: userLoginBonus.UserLoginBonuses{
					{
						UserId:             "0:test",
						MasterLoginBonusId: 1,
						ReceivedAt:         time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					},
					{
						UserId:             "0:test",
						MasterLoginBonusId: 2,
						ReceivedAt:         time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					},
					{
						UserId:             "0:test",
						MasterLoginBonusId: 3,
						ReceivedAt:         time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userLoginBonusMysqlRepository.FindListByUserId",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
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
				req: &LoginBonusGetUserRequest{
					UserId: "0:test",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userLoginBonusMysqlRepository.FindListByUserId", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &loginBonusService{
				itemService:                             tt.fields.itemService(ctrl),
				userLoginBonusMysqlRepository:           tt.fields.userLoginBonusMysqlRepository(ctrl),
				masterLoginBonusMysqlRepository:         tt.fields.masterLoginBonusMysqlRepository(ctrl),
				masterLoginBonusEventMysqlRepository:    tt.fields.masterLoginBonusEventMysqlRepository(ctrl),
				masterLoginBonusItemMysqlRepository:     tt.fields.masterLoginBonusItemMysqlRepository(ctrl),
				masterLoginBonusScheduleMysqlRepository: tt.fields.masterLoginBonusScheduleMysqlRepository(ctrl),
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
		itemService                             func(ctrl *gomock.Controller) item.ItemService
		userLoginBonusMysqlRepository           func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository
		masterLoginBonusMysqlRepository         func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository
		masterLoginBonusEventMysqlRepository    func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository
		masterLoginBonusItemMysqlRepository     func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository
		masterLoginBonusScheduleMysqlRepository func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository
	}
	type args struct {
		ctx context.Context
		req *LoginBonusGetMasterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LoginBonusGetMasterResponse
		wantErr error
	}{
		{
			name: "正常：取得できる場合",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         1,
									MasterLoginBonusScheduleId: 1,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム1",
									Count:                      1,
								},
							},
							nil,
						)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         2,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム2",
									Count:                      1,
								},
							},
							nil,
						)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(3),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         3,
									MasterLoginBonusScheduleId: 3,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム3",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &LoginBonusGetMasterRequest{
					MasterLoginBonusId: 1,
				},
			},
			want: &LoginBonusGetMasterResponse{
				MasterLoginBonus: &masterLoginBonus.MasterLoginBonus{
					Id:                      1,
					MasterLoginBonusEventId: 1,
					Name:                    "テストログインボーナス",
				},
				MasterLoginBonusEvent: &masterLoginBonusEvent.MasterLoginBonusEvent{
					Id:            1,
					Name:          "テストログインボーナスイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         nil,
				},
				MasterLoginBonusItems: masterLoginBonusItem.MasterLoginBonusItems{
					{
						Id:                         1,
						MasterLoginBonusScheduleId: 1,
						MasterItemId:               1,
						Name:                       "テストログインボーナスアイテム1",
						Count:                      1,
					},
					{
						Id:                         2,
						MasterLoginBonusScheduleId: 2,
						MasterItemId:               1,
						Name:                       "テストログインボーナスアイテム2",
						Count:                      1,
					},
					{
						Id:                         3,
						MasterLoginBonusScheduleId: 3,
						MasterItemId:               1,
						Name:                       "テストログインボーナスアイテム3",
						Count:                      1,
					},
				},
				MasterLoginBonusSchedules: masterLoginBonusSchedule.MasterLoginBonusSchedules{
					{
						Id:                 1,
						MasterLoginBonusId: 1,
						Step:               0,
						Name:               "ステップ0",
					},
					{
						Id:                 2,
						MasterLoginBonusId: 1,
						Step:               1,
						Name:               "ステップ1",
					},
					{
						Id:                 3,
						MasterLoginBonusId: 1,
						Step:               2,
						Name:               "ステップ2",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterLoginBonusMysqlRepository.Find",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
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
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &LoginBonusGetMasterRequest{
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterLoginBonusMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.masterLoginBonusEventMysqlRepository.FindByMasterLoginBonusId",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &LoginBonusGetMasterRequest{
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterLoginBonusEventMysqlRepository.FindByMasterLoginBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.masterLoginBonusScheduleMysqlRepository.FindListByMasterLoginBonusId",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &LoginBonusGetMasterRequest{
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterLoginBonusScheduleMysqlRepository.FindListByMasterLoginBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.getItems",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         1,
									MasterLoginBonusScheduleId: 1,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム1",
									Count:                      1,
								},
							},
							nil,
						)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) item.ItemService {
					m := item.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &LoginBonusGetMasterRequest{
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getItems: failed to s.masterLoginBonusItemMysqlRepository.FindListByMasterLoginBonusScheduleId", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &loginBonusService{
				itemService:                             tt.fields.itemService(ctrl),
				userLoginBonusMysqlRepository:           tt.fields.userLoginBonusMysqlRepository(ctrl),
				masterLoginBonusMysqlRepository:         tt.fields.masterLoginBonusMysqlRepository(ctrl),
				masterLoginBonusEventMysqlRepository:    tt.fields.masterLoginBonusEventMysqlRepository(ctrl),
				masterLoginBonusItemMysqlRepository:     tt.fields.masterLoginBonusItemMysqlRepository(ctrl),
				masterLoginBonusScheduleMysqlRepository: tt.fields.masterLoginBonusScheduleMysqlRepository(ctrl),
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
		itemService                             func(ctrl *gomock.Controller) item.ItemService
		userLoginBonusMysqlRepository           func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository
		masterLoginBonusMysqlRepository         func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository
		masterLoginBonusEventMysqlRepository    func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository
		masterLoginBonusItemMysqlRepository     func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository
		masterLoginBonusScheduleMysqlRepository func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		now time.Time
		req *LoginBonusReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LoginBonusReceiveResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる場合（既に取得したことがあるアイテム）",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         2,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
								},
								MasterItems: masterItem.MasterItems{
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want: &LoginBonusReceiveResponse{
				UserLoginBonus: &userLoginBonus.UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				},
				MasterLoginBonus: &masterLoginBonus.MasterLoginBonus{
					Id:                      1,
					MasterLoginBonusEventId: 1,
					Name:                    "テストログインボーナス",
				},
				MasterLoginBonusEvent: &masterLoginBonusEvent.MasterLoginBonusEvent{
					Id:            1,
					Name:          "テストログインボーナスイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         nil,
				},
				MasterLoginBonusItems: masterLoginBonusItem.MasterLoginBonusItems{
					{
						Id:                         2,
						MasterLoginBonusScheduleId: 2,
						MasterItemId:               1,
						Name:                       "テストログインボーナスアイテム",
						Count:                      1,
					},
				},
				MasterLoginBonusSchedule: &masterLoginBonusSchedule.MasterLoginBonusSchedule{
					Id:                 2,
					MasterLoginBonusId: 1,
					Step:               1,
					Name:               "ステップ1",
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：作成できる場合（取得したことがないアイテム）",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         2,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
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
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
								},
								MasterItems: masterItem.MasterItems{
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want: &LoginBonusReceiveResponse{
				UserLoginBonus: &userLoginBonus.UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				},
				MasterLoginBonus: &masterLoginBonus.MasterLoginBonus{
					Id:                      1,
					MasterLoginBonusEventId: 1,
					Name:                    "テストログインボーナス",
				},
				MasterLoginBonusEvent: &masterLoginBonusEvent.MasterLoginBonusEvent{
					Id:            1,
					Name:          "テストログインボーナスイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC),
					EndAt:         nil,
				},
				MasterLoginBonusItems: masterLoginBonusItem.MasterLoginBonusItems{
					{
						Id:                         2,
						MasterLoginBonusScheduleId: 2,
						MasterItemId:               1,
						Name:                       "テストログインボーナスアイテム",
						Count:                      1,
					},
				},
				MasterLoginBonusSchedule: &masterLoginBonusSchedule.MasterLoginBonusSchedule{
					Id:                 2,
					MasterLoginBonusId: 1,
					Step:               1,
					Name:               "ステップ1",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterLoginBonusMysqlRepository.Find",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
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
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterLoginBonusMysqlRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.getEvent",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getEvent: failed to s.masterLoginBonusEventMysqlRepository.FindByMasterLoginBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.getSchedule",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.getSchedule: failed to s.masterLoginBonusScheduleMysqlRepository.FindListByMasterLoginBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.masterLoginBonusItemMysqlRepository.FindListByMasterLoginBonusScheduleId",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterLoginBonusItemMysqlRepository.FindListByMasterLoginBonusScheduleId", errors.NewTestError()),
		},
		{
			name: "異常：s.userLoginBonusMysqlRepository.FindOrNil）",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         2,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userLoginBonusMysqlRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：already received",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         2,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("already received"),
		},
		{
			name: "異常：s.receive",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         2,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.receive: failed to s.itemService.Receive", errors.NewTestError()),
		},
		{
			name: "異常：s.update",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         2,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:WntR-PyhOJeDiE5jodeR",
							int64(1),
						).
						Return(
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
								},
								MasterItems: masterItem.MasterItems{
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.update: failed to s.userLoginBonusMysqlRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.update",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
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
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusId(
							nil,
							int64(1),
						).
						Return(
							masterLoginBonusSchedule.MasterLoginBonusSchedules{
								{
									Id:                 1,
									MasterLoginBonusId: 1,
									Step:               0,
									Name:               "ステップ0",
								},
								{
									Id:                 2,
									MasterLoginBonusId: 1,
									Step:               1,
									Name:               "ステップ1",
								},
								{
									Id:                 3,
									MasterLoginBonusId: 1,
									Step:               2,
									Name:               "ステップ2",
								},
							},
							nil,
						)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					m.EXPECT().
						FindListByMasterLoginBonusScheduleId(
							nil,
							int64(2),
						).
						Return(
							masterLoginBonusItem.MasterLoginBonusItems{
								{
									Id:                         2,
									MasterLoginBonusScheduleId: 2,
									MasterItemId:               1,
									Name:                       "テストログインボーナスアイテム",
									Count:                      1,
								},
							},
							nil,
						)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
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
							&userLoginBonus.UserLoginBonus{
								UserId:             "0:WntR-PyhOJeDiE5jodeR",
								MasterLoginBonusId: 1,
								ReceivedAt:         time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
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
								},
								MasterItems: masterItem.MasterItems{
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.update: failed to s.userLoginBonusMysqlRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：s.getEvent: outside the event period",
			fields: fields{
				masterLoginBonusMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusMysqlRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonus.MasterLoginBonus{
								Id:                      1,
								MasterLoginBonusEventId: 1,
								Name:                    "テストログインボーナス",
							},
							nil,
						)
					return m
				},
				masterLoginBonusEventMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventMysqlRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventMysqlRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterLoginBonusEvent.MasterLoginBonusEvent{
								Id:            1,
								Name:          "テストログインボーナスイベント",
								ResetHour:     9,
								IntervalHour:  24,
								RepeatSetting: true,
								StartAt:       time.Date(2023, 1, 3, 9, 0, 0, 0, time.UTC),
								EndAt:         nil,
							},
							nil,
						)
					return m
				},
				masterLoginBonusScheduleMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleMysqlRepository(ctrl)
					return m
				},
				masterLoginBonusItemMysqlRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemMysqlRepository(ctrl)
					return m
				},
				userLoginBonusMysqlRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusMysqlRepository {
					m := userLoginBonus.NewMockUserLoginBonusMysqlRepository(ctrl)
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
				now: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
				req: &LoginBonusReceiveRequest{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewError("failed to s.getEvent: outside the event period"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &loginBonusService{
				itemService:                             tt.fields.itemService(ctrl),
				userLoginBonusMysqlRepository:           tt.fields.userLoginBonusMysqlRepository(ctrl),
				masterLoginBonusMysqlRepository:         tt.fields.masterLoginBonusMysqlRepository(ctrl),
				masterLoginBonusEventMysqlRepository:    tt.fields.masterLoginBonusEventMysqlRepository(ctrl),
				masterLoginBonusItemMysqlRepository:     tt.fields.masterLoginBonusItemMysqlRepository(ctrl),
				masterLoginBonusScheduleMysqlRepository: tt.fields.masterLoginBonusScheduleMysqlRepository(ctrl),
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
