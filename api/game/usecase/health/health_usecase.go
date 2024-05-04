package health

import (
	"context"

	healthServer "github.com/game-core/gocrafter/api/game/presentation/server/health"
	"github.com/game-core/gocrafter/internal/errors"
	healthService "github.com/game-core/gocrafter/pkg/domain/model/health"
)

type HealthUsecase interface {
	Check(ctx context.Context, req *healthServer.HealthCheckRequest) (*healthServer.HealthCheckResponse, error)
}

type healthUsecase struct {
	healthService healthService.HealthService
}

func NewHealthUsecase(
	healthService healthService.HealthService,
) HealthUsecase {
	return &healthUsecase{
		healthService: healthService,
	}
}

// Check ヘルスチェック
func (s *healthUsecase) Check(ctx context.Context, req *healthServer.HealthCheckRequest) (*healthServer.HealthCheckResponse, error) {
	res, err := s.healthService.Check(ctx, healthService.SetHealthCheckRequest(req.Message))
	if err != nil {
		return nil, errors.NewMethodError("s.healthService.Check", err)
	}

	return healthServer.SetHealthCheckResponse(res.Message), nil
}
