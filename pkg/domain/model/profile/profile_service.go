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
	userProfileRepository userProfile.UserProfileRepository
}

func NewProfileService(
	userProfileRepository userProfile.UserProfileRepository,
) ProfileService {
	return &profileService{
		userProfileRepository: userProfileRepository,
	}
}

// Get プロフィールを取得する
func (s *profileService) Get(ctx context.Context, req *ProfileGetRequest) (*ProfileGetResponse, error) {
	userProfileModel, err := s.userProfileRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userProfileRepository.Find", err)
	}

	return SetProfileGetResponse(userProfileModel), nil
}

// Create プロフィールを作成する
func (s *profileService) Create(ctx context.Context, tx *gorm.DB, req *ProfileCreateRequest) (*ProfileCreateResponse, error) {
	userProfileModel, err := s.userProfileRepository.Create(ctx, tx, userProfile.SetUserProfile(req.UserId, req.Name, req.Content))
	if err != nil {
		return nil, errors.NewMethodError("s.userProfileRepository.Create", err)
	}

	return SetProfileCreateResponse(userProfileModel), nil
}

// Update プロフィールを更新する
func (s *profileService) Update(ctx context.Context, tx *gorm.DB, req *ProfileUpdateRequest) (*ProfileUpdateResponse, error) {
	userProfileModel, err := s.userProfileRepository.Update(ctx, tx, userProfile.SetUserProfile(req.UserId, req.Name, req.Content))
	if err != nil {
		return nil, errors.NewMethodError("s.userProfileRepository.Update", err)
	}

	return SetProfileUpdateResponse(userProfileModel), nil
}
