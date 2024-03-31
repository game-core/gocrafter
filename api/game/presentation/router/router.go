package router

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/game-core/gocrafter/api/game/di"
	"github.com/game-core/gocrafter/api/game/presentation/server/account"
	"github.com/game-core/gocrafter/api/game/presentation/server/friend"
	"github.com/game-core/gocrafter/api/game/presentation/server/idleBonus"
	"github.com/game-core/gocrafter/api/game/presentation/server/loginBonus"
	"github.com/game-core/gocrafter/api/game/presentation/server/profile"
	"github.com/game-core/gocrafter/api/game/presentation/server/ranking"
	"github.com/game-core/gocrafter/api/game/presentation/server/room"
)

func Router(lis net.Listener) {
	// DI
	accountHandler := di.InitializeAccountHandler()
	friendHandler := di.InitializeFriendHandler()
	authInterceptor := di.InitializeAuthInterceptor()
	loginBonusHandler := di.InitializeLoginBonusHandler()
	profileHandler := di.InitializeProfileHandler()
	idleHandler := di.InitializeIdleBonusHandler()
	rankingHandler := di.InitializeRankingHandler()
	roomHandler := di.InitializeRoomHandler()

	// Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.JwtAuth),
	)

	account.RegisterAccountServer(s, accountHandler)
	friend.RegisterFriendServer(s, friendHandler)
	loginBonus.RegisterLoginBonusServer(s, loginBonusHandler)
	profile.RegisterProfileServer(s, profileHandler)
	idleBonus.RegisterIdleBonusServer(s, idleHandler)
	ranking.RegisterRankingServer(s, rankingHandler)
	room.RegisterRoomServer(s, roomHandler)

	log.Printf("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
