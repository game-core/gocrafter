// Package health ヘルスチェックレスポンス
package health

func SetHealthCheckResponse(message string) *HealthCheckResponse {
	return &HealthCheckResponse{
		Message: message,
	}
}
