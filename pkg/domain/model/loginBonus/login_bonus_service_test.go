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
		itemService                        item.ItemService
		userLoginBonusRepository           userLoginBonus.UserLoginBonusRepository
		masterLoginBonusRepository         masterLoginBonus.MasterLoginBonusRepository
		masterLoginBonusEventRepository    masterLoginBonusEvent.MasterLoginBonusEventRepository
		masterLoginBonusItemRepository     masterLoginBonusItem.MasterLoginBonusItemRepository
		masterLoginBonusScheduleRepository masterLoginBonusSchedule.MasterLoginBonusScheduleRepository
	}
	tests := []struct {
		name string
		args args
		want LoginBonusService
	}{
		{
			name: "正常",
			args: args{
				itemService:                        nil,
				userLoginBonusRepository:           nil,
				masterLoginBonusRepository:         nil,
				masterLoginBonusEventRepository:    nil,
				masterLoginBonusItemRepository:     nil,
				masterLoginBonusScheduleRepository: nil,
			},
			want: &loginBonusService{
				itemService:                        nil,
				userLoginBonusRepository:           nil,
				masterLoginBonusRepository:         nil,
				masterLoginBonusEventRepository:    nil,
				masterLoginBonusItemRepository:     nil,
				masterLoginBonusScheduleRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLoginBonusService(
				tt.args.itemService,
				tt.args.userLoginBonusRepository,
				tt.args.masterLoginBonusRepository,
				tt.args.masterLoginBonusEventRepository,
				tt.args.masterLoginBonusItemRepository,
				tt.args.masterLoginBonusScheduleRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginBonusService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewItemService_Receive(t *testing.T) {
	type fields struct {
		itemService                        func(ctrl *gomock.Controller) item.ItemService
		userLoginBonusRepository           func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository
		masterLoginBonusRepository         func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository
		masterLoginBonusEventRepository    func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository
		masterLoginBonusItemRepository     func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository
		masterLoginBonusScheduleRepository func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository
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
			name: "正常：作成できる場合（既に取得したことがあるアイテム）",
			fields: fields{
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
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
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
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
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
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
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
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
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
			name: "異常：s.masterLoginBonusRepository.Find",
			fields: fields{
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
					return m
				},
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
					return m
				},
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
					return m
				},
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.masterLoginBonusRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.getEvent",
			fields: fields{
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
					return m
				},
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
					return m
				},
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.getEvent: failed to s.masterLoginBonusEventRepository.FindByMasterLoginBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.getSchedule",
			fields: fields{
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
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
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
					return m
				},
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.getSchedule: failed to s.masterLoginBonusScheduleRepository.FindListByMasterLoginBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.masterLoginBonusItemRepository.FindListByMasterLoginBonusScheduleId",
			fields: fields{
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
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
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
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
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.masterLoginBonusItemRepository.FindListByMasterLoginBonusScheduleId", errors.NewTestError()),
		},
		{
			name: "異常：s.userLoginBonusRepository.FindOrNil）",
			fields: fields{
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
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
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
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
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.userLoginBonusRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：already received",
			fields: fields{
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
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
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
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
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
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
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
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
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
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
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
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
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.update: failed to s.userLoginBonusRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.update",
			fields: fields{
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
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
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
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
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.update: failed to s.userLoginBonusRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：s.getEvent: outside the event period",
			fields: fields{
				masterLoginBonusRepository: func(ctrl *gomock.Controller) masterLoginBonus.MasterLoginBonusRepository {
					m := masterLoginBonus.NewMockMasterLoginBonusRepository(ctrl)
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
				masterLoginBonusEventRepository: func(ctrl *gomock.Controller) masterLoginBonusEvent.MasterLoginBonusEventRepository {
					m := masterLoginBonusEvent.NewMockMasterLoginBonusEventRepository(ctrl)
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
				masterLoginBonusScheduleRepository: func(ctrl *gomock.Controller) masterLoginBonusSchedule.MasterLoginBonusScheduleRepository {
					m := masterLoginBonusSchedule.NewMockMasterLoginBonusScheduleRepository(ctrl)
					return m
				},
				masterLoginBonusItemRepository: func(ctrl *gomock.Controller) masterLoginBonusItem.MasterLoginBonusItemRepository {
					m := masterLoginBonusItem.NewMockMasterLoginBonusItemRepository(ctrl)
					return m
				},
				userLoginBonusRepository: func(ctrl *gomock.Controller) userLoginBonus.UserLoginBonusRepository {
					m := userLoginBonus.NewMockUserLoginBonusRepository(ctrl)
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
				itemService:                        tt.fields.itemService(ctrl),
				userLoginBonusRepository:           tt.fields.userLoginBonusRepository(ctrl),
				masterLoginBonusRepository:         tt.fields.masterLoginBonusRepository(ctrl),
				masterLoginBonusEventRepository:    tt.fields.masterLoginBonusEventRepository(ctrl),
				masterLoginBonusItemRepository:     tt.fields.masterLoginBonusItemRepository(ctrl),
				masterLoginBonusScheduleRepository: tt.fields.masterLoginBonusScheduleRepository(ctrl),
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
