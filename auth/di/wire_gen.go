// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-core/gocrafter/api/presentation/middleware/account"
	account2 "github.com/game-core/gocrafter/auth/presentation/controller/account"
	"github.com/game-core/gocrafter/config/database"
	account3 "github.com/game-core/gocrafter/domain/service/auth/account"
	account4 "github.com/game-core/gocrafter/infra/dao/auth/account"
	"github.com/game-core/gocrafter/infra/dao/auth/transaction"
)

// Injectors from wire.go:

func InitializeAccountMiddleware() account.AccountMiddleware {
	accountMiddleware := account.NewAccountMiddleware()
	return accountMiddleware
}

func InitializeAccountController() account2.AccountController {
	accountService := InitializeAccountService()
	accountController := account2.NewAccountController(accountService)
	return accountController
}

func InitializeAccountService() account3.AccountService {
	sqlHandler := database.NewDB()
	transactionRepository := config.NewTransactionDao(sqlHandler)
	accountRepository := account4.NewAccountDao(sqlHandler)
	accountService := account3.NewAccountService(transactionRepository, accountRepository)
	return accountService
}
