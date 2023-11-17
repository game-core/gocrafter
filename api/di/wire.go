//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	accountController "github.com/game-core/gocrafter/api/presentation/controller/account"
	"github.com/game-core/gocrafter/api/presentation/middleware"
	"github.com/game-core/gocrafter/config/database"
	accountService "github.com/game-core/gocrafter/domain/service/account"
	dao "github.com/game-core/gocrafter/infra/dao/user"
	accountDao "github.com/game-core/gocrafter/infra/dao/user/account"
)

func InitializeAccountMiddleware() middleware.AccountMiddleware {
	wire.Build(
		middleware.NewAccountMiddleware,
	)
	return nil
}

func InitializeAccountController() accountController.AccountController {
	wire.Build(
		database.NewDB,
		dao.NewTransactionDao,
		accountDao.NewAccountDao,
		accountService.NewAccountService,
		accountController.NewAccountController,
	)

	return nil
}
