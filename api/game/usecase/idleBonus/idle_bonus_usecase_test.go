package idleBonus

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	idleBonusServer "github.com/game-core/gocrafter/api/game/presentation/server/idleBonus"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/times"
	idleBonusService "github.com/game-core/gocrafter/pkg/domain/model/idleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/masterIdleBonusSchedule"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus/userIdleBonus"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

func TestIdleBonusUsecase_NewIdleBonusUsecase(t *testing.T) {
	type args struct {
		idleBonusService   idleBonusService.IdleBonusService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want IdleBonusUsecase
	}{
		{
			name: "正常",
			args: args{
				idleBonusService:   nil,
				transactionService: nil,
			},
			want: &idleBonusUsecase{
				idleBonusService:   nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIdleBonusUsecase(tt.args.idleBonusService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdleBonusUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdleBonusUsecase_GetMaster(t *testing.T) {
	type fields struct {
		idleBonusService   func(ctrl *gomock.Controller) idleBonusService.IdleBonusService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *idleBonusServer.IdleBonusGetMasterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idleBonusServer.IdleBonusGetMasterResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				idleBonusService: func(ctrl *gomock.Controller) idleBonusService.IdleBonusService {
					m := idleBonusService.NewMockIdleBonusService(ctrl)
					m.EXPECT().
						GetMaster(
							gomock.Any(),
							&idleBonusService.IdleBonusGetMasterRequest{
								MasterIdleBonusId: 1,
							},
						).
						Return(
							&idleBonusService.IdleBonusGetMasterResponse{
								MasterIdleBonus: &masterIdleBonus.MasterIdleBonus{
									Id:                     1,
									MasterIdleBonusEventId: 1,
									Name:                   "テスト放置ボーナス",
								},
								MasterIdleBonusEvent: &masterIdleBonusEvent.MasterIdleBonusEvent{
									Id:            1,
									Name:          "テスト放置ボーナスイベント",
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
									{
										Id:                        3,
										MasterIdleBonusScheduleId: 3,
										MasterItemId:              1,
										Name:                      "テスト放置ボーナスアイテム3",
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
				req: &idleBonusServer.IdleBonusGetMasterRequest{
					UserId:            "0:test",
					MasterIdleBonusId: 1,
				},
			},
			want: &idleBonusServer.IdleBonusGetMasterResponse{
				MasterIdleBonus: &idleBonusServer.MasterIdleBonus{
					Id:                     1,
					MasterIdleBonusEventId: 1,
					Name:                   "テスト放置ボーナス",
				},
				MasterIdleBonusEvent: &idleBonusServer.MasterIdleBonusEvent{
					Id:            1,
					Name:          "テスト放置ボーナスイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC))),
					EndAt:         nil,
				},
				MasterIdleBonusItems: []*idleBonusServer.MasterIdleBonusItem{
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
					{
						Id:                        3,
						MasterIdleBonusScheduleId: 3,
						MasterItemId:              1,
						Name:                      "テスト放置ボーナスアイテム3",
						Count:                     1,
					},
				},
				MasterIdleBonusSchedules: []*idleBonusServer.MasterIdleBonusSchedule{
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
			name: "異常：s.idleBonusService.GetMaster",
			fields: fields{
				idleBonusService: func(ctrl *gomock.Controller) idleBonusService.IdleBonusService {
					m := idleBonusService.NewMockIdleBonusService(ctrl)
					m.EXPECT().
						GetMaster(
							gomock.Any(),
							&idleBonusService.IdleBonusGetMasterRequest{
								MasterIdleBonusId: 1,
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
				req: &idleBonusServer.IdleBonusGetMasterRequest{
					UserId:            "0:test",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.idleBonusService.GetMaster", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &idleBonusUsecase{
				idleBonusService:   tt.fields.idleBonusService(ctrl),
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

func TestIdleBonusUsecase_Receive(t *testing.T) {
	type fields struct {
		idleBonusService   func(ctrl *gomock.Controller) idleBonusService.IdleBonusService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *idleBonusServer.IdleBonusReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idleBonusServer.IdleBonusReceiveResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				idleBonusService: func(ctrl *gomock.Controller) idleBonusService.IdleBonusService {
					m := idleBonusService.NewMockIdleBonusService(ctrl)
					m.EXPECT().
						Receive(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&idleBonusService.IdleBonusReceiveRequest{
								UserId:            "0:WntR-PyhOJeDiE5jodeR",
								MasterIdleBonusId: 1,
							},
						).
						Return(
							&idleBonusService.IdleBonusReceiveResponse{
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
							nil,
						)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						UserBegin(
							gomock.Any(),
							"0:WntR-PyhOJeDiE5jodeR",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						UserEnd(
							gomock.Any(),
							gomock.Any(),
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &idleBonusServer.IdleBonusReceiveRequest{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
				},
			},
			want: &idleBonusServer.IdleBonusReceiveResponse{
				UserIdleBonus: &idleBonusServer.UserIdleBonus{
					UserId:            "0:WntR-PyhOJeDiE5jodeR",
					MasterIdleBonusId: 1,
					ReceivedAt:        times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC))),
				},
				MasterIdleBonus: &idleBonusServer.MasterIdleBonus{
					Id:                     1,
					MasterIdleBonusEventId: 1,
					Name:                   "テスト放置ボーナス",
				},
				MasterIdleBonusEvent: &idleBonusServer.MasterIdleBonusEvent{
					Id:            1,
					Name:          "テスト放置ボーナスイベント",
					ResetHour:     9,
					IntervalHour:  1,
					RepeatSetting: true,
					StartAt:       times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC))),
					EndAt:         nil,
				},
				MasterIdleBonusItems: []*idleBonusServer.MasterIdleBonusItem{
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
				MasterIdleBonusSchedules: []*idleBonusServer.MasterIdleBonusSchedule{
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
			name: "異常：s.transactionService.UserBegin",
			fields: fields{
				idleBonusService: func(ctrl *gomock.Controller) idleBonusService.IdleBonusService {
					m := idleBonusService.NewMockIdleBonusService(ctrl)
					return m
				},
				transactionService: func(ctrl *gomock.Controller) transactionService.TransactionService {
					m := transactionService.NewMockTransactionService(ctrl)
					m.EXPECT().
						UserBegin(
							gomock.Any(),
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
				req: &idleBonusServer.IdleBonusReceiveRequest{
					UserId:            "0:test",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.UserBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.idleBonusService.Receive",
			fields: fields{
				idleBonusService: func(ctrl *gomock.Controller) idleBonusService.IdleBonusService {
					m := idleBonusService.NewMockIdleBonusService(ctrl)
					m.EXPECT().
						Receive(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&idleBonusService.IdleBonusReceiveRequest{
								UserId:            "0:test",
								MasterIdleBonusId: 1,
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
						UserBegin(
							gomock.Any(),
							"0:test",
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						UserEnd(
							gomock.Any(),
							gomock.Any(),
							errors.NewTestError(),
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &idleBonusServer.IdleBonusReceiveRequest{
					UserId:            "0:test",
					MasterIdleBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.idleBonusService.Receive", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &idleBonusUsecase{
				idleBonusService:   tt.fields.idleBonusService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Receive(tt.args.ctx, tt.args.req)
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
