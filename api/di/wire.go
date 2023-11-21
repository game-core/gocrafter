//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/config/database"
	userDao "github.com/game-core/gocrafter/infra/dao/user"

	accountController "github.com/game-core/gocrafter/api/presentation/controller/account"
	accountMiddleware "github.com/game-core/gocrafter/api/presentation/middleware/account"
	accountService "github.com/game-core/gocrafter/domain/service/account"
	accountDao "github.com/game-core/gocrafter/infra/dao/user/account"
)

func InitializeAccountMiddleware() accountMiddleware.AccountMiddleware {
	wire.Build(
		accountMiddleware.NewAccountMiddleware,
	)

	return nil
}

func InitializeAccountController() accountController.AccountController {
	wire.Build(
		database.NewDB,
		accountController.NewAccountController,
		accountService.NewAccountService,
		accountDao.NewAccountDao,
		userDao.NewTransactionDao,
	)

	return nil
}
