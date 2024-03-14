package profile

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/profile/userProfile"
)

func TestNewProfileService_NewProfileService(t *testing.T) {
	type args struct {
		userProfileRepository userProfile.UserProfileRepository
	}
	tests := []struct {
		name string
		args args
		want ProfileService
	}{
		{
			name: "正常",
			args: args{
				userProfileRepository: nil,
			},
			want: &profileService{
				userProfileRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewProfileService(
				tt.args.userProfileRepository,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProfileService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileService_Create(t *testing.T) {
	type fields struct {
		userProfileRepository func(ctrl *gomock.Controller) userProfile.UserProfileRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *ProfileCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ProfileCreateResponse
		wantErr error
	}{
		{
			name: "正常：作成できる場合",
			fields: fields{
				userProfileRepository: func(ctrl *gomock.Controller) userProfile.UserProfileRepository {
					m := userProfile.NewMockUserProfileRepository(ctrl)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userProfile.UserProfile{
								UserId:  "0:WntR-PyhOJeDiE5jodeR",
								Name:    "test_user_profile",
								Content: "test_user_profile_content",
							},
						).
						Return(
							&userProfile.UserProfile{
								UserId:  "0:WntR-PyhOJeDiE5jodeR",
								Name:    "test_user_profile",
								Content: "test_user_profile_content",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ProfileCreateRequest{
					UserId:  "0:WntR-PyhOJeDiE5jodeR",
					Name:    "test_user_profile",
					Content: "test_user_profile_content",
				},
			},
			want: &ProfileCreateResponse{
				UserProfile: &userProfile.UserProfile{
					UserId:  "0:WntR-PyhOJeDiE5jodeR",
					Name:    "test_user_profile",
					Content: "test_user_profile_content",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userProfileRepository.Create",
			fields: fields{
				userProfileRepository: func(ctrl *gomock.Controller) userProfile.UserProfileRepository {
					m := userProfile.NewMockUserProfileRepository(ctrl)
					m.EXPECT().
						Create(
							nil,
							nil,
							&userProfile.UserProfile{
								UserId:  "0:WntR-PyhOJeDiE5jodeR",
								Name:    "test_user_profile",
								Content: "test_user_profile_content",
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
				req: &ProfileCreateRequest{
					UserId:  "0:WntR-PyhOJeDiE5jodeR",
					Name:    "test_user_profile",
					Content: "test_user_profile_content",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userProfileRepository.Create", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &profileService{
				userProfileRepository: tt.fields.userProfileRepository(ctrl),
			}

			got, err := s.Create(tt.args.ctx, tt.args.tx, tt.args.req)
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

func TestProfileService_Update(t *testing.T) {
	type fields struct {
		userProfileRepository func(ctrl *gomock.Controller) userProfile.UserProfileRepository
	}
	type args struct {
		ctx context.Context
		tx  *gorm.DB
		req *ProfileUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ProfileUpdateResponse
		wantErr error
	}{
		{
			name: "正常：更新できる場合",
			fields: fields{
				userProfileRepository: func(ctrl *gomock.Controller) userProfile.UserProfileRepository {
					m := userProfile.NewMockUserProfileRepository(ctrl)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userProfile.UserProfile{
								UserId:  "0:WntR-PyhOJeDiE5jodeR",
								Name:    "test_user_profile",
								Content: "test_user_profile_content",
							},
						).
						Return(
							&userProfile.UserProfile{
								UserId:  "0:WntR-PyhOJeDiE5jodeR",
								Name:    "test_user_profile",
								Content: "test_user_profile_content",
							},
							nil,
						)
					return m
				},
			},
			args: args{
				ctx: nil,
				tx:  nil,
				req: &ProfileUpdateRequest{
					UserId:  "0:WntR-PyhOJeDiE5jodeR",
					Name:    "test_user_profile",
					Content: "test_user_profile_content",
				},
			},
			want: &ProfileUpdateResponse{
				UserProfile: &userProfile.UserProfile{
					UserId:  "0:WntR-PyhOJeDiE5jodeR",
					Name:    "test_user_profile",
					Content: "test_user_profile_content",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常：s.userProfileRepository.Create",
			fields: fields{
				userProfileRepository: func(ctrl *gomock.Controller) userProfile.UserProfileRepository {
					m := userProfile.NewMockUserProfileRepository(ctrl)
					m.EXPECT().
						Update(
							nil,
							nil,
							&userProfile.UserProfile{
								UserId:  "0:WntR-PyhOJeDiE5jodeR",
								Name:    "test_user_profile",
								Content: "test_user_profile_content",
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
				req: &ProfileUpdateRequest{
					UserId:  "0:WntR-PyhOJeDiE5jodeR",
					Name:    "test_user_profile",
					Content: "test_user_profile_content",
				},
			},
			want:    nil,
			wantErr: errors.NewMethodError("s.userProfileRepository.Update", errors.NewTestError()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &profileService{
				userProfileRepository: tt.fields.userProfileRepository(ctrl),
			}

			got, err := s.Update(tt.args.ctx, tt.args.tx, tt.args.req)
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
