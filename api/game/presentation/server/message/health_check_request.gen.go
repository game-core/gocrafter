// Package message ヘルスチェックリクエスト
package message

func SetHealthCheckRequest(message string) *HealthCheckRequest {
	return &HealthCheckRequest{
		Message: message,
	}
}
