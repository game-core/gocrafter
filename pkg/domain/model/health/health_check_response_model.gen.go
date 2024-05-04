// Package health ヘルスチェックレスポンス
package health

type HealthCheckResponses []*HealthCheckResponse

type HealthCheckResponse struct {
	Message string
}

func NewHealthCheckResponse() *HealthCheckResponse {
	return &HealthCheckResponse{}
}

func NewHealthCheckResponses() HealthCheckResponses {
	return HealthCheckResponses{}
}

func SetHealthCheckResponse(message string) *HealthCheckResponse {
	return &HealthCheckResponse{
		Message: message,
	}
}
