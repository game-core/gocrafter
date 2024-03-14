package profile

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/profile"
	profileUsecase "github.com/game-core/gocrafter/api/game/usecase/profile"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/tokens"
)

type ProfileHandler interface {
	profile.ProfileServer
}

type profileHandler struct {
	profile.UnimplementedProfileServer
	profileUsecase profileUsecase.ProfileUsecase
}

func NewProfileHandler(
	profileUsecase profileUsecase.ProfileUsecase,
) ProfileHandler {
	return &profileHandler{
		profileUsecase: profileUsecase,
	}
}

// Create フレンドを作成する
func (s *profileHandler) Create(ctx context.Context, req *profile.ProfileCreateRequest) (*profile.ProfileCreateResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.profileUsecase.Create(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.profileUsecase.Create", err)
	}

	return res, nil
}

// Update フレンドを更新する
func (s *profileHandler) Update(ctx context.Context, req *profile.ProfileUpdateRequest) (*profile.ProfileUpdateResponse, error) {
	if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
		return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
	}
	res, err := s.profileUsecase.Update(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.profileUsecase.Update", err)
	}

	return res, nil
}
