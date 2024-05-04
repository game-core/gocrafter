// Package health ヘルスチェックリクエスト
package health

func SetHealthCheckRequest(message string) *HealthCheckRequest {
	return &HealthCheckRequest{
		Message: message,
	}
}
