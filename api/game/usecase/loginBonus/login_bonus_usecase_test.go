package loginBonus

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	loginBonusServer "github.com/game-core/gocrafter/api/game/presentation/server/loginBonus"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/times"
	loginBonusService "github.com/game-core/gocrafter/pkg/domain/model/loginBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusEvent"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusItem"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonusSchedule"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/userLoginBonus"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

func TestLoginBonusUsecase_NewLoginBonusUsecase(t *testing.T) {
	type args struct {
		loginBonusService  loginBonusService.LoginBonusService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want LoginBonusUsecase
	}{
		{
			name: "正常",
			args: args{
				loginBonusService:  nil,
				transactionService: nil,
			},
			want: &loginBonusUsecase{
				loginBonusService:  nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLoginBonusUsecase(tt.args.loginBonusService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginBonusUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginBonusUsecase_GetUser(t *testing.T) {
	type fields struct {
		loginBonusService  func(ctrl *gomock.Controller) loginBonusService.LoginBonusService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *loginBonusServer.LoginBonusGetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *loginBonusServer.LoginBonusGetUserResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
					m.EXPECT().
						GetUser(
							gomock.Any(),
							&loginBonusService.LoginBonusGetUserRequest{
								UserId: "0:test",
							},
						).
						Return(
							&loginBonusService.LoginBonusGetUserResponse{
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
				req: &loginBonusServer.LoginBonusGetUserRequest{
					UserId: "0:test",
				},
			},
			want: &loginBonusServer.LoginBonusGetUserResponse{
				UserLoginBonuses: []*loginBonusServer.UserLoginBonus{
					{
						UserId:             "0:test",
						MasterLoginBonusId: 1,
						ReceivedAt:         times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC))),
					},
					{
						UserId:             "0:test",
						MasterLoginBonusId: 2,
						ReceivedAt:         times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC))),
					},
					{
						UserId:             "0:test",
						MasterLoginBonusId: 3,
						ReceivedAt:         times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC))),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.loginBonusService.GetUser",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
					m.EXPECT().
						GetUser(
							gomock.Any(),
							&loginBonusService.LoginBonusGetUserRequest{
								UserId: "0:test",
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
				req: &loginBonusServer.LoginBonusGetUserRequest{
					UserId: "0:test",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.loginBonusService.GetUser", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &loginBonusUsecase{
				loginBonusService:  tt.fields.loginBonusService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.GetUser(tt.args.ctx, tt.args.req)
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

func TestLoginBonusUsecase_GetMaster(t *testing.T) {
	type fields struct {
		loginBonusService  func(ctrl *gomock.Controller) loginBonusService.LoginBonusService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *loginBonusServer.LoginBonusGetMasterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *loginBonusServer.LoginBonusGetMasterResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
					m.EXPECT().
						GetMaster(
							gomock.Any(),
							&loginBonusService.LoginBonusGetMasterRequest{
								MasterLoginBonusId: 1,
							},
						).
						Return(
							&loginBonusService.LoginBonusGetMasterResponse{
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
				req: &loginBonusServer.LoginBonusGetMasterRequest{
					UserId:             "0:test",
					MasterLoginBonusId: 1,
				},
			},
			want: &loginBonusServer.LoginBonusGetMasterResponse{
				MasterLoginBonus: &loginBonusServer.MasterLoginBonus{
					Id:                      1,
					MasterLoginBonusEventId: 1,
					Name:                    "テストログインボーナス",
				},
				MasterLoginBonusEvent: &loginBonusServer.MasterLoginBonusEvent{
					Id:            1,
					Name:          "テストログインボーナスイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC))),
					EndAt:         nil,
				},
				MasterLoginBonusItems: []*loginBonusServer.MasterLoginBonusItem{
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
				MasterLoginBonusSchedules: []*loginBonusServer.MasterLoginBonusSchedule{
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
			name: "異常：s.loginBonusService.GetMaster",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
					m.EXPECT().
						GetMaster(
							gomock.Any(),
							&loginBonusService.LoginBonusGetMasterRequest{
								MasterLoginBonusId: 1,
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
				req: &loginBonusServer.LoginBonusGetMasterRequest{
					UserId:             "0:test",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.loginBonusService.GetMaster", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &loginBonusUsecase{
				loginBonusService:  tt.fields.loginBonusService(ctrl),
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

func TestLoginBonusUsecase_Receive(t *testing.T) {
	type fields struct {
		loginBonusService  func(ctrl *gomock.Controller) loginBonusService.LoginBonusService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *loginBonusServer.LoginBonusReceiveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *loginBonusServer.LoginBonusReceiveResponse
		wantErr error
	}{
		{
			name: "正常：受け取りできる",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
					m.EXPECT().
						Receive(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&loginBonusService.LoginBonusReceiveRequest{
								UserId:             "0:test",
								MasterLoginBonusId: 1,
							},
						).
						Return(
							&loginBonusService.LoginBonusReceiveResponse{
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
							nil,
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
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				req: &loginBonusServer.LoginBonusReceiveRequest{
					UserId:             "0:test",
					MasterLoginBonusId: 1,
				},
			},
			want: &loginBonusServer.LoginBonusReceiveResponse{
				UserLoginBonus: &loginBonusServer.UserLoginBonus{
					UserId:             "0:WntR-PyhOJeDiE5jodeR",
					MasterLoginBonusId: 1,
					ReceivedAt:         times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC))),
				},
				MasterLoginBonus: &loginBonusServer.MasterLoginBonus{
					Id:                      1,
					MasterLoginBonusEventId: 1,
					Name:                    "テストログインボーナス",
				},
				MasterLoginBonusEvent: &loginBonusServer.MasterLoginBonusEvent{
					Id:            1,
					Name:          "テストログインボーナスイベント",
					ResetHour:     9,
					IntervalHour:  24,
					RepeatSetting: true,
					StartAt:       times.TimeToPb(times.TimeToPointer(time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC))),
					EndAt:         nil,
				},
				MasterLoginBonusItems: []*loginBonusServer.MasterLoginBonusItem{
					{
						Id:                         2,
						MasterLoginBonusScheduleId: 2,
						MasterItemId:               1,
						Name:                       "テストログインボーナスアイテム",
						Count:                      1,
					},
				},
				MasterLoginBonusSchedule: &loginBonusServer.MasterLoginBonusSchedule{
					Id:                 2,
					MasterLoginBonusId: 1,
					Step:               1,
					Name:               "ステップ1",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.UserBegin",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
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
				req: &loginBonusServer.LoginBonusReceiveRequest{
					UserId:             "0:test",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.UserBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.loginBonusService.Receive",
			fields: fields{
				loginBonusService: func(ctrl *gomock.Controller) loginBonusService.LoginBonusService {
					m := loginBonusService.NewMockLoginBonusService(ctrl)
					m.EXPECT().
						Receive(
							gomock.Any(),
							gomock.Any(),
							gomock.Any(),
							&loginBonusService.LoginBonusReceiveRequest{
								UserId:             "0:test",
								MasterLoginBonusId: 1,
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
				req: &loginBonusServer.LoginBonusReceiveRequest{
					UserId:             "0:test",
					MasterLoginBonusId: 1,
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.loginBonusService.Receive", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &loginBonusUsecase{
				loginBonusService:  tt.fields.loginBonusService(ctrl),
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
