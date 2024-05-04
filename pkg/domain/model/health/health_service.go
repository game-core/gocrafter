//go:generate mockgen -source=./health_service.go -destination=./health_service_mock.gen.go -package=health
package health

import (
	"context"
)

type HealthService interface {
	Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error)
}

type healthService struct {
}

func NewHealthService() HealthService {
	return &healthService{}
}

// Check ヘルスチェック
func (s *healthService) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return SetHealthCheckResponse("StatusOK"), nil
}
