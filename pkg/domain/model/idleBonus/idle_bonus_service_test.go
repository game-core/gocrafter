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
		itemService                       item.ItemService
		userIdleBonusRepository           userIdleBonus.UserIdleBonusRepository
		masterIdleBonusRepository         masterIdleBonus.MasterIdleBonusRepository
		masterIdleBonusEventRepository    masterIdleBonusEvent.MasterIdleBonusEventRepository
		masterIdleBonusItemRepository     masterIdleBonusItem.MasterIdleBonusItemRepository
		masterIdleBonusScheduleRepository masterIdleBonusSchedule.MasterIdleBonusScheduleRepository
	}
	tests := []struct {
		name string
		args args
		want IdleBonusService
	}{
		{
			name: "正常",
			args: args{
				itemService:                       nil,
				userIdleBonusRepository:           nil,
				masterIdleBonusRepository:         nil,
				masterIdleBonusEventRepository:    nil,
				masterIdleBonusItemRepository:     nil,
				masterIdleBonusScheduleRepository: nil,
			},
			want: &idleBonusService{
				itemService:                       nil,
				userIdleBonusRepository:           nil,
				masterIdleBonusRepository:         nil,
				masterIdleBonusEventRepository:    nil,
				masterIdleBonusItemRepository:     nil,
				masterIdleBonusScheduleRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIdleBonusService(
				tt.args.itemService,
				tt.args.userIdleBonusRepository,
				tt.args.masterIdleBonusRepository,
				tt.args.masterIdleBonusEventRepository,
				tt.args.masterIdleBonusItemRepository,
				tt.args.masterIdleBonusScheduleRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdleBonusService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewItemService_Receive(t *testing.T) {
	type fields struct {
		itemService                       func(ctrl *gomock.Controller) item.ItemService
		userIdleBonusRepository           func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository
		masterIdleBonusRepository         func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository
		masterIdleBonusEventRepository    func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository
		masterIdleBonusItemRepository     func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository
		masterIdleBonusScheduleRepository func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository
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
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
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
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
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
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
					return m
				},
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
					return m
				},
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			name: "異常：s.masterIdleBonusRepository.Find",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
					return m
				},
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
					return m
				},
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
					return m
				},
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.masterIdleBonusRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：failed to s.getEvent: failed to s.masterIdleBonusEventRepository.FindByMasterIdleBonusId",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
					return m
				},
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
					return m
				},
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.getEvent: failed to s.masterIdleBonusEventRepository.FindByMasterIdleBonusId", errors.NewTestError()),
		},
		{
			name: "異常：failed to s.getEvent: outside the event period",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
					return m
				},
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
					return m
				},
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			name: "異常：s.userIdleBonusRepository.FindOrNil",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
					return m
				},
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
					return m
				},
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.userIdleBonusRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：s.masterIdleBonusScheduleRepository.FindListByMasterIdleBonusId",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
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
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
					return m
				},
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.getSchedules: failed to s.masterIdleBonusScheduleRepository.FindListByMasterIdleBonusId", errors.NewTestError()),
		},
		{
			name: "異常：s.masterIdleBonusItemRepository.FindListByMasterIdleBonusScheduleId",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
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
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
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
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.getItems: failed to s.masterIdleBonusItemRepository.FindListByMasterIdleBonusScheduleId", errors.NewTestError()),
		},
		{
			name: "異常：s.itemService.Receive",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
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
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
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
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			name: "異常：s.userIdleBonusRepository.Update",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
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
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
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
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.update: failed to s.userIdleBonusRepository.Update", errors.NewTestError()),
		},
		{
			name: "異常：s.userIdleBonusRepository.Create",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
					return m
				},
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
					return m
				},
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
			wantErr: errors.NewMethodError("s.userIdleBonusRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：masterIdleBonusSchedules.GetStep",
			fields: fields{
				masterIdleBonusRepository: func(ctrl *gomock.Controller) masterIdleBonus.MasterIdleBonusRepository {
					m := masterIdleBonus.NewMockMasterIdleBonusRepository(ctrl)
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
				masterIdleBonusEventRepository: func(ctrl *gomock.Controller) masterIdleBonusEvent.MasterIdleBonusEventRepository {
					m := masterIdleBonusEvent.NewMockMasterIdleBonusEventRepository(ctrl)
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
				masterIdleBonusScheduleRepository: func(ctrl *gomock.Controller) masterIdleBonusSchedule.MasterIdleBonusScheduleRepository {
					m := masterIdleBonusSchedule.NewMockMasterIdleBonusScheduleRepository(ctrl)
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
				masterIdleBonusItemRepository: func(ctrl *gomock.Controller) masterIdleBonusItem.MasterIdleBonusItemRepository {
					m := masterIdleBonusItem.NewMockMasterIdleBonusItemRepository(ctrl)
					return m
				},
				userIdleBonusRepository: func(ctrl *gomock.Controller) userIdleBonus.UserIdleBonusRepository {
					m := userIdleBonus.NewMockUserIdleBonusRepository(ctrl)
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
				itemService:                       tt.fields.itemService(ctrl),
				userIdleBonusRepository:           tt.fields.userIdleBonusRepository(ctrl),
				masterIdleBonusRepository:         tt.fields.masterIdleBonusRepository(ctrl),
				masterIdleBonusEventRepository:    tt.fields.masterIdleBonusEventRepository(ctrl),
				masterIdleBonusItemRepository:     tt.fields.masterIdleBonusItemRepository(ctrl),
				masterIdleBonusScheduleRepository: tt.fields.masterIdleBonusScheduleRepository(ctrl),
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
