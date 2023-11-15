// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/game-core/gocrafter/config/database"
    "github.com/game-core/gocrafter/api/presentation/middleware"
    dao "github.com/game-core/gocrafter/infra/dao/user"
    accountDao "github.com/game-core/gocrafter/infra/dao/user/account"
    accountController "github.com/game-core/gocrafter/api/presentation/controller/account"	
	accountService "github.com/game-core/gocrafter/domain/service/account"
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
