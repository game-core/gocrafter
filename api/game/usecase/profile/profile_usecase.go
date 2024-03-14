package profile

import (
	"context"

	profileServer "github.com/game-core/gocrafter/api/game/presentation/server/profile"
	"github.com/game-core/gocrafter/internal/errors"
	profileService "github.com/game-core/gocrafter/pkg/domain/model/profile"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

type ProfileUsecase interface {
	Create(ctx context.Context, req *profileServer.ProfileCreateRequest) (*profileServer.ProfileCreateResponse, error)
	Update(ctx context.Context, req *profileServer.ProfileUpdateRequest) (*profileServer.ProfileUpdateResponse, error)
}

type profileUsecase struct {
	profileService     profileService.ProfileService
	transactionService transactionService.TransactionService
}

func NewProfileUsecase(
	profileService profileService.ProfileService,
	transactionService transactionService.TransactionService,
) ProfileUsecase {
	return &profileUsecase{
		profileService:     profileService,
		transactionService: transactionService,
	}
}

// Create アカウントを作成する
func (s *profileUsecase) Create(ctx context.Context, req *profileServer.ProfileCreateRequest) (*profileServer.ProfileCreateResponse, error) {
	// transaction
	tx, err := s.transactionService.UserBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserBegin", err)
	}
	defer func() {
		s.transactionService.UserEnd(ctx, tx, err)
	}()

	result, err := s.profileService.Create(ctx, tx, profileService.SetProfileCreateRequest(req.UserId, req.Name, req.Content))
	if err != nil {
		return nil, errors.NewMethodError("s.profileService.Create", err)
	}

	return profileServer.SetProfileCreateResponse(
		profileServer.SetUserProfile(
			result.UserProfile.UserId,
			result.UserProfile.Name,
			result.UserProfile.Content,
		),
	), nil
}

// Update アカウントを更新する
func (s *profileUsecase) Update(ctx context.Context, req *profileServer.ProfileUpdateRequest) (*profileServer.ProfileUpdateResponse, error) {
	// transaction
	tx, err := s.transactionService.UserBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserBegin", err)
	}
	defer func() {
		s.transactionService.UserEnd(ctx, tx, err)
	}()

	result, err := s.profileService.Update(ctx, tx, profileService.SetProfileUpdateRequest(req.UserId, req.Name, req.Content))
	if err != nil {
		return nil, errors.NewMethodError("s.profileService.Update", err)
	}

	return profileServer.SetProfileUpdateResponse(
		profileServer.SetUserProfile(
			result.UserProfile.UserId,
			result.UserProfile.Name,
			result.UserProfile.Content,
		),
	), nil
}
