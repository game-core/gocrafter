package router

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/game-core/gocrafter/api/game/di"
	"github.com/game-core/gocrafter/api/game/presentation/server/account"
)

func Router(lis net.Listener) {
	// DI
	accountHandler := di.InitializeAccountHandler()
	authInterceptor := di.InitializeAuthInterceptor()

	// Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.JwtAuth),
	)

	account.RegisterAccountServer(s, accountHandler)

	log.Printf("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
