// Package health ヘルスチェックリクエスト
package health

type HealthCheckRequests []*HealthCheckRequest

type HealthCheckRequest struct {
	Message string
}

func NewHealthCheckRequest() *HealthCheckRequest {
	return &HealthCheckRequest{}
}

func NewHealthCheckRequests() HealthCheckRequests {
	return HealthCheckRequests{}
}

func SetHealthCheckRequest(message string) *HealthCheckRequest {
	return &HealthCheckRequest{
		Message: message,
	}
}
