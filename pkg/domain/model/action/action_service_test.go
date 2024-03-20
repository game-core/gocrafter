package action

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/pointers"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionRun"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionStep"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionTrigger"
	"github.com/game-core/gocrafter/pkg/domain/model/action/userAction"
)

func TestNewActionService_NewActionService(t *testing.T) {
	type args struct {
		masterActionRepository        masterAction.MasterActionRepository
		masterActionRunRepository     masterActionRun.MasterActionRunRepository
		masterActionStepRepository    masterActionStep.MasterActionStepRepository
		masterActionTriggerRepository masterActionTrigger.MasterActionTriggerRepository
		userActionRepository          userAction.UserActionRepository
	}
	tests := []struct {
		name string
		args args
		want ActionService
	}{
		{
			name: "正常",
			args: args{
				masterActionRepository:        nil,
				masterActionRunRepository:     nil,
				masterActionStepRepository:    nil,
				masterActionTriggerRepository: nil,
				userActionRepository:          nil,
			},
			want: &actionService{
				masterActionRepository:        nil,
				masterActionRunRepository:     nil,
				masterActionStepRepository:    nil,
				masterActionTriggerRepository: nil,
				userActionRepository:          nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewActionService(
				tt.args.masterActionRepository,
				tt.args.masterActionRunRepository,
				tt.args.masterActionStepRepository,
				tt.args.masterActionTriggerRepository,
				tt.args.userActionRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActionService_GetMaster(t *testing.T) {
	type fields struct {
		masterActionRepository        func(ctrl *gomock.Controller) masterAction.MasterActionRepository
		masterActionRunRepository     func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository
		masterActionStepRepository    func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository
		masterActionTriggerRepository func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository
		userActionRepository          func(ctrl *gomock.Controller) userAction.UserActionRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ActionGetMasterResponse
		wantErr error
	}{
		{
			name: "正常：取得できる",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterAction.MasterActions{
								{
									Id:                1,
									Name:              "テストアクション1",
									ActionStepType:    enum.ActionStepType_None,
									ActionTriggerType: enum.ActionTriggerType_Continuation,
									AnyId:             nil,
									TriggerActionId:   nil,
									Expiration:        nil,
								},
								{
									Id:                2,
									Name:              "テストアクション2",
									ActionStepType:    enum.ActionStepType_None,
									ActionTriggerType: enum.ActionTriggerType_Continuation,
									AnyId:             pointers.Int64ToPointer(2),
									TriggerActionId:   pointers.Int64ToPointer(1),
									Expiration:        pointers.Int32ToPointer(24),
								},
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									Id:       1,
									Name:     "テストアクション1",
									ActionId: 1,
								},
								{
									Id:       2,
									Name:     "テストアクション2",
									ActionId: 2,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionStep.MasterActionSteps{
								{
									Id:             1,
									Name:           "無",
									ActionStepType: enum.ActionStepType_None,
								},
								{
									Id:             2,
									Name:           "チュートリアル突破",
									ActionStepType: enum.ActionStepType_PassedTutorial,
								},
							},
							nil,
						)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionTrigger.MasterActionTriggers{
								{
									Id:                1,
									Name:              "無期限",
									ActionTriggerType: enum.ActionTriggerType_Continuation,
								},
								{
									Id:                2,
									Name:              "期限あり",
									ActionTriggerType: enum.ActionTriggerType_Discontinuation,
								},
							},
							nil,
						)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want: &ActionGetMasterResponse{
				MasterActions: masterAction.MasterActions{
					{
						Id:                1,
						Name:              "テストアクション1",
						ActionStepType:    enum.ActionStepType_None,
						ActionTriggerType: enum.ActionTriggerType_Continuation,
						AnyId:             nil,
						TriggerActionId:   nil,
						Expiration:        nil,
					},
					{
						Id:                2,
						Name:              "テストアクション2",
						ActionStepType:    enum.ActionStepType_None,
						ActionTriggerType: enum.ActionTriggerType_Continuation,
						AnyId:             pointers.Int64ToPointer(2),
						TriggerActionId:   pointers.Int64ToPointer(1),
						Expiration:        pointers.Int32ToPointer(24),
					},
				},
				MasterActionRuns: masterActionRun.MasterActionRuns{
					{
						Id:       1,
						Name:     "テストアクション1",
						ActionId: 1,
					},
					{
						Id:       2,
						Name:     "テストアクション2",
						ActionId: 2,
					},
				},
				MasterActionSteps: masterActionStep.MasterActionSteps{
					{
						Id:             1,
						Name:           "無",
						ActionStepType: enum.ActionStepType_None,
					},
					{
						Id:             2,
						Name:           "チュートリアル突破",
						ActionStepType: enum.ActionStepType_PassedTutorial,
					},
				},
				MasterActionTriggers: masterActionTrigger.MasterActionTriggers{
					{
						Id:                1,
						Name:              "無期限",
						ActionTriggerType: enum.ActionTriggerType_Continuation,
					},
					{
						Id:                2,
						Name:              "期限あり",
						ActionTriggerType: enum.ActionTriggerType_Discontinuation,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.masterActionRepository.FindList",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterActionRepository.FindList", errors.NewTestError()),
		},
		{
			name: "異常：s.masterActionRunRepository.FindList",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterAction.MasterActions{
								{
									Id:                1,
									Name:              "テストアクション1",
									ActionStepType:    enum.ActionStepType_None,
									ActionTriggerType: enum.ActionTriggerType_Continuation,
									AnyId:             nil,
									TriggerActionId:   nil,
									Expiration:        nil,
								},
								{
									Id:                2,
									Name:              "テストアクション2",
									ActionStepType:    enum.ActionStepType_None,
									ActionTriggerType: enum.ActionTriggerType_Continuation,
									AnyId:             pointers.Int64ToPointer(2),
									TriggerActionId:   pointers.Int64ToPointer(1),
									Expiration:        pointers.Int32ToPointer(24),
								},
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterActionRunRepository.FindList", errors.NewTestError()),
		},
		{
			name: "異常：s.masterActionStepRepository.FindList",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterAction.MasterActions{
								{
									Id:                1,
									Name:              "テストアクション1",
									ActionStepType:    enum.ActionStepType_None,
									ActionTriggerType: enum.ActionTriggerType_Continuation,
									AnyId:             nil,
									TriggerActionId:   nil,
									Expiration:        nil,
								},
								{
									Id:                2,
									Name:              "テストアクション2",
									ActionStepType:    enum.ActionStepType_None,
									ActionTriggerType: enum.ActionTriggerType_Continuation,
									AnyId:             pointers.Int64ToPointer(2),
									TriggerActionId:   pointers.Int64ToPointer(1),
									Expiration:        pointers.Int32ToPointer(24),
								},
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									Id:       1,
									Name:     "テストアクション1",
									ActionId: 1,
								},
								{
									Id:       2,
									Name:     "テストアクション2",
									ActionId: 2,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterActionStepRepository.FindList", errors.NewTestError()),
		},
		{
			name: "正常：取得できる",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterAction.MasterActions{
								{
									Id:                1,
									Name:              "テストアクション1",
									ActionStepType:    enum.ActionStepType_None,
									ActionTriggerType: enum.ActionTriggerType_Continuation,
									AnyId:             nil,
									TriggerActionId:   nil,
									Expiration:        nil,
								},
								{
									Id:                2,
									Name:              "テストアクション2",
									ActionStepType:    enum.ActionStepType_None,
									ActionTriggerType: enum.ActionTriggerType_Continuation,
									AnyId:             pointers.Int64ToPointer(2),
									TriggerActionId:   pointers.Int64ToPointer(1),
									Expiration:        pointers.Int32ToPointer(24),
								},
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									Id:       1,
									Name:     "テストアクション1",
									ActionId: 1,
								},
								{
									Id:       2,
									Name:     "テストアクション2",
									ActionId: 2,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							masterActionStep.MasterActionSteps{
								{
									Id:             1,
									Name:           "無",
									ActionStepType: enum.ActionStepType_None,
								},
								{
									Id:             2,
									Name:           "チュートリアル突破",
									ActionStepType: enum.ActionStepType_PassedTutorial,
								},
							},
							nil,
						)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					m.EXPECT().
						FindList(
							nil,
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.masterActionTriggerRepository.FindList", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &actionService{
				masterActionRepository:        tt.fields.masterActionRepository(ctrl),
				masterActionRunRepository:     tt.fields.masterActionRunRepository(ctrl),
				masterActionStepRepository:    tt.fields.masterActionStepRepository(ctrl),
				masterActionTriggerRepository: tt.fields.masterActionTriggerRepository(ctrl),
				userActionRepository:          tt.fields.userActionRepository(ctrl),
			}

			got, err := s.GetMaster(tt.args.ctx)
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

func TestActionService_Check(t *testing.T) {
	type fields struct {
		masterActionRepository        func(ctrl *gomock.Controller) masterAction.MasterActionRepository
		masterActionRunRepository     func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository
		masterActionStepRepository    func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository
		masterActionTriggerRepository func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository
		userActionRepository          func(ctrl *gomock.Controller) userAction.UserActionRepository
	}
	type args struct {
		ctx context.Context
		now time.Time
		req *ActionCheckRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "正常：確認できる（AnyIdがある場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(1),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：確認できる（AnyIdがない場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepType(
							nil,
							enum.ActionStepType_None,
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             nil,
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          nil,
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：アクションが存在しない（AnyIdがある場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							nil,
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(1),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：アクションが存在しない（AnyIdがない場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepType(
							nil,
							enum.ActionStepType_None,
						).
						Return(
							nil,
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          nil,
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.getMasterAction: failed to s.masterActionRepository.FindOrNilByActionStepTypeAndAnyId",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(1),
				},
			},
			wantErr: errors.NewMethodError("s.getMasterAction: failed to s.masterActionRepository.FindOrNilByActionStepTypeAndAnyId", errors.NewTestError()),
		},
		{
			name: "異常：s.getUserAction: failed to s.userActionRepository.Find",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
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
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(1),
				},
			},
			wantErr: errors.NewMethodError("s.getUserAction: failed to s.userActionRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.getMasterAction: failed to s.masterActionRepository.FindOrNilByActionStepType",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepType(
							nil,
							enum.ActionStepType_None,
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          nil,
				},
			},
			wantErr: errors.NewMethodError("s.getMasterAction: failed to s.masterActionRepository.FindOrNilByActionStepType", errors.NewTestError()),
		},
		{
			name: "異常：s.getUserAction: failed to s.userActionRepository.Find",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepType(
							nil,
							enum.ActionStepType_None,
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             nil,
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
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
				now: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          nil,
				},
			},
			wantErr: errors.NewMethodError("s.getUserAction: failed to s.userActionRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.getUserAction: expiration date has expired",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				now: time.Date(2023, 1, 2, 1, 0, 0, 0, time.UTC),
				req: &ActionCheckRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(1),
				},
			},
			wantErr: errors.NewError("failed to s.getUserAction: expiration date has expired"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &actionService{
				masterActionRepository:        tt.fields.masterActionRepository(ctrl),
				masterActionRunRepository:     tt.fields.masterActionRunRepository(ctrl),
				masterActionStepRepository:    tt.fields.masterActionStepRepository(ctrl),
				masterActionTriggerRepository: tt.fields.masterActionTriggerRepository(ctrl),
				userActionRepository:          tt.fields.userActionRepository(ctrl),
			}

			err := s.Check(tt.args.ctx, tt.args.now, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestActionService_Run(t *testing.T) {
	type fields struct {
		masterActionRepository        func(ctrl *gomock.Controller) masterAction.MasterActionRepository
		masterActionRunRepository     func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository
		masterActionStepRepository    func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository
		masterActionTriggerRepository func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository
		userActionRepository          func(ctrl *gomock.Controller) userAction.UserActionRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		now time.Time
		req *ActionRunRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "正常：実行できる（AnyIdがある場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindListByActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（AnyIdがない場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepType(
							nil,
							enum.ActionStepType_None,
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindListByActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          nil,
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（トリガーアクションがない場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindListByActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（トリガーアクションがDiscontinuationの場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Discontinuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindListByActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：アクションがない場合",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							nil,
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（ActionRunが存在する場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindListByActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									Id:       3,
									Name:     "アクションRun3",
									ActionId: 3,
								},
								{
									Id:       4,
									Name:     "アクションRun4",
									ActionId: 4,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(3),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 3,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 3,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(4),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 4,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 4,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "正常：実行できる（期限切れにより再実行された場合）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Discontinuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        pointers.Int32ToPointer(32),
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindListByActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2022, 12, 30, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.getMasterAction: failed to s.masterActionRepository.FindOrNilByActionStepTypeAndAnyId",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.getMasterAction: failed to s.masterActionRepository.FindOrNilByActionStepTypeAndAnyId", errors.NewTestError()),
		},
		{
			name: "異常：s.checkTriggerUserAction: failed to s.masterActionRepository.Find）",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
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
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.checkTriggerUserAction: failed to s.masterActionRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.deleteTriggerAction: failed to s.userActionRepository.Delete",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Discontinuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
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
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.deleteTriggerAction: failed to s.userActionRepository.Delete", errors.NewTestError()),
		},
		{
			name: "異常：s.run: failed to s.update: failed to s.userActionRepository.FindOrNil",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.update: failed to s.userActionRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：s.run: failed to s.update: failed to s.userActionRepository.Create",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.update: failed to s.userActionRepository.Create", errors.NewTestError()),
		},
		{
			name: "異常：s.getMasterAction: failed to s.masterActionRepository.FindOrNilByActionStepType",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepType(
							nil,
							enum.ActionStepType_None,
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          nil,
				},
			},
			wantErr: errors.NewMethodError("s.getMasterAction: failed to s.masterActionRepository.FindOrNilByActionStepType", errors.NewTestError()),
		},
		{
			name: "異常：s.checkTriggerUserAction: failed to s.getUserAction: failed to s.userActionRepository.Find",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.checkTriggerUserAction: failed to s.getUserAction: failed to s.userActionRepository.Find", errors.NewTestError()),
		},
		{
			name: "異常：s.checkTriggerUserAction: failed to s.getUserAction: expiration date has expired",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Discontinuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 3, 1, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewError("failed to s.checkTriggerUserAction: failed to s.getUserAction: expiration date has expired"),
		},
		{
			name: "異常：s.deleteTriggerAction: ActionTriggerType does not exist",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: 999,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        pointers.Int32ToPointer(24),
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewError("failed to s.deleteTriggerAction: ActionTriggerType does not exist"),
		},
		{
			name: "異常：s.run: failed to s.masterActionRunRepository.FindListByActionId",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindListByActionId(
							nil,
							int64(2),
						).
						Return(
							nil,
							errors.NewTestError(),
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.masterActionRunRepository.FindListByActionId", errors.NewTestError()),
		},
		{
			name: "異常：s.run: failed to s.userActionRepository.Create: failed to s.userActionRepository.FindOrNil",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        nil,
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					m.EXPECT().
						FindListByActionId(
							nil,
							int64(2),
						).
						Return(
							masterActionRun.MasterActionRuns{
								{
									Id:       3,
									Name:     "アクションRun3",
									ActionId: 3,
								},
								{
									Id:       4,
									Name:     "アクションRun4",
									ActionId: 4,
								},
							},
							nil,
						)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							nil,
							nil,
						)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(3),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.userActionRepository.Create: failed to s.userActionRepository.FindOrNil", errors.NewTestError()),
		},
		{
			name: "異常：s.run: failed to s.update: failed to s.userActionRepository.Update",
			fields: fields{
				masterActionRepository: func(ctrl *gomock.Controller) masterAction.MasterActionRepository {
					m := masterAction.NewMockMasterActionRepository(ctrl)
					m.EXPECT().
						FindOrNilByActionStepTypeAndAnyId(
							nil,
							enum.ActionStepType_None,
							pointers.Int64ToPointer(2),
						).
						Return(
							&masterAction.MasterAction{
								Id:                2,
								Name:              "テストアクション2",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Continuation,
								AnyId:             pointers.Int64ToPointer(2),
								TriggerActionId:   pointers.Int64ToPointer(1),
								Expiration:        nil,
							},
							nil,
						)
					m.EXPECT().
						Find(
							nil,
							int64(1),
						).
						Return(
							&masterAction.MasterAction{
								Id:                1,
								Name:              "テストアクション1",
								ActionStepType:    enum.ActionStepType_None,
								ActionTriggerType: enum.ActionTriggerType_Discontinuation,
								AnyId:             pointers.Int64ToPointer(1),
								TriggerActionId:   nil,
								Expiration:        pointers.Int32ToPointer(32),
							},
							nil,
						)
					return m
				},
				masterActionRunRepository: func(ctrl *gomock.Controller) masterActionRun.MasterActionRunRepository {
					m := masterActionRun.NewMockMasterActionRunRepository(ctrl)
					return m
				},
				masterActionStepRepository: func(ctrl *gomock.Controller) masterActionStep.MasterActionStepRepository {
					m := masterActionStep.NewMockMasterActionStepRepository(ctrl)
					return m
				},
				masterActionTriggerRepository: func(ctrl *gomock.Controller) masterActionTrigger.MasterActionTriggerRepository {
					m := masterActionTrigger.NewMockMasterActionTriggerRepository(ctrl)
					return m
				},
				userActionRepository: func(ctrl *gomock.Controller) userAction.UserActionRepository {
					m := userAction.NewMockUserActionRepository(ctrl)
					m.EXPECT().
						Find(
							nil,
							"0:test",
							int64(1),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Delete(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 1,
								StartedAt:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						).
						Return(
							nil,
						)
					m.EXPECT().
						FindOrNil(
							nil,
							"0:test",
							int64(2),
						).
						Return(
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2022, 12, 30, 0, 0, 0, 0, time.UTC),
							},
							nil,
						)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userAction.UserAction{
								UserId:         "0:test",
								MasterActionId: 2,
								StartedAt:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
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
				now: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
				req: &ActionRunRequest{
					UserId:         "0:test",
					ActionStepType: enum.ActionStepType_None,
					AnyId:          pointers.Int64ToPointer(2),
				},
			},
			wantErr: errors.NewMethodError("s.run: failed to s.update: failed to s.userActionRepository.Update", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &actionService{
				masterActionRepository:        tt.fields.masterActionRepository(ctrl),
				masterActionRunRepository:     tt.fields.masterActionRunRepository(ctrl),
				masterActionStepRepository:    tt.fields.masterActionStepRepository(ctrl),
				masterActionTriggerRepository: tt.fields.masterActionTriggerRepository(ctrl),
				userActionRepository:          tt.fields.userActionRepository(ctrl),
			}

			err := s.Run(tt.args.ctx, tt.args.tx, tt.args.now, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
