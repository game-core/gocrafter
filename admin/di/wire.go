//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/config/database"

	exampleController "github.com/game-core/gocrafter/admin/presentation/controller/example"
	accountMiddleware "github.com/game-core/gocrafter/admin/presentation/middleware/account"
	exampleService "github.com/game-core/gocrafter/domain/service/admin/example"
	exampleDao "github.com/game-core/gocrafter/infra/dao/admin/example"
	transactionDao "github.com/game-core/gocrafter/infra/dao/admin/transaction"
)

func InitializeAccountMiddleware() accountMiddleware.AccountMiddleware {
	wire.Build(
		accountMiddleware.NewAccountMiddleware,
	)

	return nil
}

func InitializeExampleController() exampleController.ExampleController {
	wire.Build(
		exampleController.NewExampleController,
		InitializeExampleService,
	)

	return nil
}

func InitializeExampleService() exampleService.ExampleService {
	wire.Build(
		database.NewDB,
		exampleService.NewExampleService,
		exampleDao.NewExampleDao,
		transactionDao.NewTransactionDao,
	)

	return nil
}
