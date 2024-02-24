// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-core/gocrafter/api/game/presentation/handler/account"
	"github.com/game-core/gocrafter/api/game/presentation/interceptor/auth"
	account2 "github.com/game-core/gocrafter/api/game/usecase/account"
	"github.com/game-core/gocrafter/configs/database"
	account3 "github.com/game-core/gocrafter/pkg/domain/model/account"
	"github.com/game-core/gocrafter/pkg/domain/model/shard"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonShard"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonTransaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterTransaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userAccount"
	masterTransaction2 "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userTransaction"
)

// Injectors from wire.go:

func InitializeAuthInterceptor() auth.AuthInterceptor {
	authInterceptor := auth.NewAuthInterceptor()
	return authInterceptor
}

func InitializeAccountHandler() account.AccountHandler {
	accountUsecase := InitializeAccountUsecase()
	accountHandler := account.NewAccountHandler(accountUsecase)
	return accountHandler
}

func InitializeAccountUsecase() account2.AccountUsecase {
	accountService := InitializeAccountService()
	transactionService := InitializeTransactionService()
	accountUsecase := account2.NewAccountUsecase(accountService, transactionService)
	return accountUsecase
}

func InitializeAccountService() account3.AccountService {
	shardService := InitializeShardService()
	sqlHandler := database.NewDB()
	userAccountRepository := userAccount.NewUserAccountDao(sqlHandler)
	accountService := account3.NewAccountService(shardService, userAccountRepository)
	return accountService
}

func InitializeShardService() shard.ShardService {
	sqlHandler := database.NewDB()
	commonShardRepository := commonShard.NewCommonShardDao(sqlHandler)
	shardService := shard.NewShardService(commonShardRepository)
	return shardService
}

func InitializeTransactionService() transaction.TransactionService {
	sqlHandler := database.NewDB()
	commonTransactionRepository := commonTransaction.NewCommonTransactionDao(sqlHandler)
	masterTransactionRepository := masterTransaction.NewMasterTransactionDao(sqlHandler)
	userTransactionRepository := masterTransaction2.NewUserTransactionDao(sqlHandler)
	transactionService := transaction.NewTransactionService(commonTransactionRepository, masterTransactionRepository, userTransactionRepository)
	return transactionService
}