//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/config/database"
	authDao "github.com/game-core/gocrafter/infra/dao/auth"

	accountMiddleware "github.com/game-core/gocrafter/api/presentation/middleware/account"
	accountController "github.com/game-core/gocrafter/auth/presentation/controller/account"
	accountService "github.com/game-core/gocrafter/domain/service/auth/account"
	accountDao "github.com/game-core/gocrafter/infra/dao/auth/account"
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
		authDao.NewTransactionDao,
	)

	return nil
}
