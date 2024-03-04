package api

// GetAppConfig 設定を取得する
func GetAppConfig() *APIConfig {
	return &APIConfig{
		Port: Port{
			HTTPPort: 80,
			GrpcPort: 50051,
		},
	}
}
