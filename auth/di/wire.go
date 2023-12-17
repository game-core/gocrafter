//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/config/database"

	accountMiddleware "github.com/game-core/gocrafter/api/presentation/middleware/account"
	accountController "github.com/game-core/gocrafter/auth/presentation/controller/account"
	accountService "github.com/game-core/gocrafter/domain/service/auth/account"
	accountDao "github.com/game-core/gocrafter/infra/dao/auth/account"
	transactionDao "github.com/game-core/gocrafter/infra/dao/auth/transaction"
)

func InitializeAccountMiddleware() accountMiddleware.AccountMiddleware {
	wire.Build(
		accountMiddleware.NewAccountMiddleware,
	)

	return nil
}

func InitializeAccountController() accountController.AccountController {
	wire.Build(
		accountController.NewAccountController,
		InitializeAccountService,
	)

	return nil
}

func InitializeAccountService() accountService.AccountService {
	wire.Build(
		database.NewDB,
		accountService.NewAccountService,
		accountDao.NewAccountDao,
		transactionDao.NewTransactionDao,
	)

	return nil
}
