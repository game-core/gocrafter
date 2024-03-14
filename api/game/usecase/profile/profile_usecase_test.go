package profile

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	profileServer "github.com/game-core/gocrafter/api/game/presentation/server/profile"
	"github.com/game-core/gocrafter/internal/errors"
	profileService "github.com/game-core/gocrafter/pkg/domain/model/profile"
	"github.com/game-core/gocrafter/pkg/domain/model/profile/userProfile"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

func TestProfileUsecase_NewProfileUsecase(t *testing.T) {
	type args struct {
		profileService     profileService.ProfileService
		transactionService transactionService.TransactionService
	}
	tests := []struct {
		name string
		args args
		want ProfileUsecase
	}{
		{
			name: "正常",
			args: args{
				profileService:     nil,
				transactionService: nil,
			},
			want: &profileUsecase{
				profileService:     nil,
				transactionService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewProfileUsecase(tt.args.profileService, tt.args.transactionService)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProfileUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileUsecase_Create(t *testing.T) {
	type fields struct {
		profileService     func(ctrl *gomock.Controller) profileService.ProfileService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *profileServer.ProfileCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *profileServer.ProfileCreateResponse
		wantErr error
	}{
		{
			name: "正常：プロフィールが作成できる",
			fields: fields{
				profileService: func(ctrl *gomock.Controller) profileService.ProfileService {
					m := profileService.NewMockProfileService(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							&profileService.ProfileCreateRequest{
								UserId:  "0:test",
								Name:    "test_user_profile",
								Content: "test_user_profile_contest",
							},
						).
						Return(
							&profileService.ProfileCreateResponse{
								UserProfile: &userProfile.UserProfile{
									UserId:  "0:test",
									Name:    "test_user_profile",
									Content: "test_user_profile_contest",
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
				req: &profileServer.ProfileCreateRequest{
					UserId:  "0:test",
					Name:    "test_user_profile",
					Content: "test_user_profile_contest",
				},
			},
			want: &profileServer.ProfileCreateResponse{
				UserProfile: &profileServer.UserProfile{
					UserId:  "0:test",
					Name:    "test_user_profile",
					Content: "test_user_profile_contest",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.UserBegin",
			fields: fields{
				profileService: func(ctrl *gomock.Controller) profileService.ProfileService {
					m := profileService.NewMockProfileService(ctrl)
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
				req: &profileServer.ProfileCreateRequest{
					UserId:  "0:test",
					Name:    "test_user_profile",
					Content: "test_user_profile_contest",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.UserBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.profileService.Create",
			fields: fields{
				profileService: func(ctrl *gomock.Controller) profileService.ProfileService {
					m := profileService.NewMockProfileService(ctrl)
					m.EXPECT().
						Create(
							gomock.Any(),
							gomock.Any(),
							&profileService.ProfileCreateRequest{
								UserId:  "0:test",
								Name:    "test_user_profile",
								Content: "test_user_profile_contest",
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
				req: &profileServer.ProfileCreateRequest{
					UserId:  "0:test",
					Name:    "test_user_profile",
					Content: "test_user_profile_contest",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.profileService.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &profileUsecase{
				profileService:     tt.fields.profileService(ctrl),
				transactionService: tt.fields.transactionService(ctrl),
			}

			got, err := u.Create(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileUsecase_Update(t *testing.T) {
	type fields struct {
		profileService     func(ctrl *gomock.Controller) profileService.ProfileService
		transactionService func(ctrl *gomock.Controller) transactionService.TransactionService
	}
	type args struct {
		ctx context.Context
		req *profileServer.ProfileUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *profileServer.ProfileUpdateResponse
		wantErr error
	}{
		{
			name: "正常：プロフィールが更新できる",
			fields: fields{
				profileService: func(ctrl *gomock.Controller) profileService.ProfileService {
					m := profileService.NewMockProfileService(ctrl)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							&profileService.ProfileUpdateRequest{
								UserId:  "0:test",
								Name:    "test_user_profile",
								Content: "test_user_profile_contest",
							},
						).
						Return(
							&profileService.ProfileUpdateResponse{
								UserProfile: &userProfile.UserProfile{
									UserId:  "0:test",
									Name:    "test_user_profile",
									Content: "test_user_profile_contest",
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
				req: &profileServer.ProfileUpdateRequest{
					UserId:  "0:test",
					Name:    "test_user_profile",
					Content: "test_user_profile_contest",
				},
			},
			want: &profileServer.ProfileUpdateResponse{
				UserProfile: &profileServer.UserProfile{
					UserId:  "0:test",
					Name:    "test_user_profile",
					Content: "test_user_profile_contest",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.transactionService.UserBegin",
			fields: fields{
				profileService: func(ctrl *gomock.Controller) profileService.ProfileService {
					m := profileService.NewMockProfileService(ctrl)
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
				req: &profileServer.ProfileUpdateRequest{
					UserId:  "0:test",
					Name:    "test_user_profile",
					Content: "test_user_profile_contest",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.transactionService.UserBegin", errors.NewTestError()),
		},
		{
			name: "異常：s.profileService.Update",
			fields: fields{
				profileService: func(ctrl *gomock.Controller) profileService.ProfileService {
					m := profileService.NewMockProfileService(ctrl)
					m.EXPECT().
						Update(
							gomock.Any(),
							gomock.Any(),
							&profileService.ProfileUpdateRequest{
								UserId:  "0:test",
								Name:    "test_user_profile",
								Content: "test_user_profile_contest",
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
				req: &profileServer.ProfileUpdateRequest{
					UserId:  "0:test",
					Name:    "test_user_profile",
					Content: "test_user_profile_contest",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.profileService.Update", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			u := &profileUsecase{
				profileService:     tt.fields.profileService(ctrl),
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
