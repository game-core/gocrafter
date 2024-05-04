// Package message ヘルスチェックレスポンス
package message

func SetHealthCheckResponse(message string) *HealthCheckResponse {
	return &HealthCheckResponse{
		Message: message,
	}
}
