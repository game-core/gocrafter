//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	accountController "github.com/game-core/gocrafter/api/presentation/controller/account"
	accountMiddleware "github.com/game-core/gocrafter/api/presentation/middleware/account"
	"github.com/game-core/gocrafter/config/database"
	accountService "github.com/game-core/gocrafter/domain/service/account"
	userDao "github.com/game-core/gocrafter/infra/dao/user"
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
