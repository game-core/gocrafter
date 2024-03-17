package api

type APIConfig struct {
	Port Port
}

type Port struct {
	HTTPPort int
	GrpcPort int
}
