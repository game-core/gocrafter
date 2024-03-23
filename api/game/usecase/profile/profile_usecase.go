package profile

import (
	"context"

	profileServer "github.com/game-core/gocrafter/api/game/presentation/server/profile"
	"github.com/game-core/gocrafter/internal/errors"
	profileService "github.com/game-core/gocrafter/pkg/domain/model/profile"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
)

type ProfileUsecase interface {
	Get(ctx context.Context, req *profileServer.ProfileGetRequest) (*profileServer.ProfileGetResponse, error)
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

// Get プロフィールを作成する
func (s *profileUsecase) Get(ctx context.Context, req *profileServer.ProfileGetRequest) (*profileServer.ProfileGetResponse, error) {
	result, err := s.profileService.Get(ctx, profileService.SetProfileGetRequest(req.UserId))
	if err != nil {
		return nil, errors.NewMethodError("s.profileService.Get", err)
	}

	return profileServer.SetProfileGetResponse(
		profileServer.SetUserProfile(
			result.UserProfile.UserId,
			result.UserProfile.Name,
			result.UserProfile.Content,
		),
	), nil
}

// Create プロフィールを作成する
func (s *profileUsecase) Create(ctx context.Context, req *profileServer.ProfileCreateRequest) (*profileServer.ProfileCreateResponse, error) {
	// transaction
	tx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
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

// Update プロフィールを更新する
func (s *profileUsecase) Update(ctx context.Context, req *profileServer.ProfileUpdateRequest) (*profileServer.ProfileUpdateResponse, error) {
	// transaction
	tx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
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
