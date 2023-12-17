// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-core/gocrafter/admin/presentation/controller/example"
	"github.com/game-core/gocrafter/admin/presentation/middleware/account"
	"github.com/game-core/gocrafter/config/database"
	example2 "github.com/game-core/gocrafter/domain/service/admin/example"
	example3 "github.com/game-core/gocrafter/infra/dao/admin/example"
	"github.com/game-core/gocrafter/infra/dao/admin/transaction"
)

// Injectors from wire.go:

func InitializeAccountMiddleware() account.AccountMiddleware {
	accountMiddleware := account.NewAccountMiddleware()
	return accountMiddleware
}

func InitializeExampleController() example.ExampleController {
	exampleService := InitializeExampleService()
	exampleController := example.NewExampleController(exampleService)
	return exampleController
}

func InitializeExampleService() example2.ExampleService {
	sqlHandler := database.NewDB()
	transactionRepository := config.NewTransactionDao(sqlHandler)
	exampleRepository := example3.NewExampleDao(sqlHandler)
	exampleService := example2.NewExampleService(transactionRepository, exampleRepository)
	return exampleService
}
