package router

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/game-core/gocrafter/api/game/di"
	"github.com/game-core/gocrafter/api/game/presentation/server/account"
)

func Router(lis net.Listener) {
	accountHandler := di.InitializeAccountHandler()

	s := grpc.NewServer()

	account.RegisterAccountServer(s, accountHandler)

	log.Printf("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
