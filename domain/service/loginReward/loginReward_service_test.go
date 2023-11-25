package loginReward

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	request "github.com/game-core/gocrafter/api/presentation/request/loginReward"
	itemResponse "github.com/game-core/gocrafter/api/presentation/response/item"
	response "github.com/game-core/gocrafter/api/presentation/response/loginReward"
	"github.com/game-core/gocrafter/config/pointer"
	masterEventEntity "github.com/game-core/gocrafter/domain/entity/master/event"
	masterLoginRewardEntity "github.com/game-core/gocrafter/domain/entity/master/loginReward"
	userLoginRewarEntity "github.com/game-core/gocrafter/domain/entity/user/loginReward"
	masterLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/master/loginReward"
	userRepository "github.com/game-core/gocrafter/domain/repository/user"
	userLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/user/loginReward"
	eventService "github.com/game-core/gocrafter/domain/service/event"
	itemService "github.com/game-core/gocrafter/domain/service/item"
)

func TestExampleService_GetLoginRewardModel(t *testing.T) {
	type fields struct {
		loginRewardModelRepository  func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository
		loginRewardRewardRepository func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository
		eventService                func(ctrl *gomock.Controller) eventService.EventService
	}
	type args struct {
		req *request.GetLoginRewardModel
		now time.Time
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		want    *response.GetLoginRewardModel
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardRewards{
								{
									ID:                   1,
									LoginRewardModelName: "loginReward",
									Name:                 "reward1",
									StepNumber:           0,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   2,
									LoginRewardModelName: "loginReward",
									Name:                 "reward2",
									StepNumber:           1,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					m.EXPECT().
						GetEventToEntity(
							"event",
						).
						Return(
							&masterEventEntity.Event{
								ID:            1,
								Name:          "event",
								ResetHour:     9,
								RepeatSetting: false,
								RepeatStartAt: nil,
								StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
								EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
								CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				req: &request.GetLoginRewardModel{
					LoginRewardModelName: "loginReward",
				},
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want: &response.GetLoginRewardModel{
				Status: 200,
				Item: response.LoginRewardModel{
					ID:   1,
					Name: "loginReward",
					Event: response.Event{
						ID:            1,
						Name:          "event",
						ResetHour:     9,
						RepeatSetting: false,
						RepeatStartAt: nil,
						StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
						EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
					},
					LoginRewardRewards: response.LoginRewardRewards{
						{
							ID:         1,
							Name:       "reward1",
							StepNumber: 0,
							Items: response.Items{
								{
									Name:  "item1",
									Count: 1,
								},
								{
									Name:  "item2",
									Count: 2,
								},
							},
						},
						{
							ID:         2,
							StepNumber: 1,
							Name:       "reward2",
							Items: response.Items{
								{
									Name:  "item1",
									Count: 1,
								},
								{
									Name:  "item2",
									Count: 2,
								},
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー",
			fields: fields{
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardRewards{
								{
									ID:                   1,
									LoginRewardModelName: "loginReward",
									Name:                 "reward1",
									StepNumber:           0,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   2,
									LoginRewardModelName: "loginReward",
									Name:                 "reward2",
									StepNumber:           1,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					m.EXPECT().
						GetEventToEntity(
							"event",
						).
						Return(
							nil,
							errors.New("getLoginRewardModelAndRewardsAndEvent"),
						)
					return m
				},
			},
			args: args{
				req: &request.GetLoginRewardModel{
					LoginRewardModelName: "loginReward",
				},
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: errors.New("getLoginRewardModelAndRewardsAndEvent"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &loginRewardService{
				loginRewardModelRepository:  tt.fields.loginRewardModelRepository(ctrl),
				loginRewardRewardRepository: tt.fields.loginRewardRewardRepository(ctrl),
				eventService:                tt.fields.eventService(ctrl),
			}

			got, err := s.GetLoginRewardModel(tt.args.req, tt.args.now)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetLoginRewardModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLoginRewardModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleService_ReceiveLoginReward(t *testing.T) {
	type fields struct {
		transactionRepository       func(ctrl *gomock.Controller) userRepository.TransactionRepository
		loginRewardModelRepository  func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository
		loginRewardRewardRepository func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository
		loginRewardStatusRepository func(ctrl *gomock.Controller) userLoginRewardRepository.LoginRewardStatusRepository
		eventService                func(ctrl *gomock.Controller) eventService.EventService
		itemService                 func(ctrl *gomock.Controller) itemService.ItemService
	}
	type args struct {
		req *request.ReceiveLoginReward
		now time.Time
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		want    *response.ReceiveLoginReward
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin(
							"SHARD_1",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Commit(
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardRewards{
								{
									ID:                   1,
									LoginRewardModelName: "loginReward",
									Name:                 "reward1",
									StepNumber:           0,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   2,
									LoginRewardModelName: "loginReward",
									Name:                 "reward2",
									StepNumber:           1,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   3,
									LoginRewardModelName: "loginReward",
									Name:                 "reward3",
									StepNumber:           2,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					m.EXPECT().
						GetEventToEntity(
							"event",
						).
						Return(
							&masterEventEntity.Event{
								ID:            1,
								Name:          "event",
								ResetHour:     9,
								RepeatSetting: true,
								RepeatStartAt: pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
								StartAt:       nil,
								EndAt:         nil,
								CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardStatusRepository: func(ctrl *gomock.Controller) userLoginRewardRepository.LoginRewardStatusRepository {
					m := userLoginRewardRepository.NewMockLoginRewardStatusRepository(ctrl)
					m.EXPECT().
						FindOrNilByLoginRewardModelName(
							"loginReward",
							"SHARD_1",
						).
						Return(
							&userLoginRewarEntity.LoginRewardStatus{
								ID:                   1,
								ShardKey:             "SHARD_1",
								AccountID:            1,
								LoginRewardModelName: "loginReward",
								LastReceivedAt:       time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC),
								CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Save(
							gomock.Any(),
							"SHARD_1",
							nil,
						).
						Return(
							&userLoginRewarEntity.LoginRewardStatus{
								ID:                   1,
								ShardKey:             "SHARD_1",
								AccountID:            1,
								LoginRewardModelName: "loginReward",
								LastReceivedAt:       time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
								CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) itemService.ItemService {
					m := itemService.NewMockItemService(ctrl)
					m.EXPECT().
						ReceiveItemInBox(
							gomock.Any(),
						).
						Return(
							&itemResponse.ReceiveItemInBox{
								Status: 200,
								Items: itemResponse.Items{
									{
										ID:     1,
										Name:   "item1",
										Detail: "detail1",
										Count:  1,
									},
									{
										ID:     2,
										Name:   "item2",
										Detail: "detail2",
										Count:  2,
									},
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				req: &request.ReceiveLoginReward{
					ShardKey:             "SHARD_1",
					AccountID:            1,
					UUID:                 "uuid",
					LoginRewardModelName: "loginReward",
				},
				now: time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
			},
			want: &response.ReceiveLoginReward{
				Status: 200,
				Item: response.LoginRewardStatus{
					ID: 1,
					LoginRewardModel: response.LoginRewardModel{
						ID:   1,
						Name: "loginReward",
						Event: response.Event{
							ID:            1,
							Name:          "event",
							ResetHour:     9,
							RepeatSetting: true,
							RepeatStartAt: pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
							StartAt:       nil,
							EndAt:         nil,
						},
					},
					Items: response.Items{
						{
							Name:  "item1",
							Count: 1,
						},
						{
							Name:  "item2",
							Count: 2,
						},
					},
					LastReceivedAt: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：トランザクションエラー",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin(
							"SHARD_1",
						).
						Return(
							nil,
							errors.New("transactionRepository.Begin"),
						)
					return m
				},
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					return m
				},
				loginRewardStatusRepository: func(ctrl *gomock.Controller) userLoginRewardRepository.LoginRewardStatusRepository {
					m := userLoginRewardRepository.NewMockLoginRewardStatusRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) itemService.ItemService {
					m := itemService.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				req: &request.ReceiveLoginReward{
					ShardKey:             "SHARD_1",
					AccountID:            1,
					UUID:                 "uuid",
					LoginRewardModelName: "loginReward",
				},
				now: time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: errors.New("transactionRepository.Begin"),
		},
		{
			name: "異常：エラー",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin(
							"SHARD_1",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Rollback(
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							nil,
							errors.New("s.getLoginRewardModelAndRewardsAndEvent"),
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					return m
				},
				loginRewardStatusRepository: func(ctrl *gomock.Controller) userLoginRewardRepository.LoginRewardStatusRepository {
					m := userLoginRewardRepository.NewMockLoginRewardStatusRepository(ctrl)
					return m
				},
				itemService: func(ctrl *gomock.Controller) itemService.ItemService {
					m := itemService.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				req: &request.ReceiveLoginReward{
					ShardKey:             "SHARD_1",
					AccountID:            1,
					UUID:                 "uuid",
					LoginRewardModelName: "loginReward",
				},
				now: time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: errors.New("s.getLoginRewardModelAndRewardsAndEvent"),
		},
		{
			name: "異常：エラー",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin(
							"SHARD_1",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Rollback(
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardRewards{
								{
									ID:                   1,
									LoginRewardModelName: "loginReward",
									Name:                 "reward1",
									StepNumber:           0,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   2,
									LoginRewardModelName: "loginReward",
									Name:                 "reward2",
									StepNumber:           1,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   3,
									LoginRewardModelName: "loginReward",
									Name:                 "reward3",
									StepNumber:           2,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					m.EXPECT().
						GetEventToEntity(
							"event",
						).
						Return(
							&masterEventEntity.Event{
								ID:            1,
								Name:          "event",
								ResetHour:     9,
								RepeatSetting: true,
								RepeatStartAt: pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
								StartAt:       nil,
								EndAt:         nil,
								CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardStatusRepository: func(ctrl *gomock.Controller) userLoginRewardRepository.LoginRewardStatusRepository {
					m := userLoginRewardRepository.NewMockLoginRewardStatusRepository(ctrl)
					m.EXPECT().
						FindOrNilByLoginRewardModelName(
							"loginReward",
							"SHARD_1",
						).
						Return(
							nil,
							errors.New("loginRewardStatusRepository.FindOrNilByLoginRewardModelName"),
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) itemService.ItemService {
					m := itemService.NewMockItemService(ctrl)
					return m
				},
			},
			args: args{
				req: &request.ReceiveLoginReward{
					ShardKey:             "SHARD_1",
					AccountID:            1,
					UUID:                 "uuid",
					LoginRewardModelName: "loginReward",
				},
				now: time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: errors.New("loginRewardStatusRepository.FindOrNilByLoginRewardModelName"),
		},
		{
			name: "異常：エラー",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin(
							"SHARD_1",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Rollback(
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardRewards{
								{
									ID:                   1,
									LoginRewardModelName: "loginReward",
									Name:                 "reward1",
									StepNumber:           0,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   2,
									LoginRewardModelName: "loginReward",
									Name:                 "reward2",
									StepNumber:           1,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   3,
									LoginRewardModelName: "loginReward",
									Name:                 "reward3",
									StepNumber:           2,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					m.EXPECT().
						GetEventToEntity(
							"event",
						).
						Return(
							&masterEventEntity.Event{
								ID:            1,
								Name:          "event",
								ResetHour:     9,
								RepeatSetting: true,
								RepeatStartAt: pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
								StartAt:       nil,
								EndAt:         nil,
								CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardStatusRepository: func(ctrl *gomock.Controller) userLoginRewardRepository.LoginRewardStatusRepository {
					m := userLoginRewardRepository.NewMockLoginRewardStatusRepository(ctrl)
					m.EXPECT().
						FindOrNilByLoginRewardModelName(
							"loginReward",
							"SHARD_1",
						).
						Return(
							&userLoginRewarEntity.LoginRewardStatus{
								ID:                   1,
								ShardKey:             "SHARD_1",
								AccountID:            1,
								LoginRewardModelName: "loginReward",
								LastReceivedAt:       time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC),
								CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) itemService.ItemService {
					m := itemService.NewMockItemService(ctrl)
					m.EXPECT().
						ReceiveItemInBox(
							gomock.Any(),
						).
						Return(
							nil,
							errors.New("itemService.ReceiveItemInBox"),
						)
					return m
				},
			},
			args: args{
				req: &request.ReceiveLoginReward{
					ShardKey:             "SHARD_1",
					AccountID:            1,
					UUID:                 "uuid",
					LoginRewardModelName: "loginReward",
				},
				now: time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: errors.New("itemService.ReceiveItemInBox"),
		},
		{
			name: "異常：エラー",
			fields: fields{
				transactionRepository: func(ctrl *gomock.Controller) userRepository.TransactionRepository {
					m := userRepository.NewMockTransactionRepository(ctrl)
					m.EXPECT().
						Begin(
							"SHARD_1",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Rollback(
							gomock.Any(),
						).
						Return(
							nil,
						)
					return m
				},
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardRewards{
								{
									ID:                   1,
									LoginRewardModelName: "loginReward",
									Name:                 "reward1",
									StepNumber:           0,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   2,
									LoginRewardModelName: "loginReward",
									Name:                 "reward2",
									StepNumber:           1,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   3,
									LoginRewardModelName: "loginReward",
									Name:                 "reward3",
									StepNumber:           2,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					m.EXPECT().
						GetEventToEntity(
							"event",
						).
						Return(
							&masterEventEntity.Event{
								ID:            1,
								Name:          "event",
								ResetHour:     9,
								RepeatSetting: true,
								RepeatStartAt: pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
								StartAt:       nil,
								EndAt:         nil,
								CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardStatusRepository: func(ctrl *gomock.Controller) userLoginRewardRepository.LoginRewardStatusRepository {
					m := userLoginRewardRepository.NewMockLoginRewardStatusRepository(ctrl)
					m.EXPECT().
						FindOrNilByLoginRewardModelName(
							"loginReward",
							"SHARD_1",
						).
						Return(
							&userLoginRewarEntity.LoginRewardStatus{
								ID:                   1,
								ShardKey:             "SHARD_1",
								AccountID:            1,
								LoginRewardModelName: "loginReward",
								LastReceivedAt:       time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC),
								CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Save(
							gomock.Any(),
							"SHARD_1",
							nil,
						).
						Return(
							nil,
							errors.New("loginRewardStatusRepository.Save"),
						)
					return m
				},
				itemService: func(ctrl *gomock.Controller) itemService.ItemService {
					m := itemService.NewMockItemService(ctrl)
					m.EXPECT().
						ReceiveItemInBox(
							gomock.Any(),
						).
						Return(
							&itemResponse.ReceiveItemInBox{
								Status: 200,
								Items: itemResponse.Items{
									{
										ID:     1,
										Name:   "item1",
										Detail: "detail1",
										Count:  1,
									},
									{
										ID:     2,
										Name:   "item2",
										Detail: "detail2",
										Count:  2,
									},
								},
							},
							nil,
						)
					return m
				},
			},
			args: args{
				req: &request.ReceiveLoginReward{
					ShardKey:             "SHARD_1",
					AccountID:            1,
					UUID:                 "uuid",
					LoginRewardModelName: "loginReward",
				},
				now: time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: errors.New("loginRewardStatusRepository.Save"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &loginRewardService{
				transactionRepository:       tt.fields.transactionRepository(ctrl),
				loginRewardModelRepository:  tt.fields.loginRewardModelRepository(ctrl),
				loginRewardRewardRepository: tt.fields.loginRewardRewardRepository(ctrl),
				loginRewardStatusRepository: tt.fields.loginRewardStatusRepository(ctrl),
				eventService:                tt.fields.eventService(ctrl),
				itemService:                 tt.fields.itemService(ctrl),
			}

			got, err := s.ReceiveLoginReward(tt.args.req, tt.args.now)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("ReceiveLoginReward() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReceiveLoginRewar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExampleService_getLoginRewardModelAndRewardsAndEvent(t *testing.T) {
	type fields struct {
		loginRewardModelRepository  func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository
		loginRewardRewardRepository func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository
		eventService                func(ctrl *gomock.Controller) eventService.EventService
	}
	type args struct {
		loginRewardModelName string
		now                  time.Time
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		want1   *masterLoginRewardEntity.LoginRewardModel
		want2   *masterLoginRewardEntity.LoginRewardRewards
		want3   *masterEventEntity.Event
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardRewards{
								{
									ID:                   1,
									LoginRewardModelName: "loginReward",
									Name:                 "reward1",
									StepNumber:           0,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   2,
									LoginRewardModelName: "loginReward",
									Name:                 "reward2",
									StepNumber:           1,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					m.EXPECT().
						GetEventToEntity(
							"event",
						).
						Return(
							&masterEventEntity.Event{
								ID:            1,
								Name:          "event",
								ResetHour:     9,
								RepeatSetting: false,
								RepeatStartAt: nil,
								StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
								EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
								CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				loginRewardModelName: "loginReward",
				now:                  time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want1: &masterLoginRewardEntity.LoginRewardModel{
				ID:        1,
				Name:      "loginReward",
				EventName: "event",
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want2: &masterLoginRewardEntity.LoginRewardRewards{
				{
					ID:                   1,
					LoginRewardModelName: "loginReward",
					Name:                 "reward1",
					StepNumber:           0,
					Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
					CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:                   2,
					LoginRewardModelName: "loginReward",
					Name:                 "reward2",
					StepNumber:           1,
					Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
					CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want3: &masterEventEntity.Event{
				ID:            1,
				Name:          "event",
				ResetHour:     9,
				RepeatSetting: false,
				RepeatStartAt: nil,
				StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)),
				EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
				CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "異常：エラー",
			fields: fields{
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							nil,
							errors.New("loginRewardModelRepository.FindByName"),
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					return m
				},
			},
			args: args{
				loginRewardModelName: "loginReward",
				now:                  time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want1:   nil,
			want2:   nil,
			want3:   nil,
			wantErr: errors.New("loginRewardModelRepository.FindByName"),
		},
		{
			name: "異常：エラー",
			fields: fields{
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							nil,
							errors.New("loginRewardRewardRepository.ListByLoginRewardModelName"),
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					return m
				},
			},
			args: args{
				loginRewardModelName: "loginReward",
				now:                  time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want1:   nil,
			want2:   nil,
			want3:   nil,
			wantErr: errors.New("loginRewardRewardRepository.ListByLoginRewardModelName"),
		},
		{
			name: "異常：エラー",
			fields: fields{
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardRewards{
								{
									ID:                   1,
									LoginRewardModelName: "loginReward",
									Name:                 "reward1",
									StepNumber:           0,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   2,
									LoginRewardModelName: "loginReward",
									Name:                 "reward2",
									StepNumber:           1,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					m.EXPECT().
						GetEventToEntity(
							"event",
						).
						Return(
							nil,
							errors.New("eventService.GetEventToEntity"),
						)
					return m
				},
			},
			args: args{
				loginRewardModelName: "loginReward",
				now:                  time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want1:   nil,
			want2:   nil,
			want3:   nil,
			wantErr: errors.New("eventService.GetEventToEntity"),
		},
		{
			name: "異常：エラー",
			fields: fields{
				loginRewardModelRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardModelRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardModelRepository(ctrl)
					m.EXPECT().
						FindByName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardModel{
								ID:        1,
								Name:      "loginReward",
								EventName: "event",
								CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
				loginRewardRewardRepository: func(ctrl *gomock.Controller) masterLoginRewardRepository.LoginRewardRewardRepository {
					m := masterLoginRewardRepository.NewMockLoginRewardRewardRepository(ctrl)
					m.EXPECT().
						ListByLoginRewardModelName(
							"loginReward",
						).
						Return(
							&masterLoginRewardEntity.LoginRewardRewards{
								{
									ID:                   1,
									LoginRewardModelName: "loginReward",
									Name:                 "reward1",
									StepNumber:           0,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
								{
									ID:                   2,
									LoginRewardModelName: "loginReward",
									Name:                 "reward2",
									StepNumber:           1,
									Items:                "[{\"name\":\"item1\",\"count\":1},{\"name\":\"item2\",\"count\":2}]",
									CreatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
									UpdatedAt:            time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							nil,
						)
					return m
				},
				eventService: func(ctrl *gomock.Controller) eventService.EventService {
					m := eventService.NewMockEventService(ctrl)
					m.EXPECT().
						GetEventToEntity(
							"event",
						).
						Return(
							&masterEventEntity.Event{
								ID:            1,
								Name:          "event",
								ResetHour:     9,
								RepeatSetting: false,
								RepeatStartAt: nil,
								StartAt:       pointer.TimeToPointer(time.Date(2023, 1, 3, 9, 0, 0, 0, time.UTC)),
								EndAt:         pointer.TimeToPointer(time.Date(2023, 1, 31, 8, 59, 59, 0, time.UTC)),
								CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				loginRewardModelName: "loginReward",
				now:                  time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want1:   nil,
			want2:   nil,
			want3:   nil,
			wantErr: errors.New("outside the event period"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &loginRewardService{
				loginRewardModelRepository:  tt.fields.loginRewardModelRepository(ctrl),
				loginRewardRewardRepository: tt.fields.loginRewardRewardRepository(ctrl),
				eventService:                tt.fields.eventService(ctrl),
			}

			want1, want2, want3, err := s.getLoginRewardModelAndRewardsAndEvent(tt.args.loginRewardModelName, tt.args.now)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetLoginRewardModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(want1, tt.want1) {
				t.Errorf("getLoginRewardModelAndRewardsAndEvent() = %v, want %v", want1, tt.want1)
			}
			if !reflect.DeepEqual(want2, tt.want2) {
				t.Errorf("getLoginRewardModelAndRewardsAndEvent() = %v, want %v", want2, tt.want2)
			}
			if !reflect.DeepEqual(want1, tt.want1) {
				t.Errorf("getLoginRewardModelAndRewardsAndEvent() = %v, want %v", want3, tt.want3)
			}
		})
	}
}
