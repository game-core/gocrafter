//go:generate mockgen -source=./profile_service.go -destination=./profile_service_mock.gen.go -package=profile
package profile

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/profile/userProfile"
)

type ProfileService interface {
	Get(ctx context.Context, req *ProfileGetRequest) (*ProfileGetResponse, error)
	Create(ctx context.Context, tx *gorm.DB, req *ProfileCreateRequest) (*ProfileCreateResponse, error)
	Update(ctx context.Context, tx *gorm.DB, req *ProfileUpdateRequest) (*ProfileUpdateResponse, error)
}

type profileService struct {
	userProfileMysqlRepository userProfile.UserProfileMysqlRepository
}

func NewProfileService(
	userProfileMysqlRepository userProfile.UserProfileMysqlRepository,
) ProfileService {
	return &profileService{
		userProfileMysqlRepository: userProfileMysqlRepository,
	}
}

// Get プロフィールを取得する
func (s *profileService) Get(ctx context.Context, req *ProfileGetRequest) (*ProfileGetResponse, error) {
	userProfileModel, err := s.userProfileMysqlRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userProfileMysqlRepository.Find", err)
	}

	return SetProfileGetResponse(userProfileModel), nil
}

// Create プロフィールを作成する
func (s *profileService) Create(ctx context.Context, tx *gorm.DB, req *ProfileCreateRequest) (*ProfileCreateResponse, error) {
	userProfileModel, err := s.userProfileMysqlRepository.Create(ctx, tx, userProfile.SetUserProfile(req.UserId, req.Name, req.Content))
	if err != nil {
		return nil, errors.NewMethodError("s.userProfileMysqlRepository.Create", err)
	}

	return SetProfileCreateResponse(userProfileModel), nil
}

// Update プロフィールを更新する
func (s *profileService) Update(ctx context.Context, tx *gorm.DB, req *ProfileUpdateRequest) (*ProfileUpdateResponse, error) {
	userProfileModel, err := s.userProfileMysqlRepository.Update(ctx, tx, userProfile.SetUserProfile(req.UserId, req.Name, req.Content))
	if err != nil {
		return nil, errors.NewMethodError("s.userProfileMysqlRepository.Update", err)
	}

	return SetProfileUpdateResponse(userProfileModel), nil
}
