package router

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/game-core/gocrafter/api/multi/di"
)

func Router(lis net.Listener) {
	// DI
	authInterceptor := di.InitializeAuthInterceptor()

	// Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.JwtAuth),
	)

	log.Printf("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
