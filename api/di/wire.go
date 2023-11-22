//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/config/database"
	configDao "github.com/game-core/gocrafter/infra/dao/config"
	userDao "github.com/game-core/gocrafter/infra/dao/user"

	accountController "github.com/game-core/gocrafter/api/presentation/controller/account"
	accountMiddleware "github.com/game-core/gocrafter/api/presentation/middleware/account"
	accountService "github.com/game-core/gocrafter/domain/service/account"
	shardService "github.com/game-core/gocrafter/domain/service/shard"
	shardDao "github.com/game-core/gocrafter/infra/dao/config/shard"
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
		accountController.NewAccountController,
		InitializeAccountService,
	)

	return nil
}

func InitializeAccountService() accountService.AccountService {
	wire.Build(
		database.NewDB,
		accountService.NewAccountService,
		InitializeShardService,
		accountDao.NewAccountDao,
		userDao.NewTransactionDao,
	)

	return nil
}

func InitializeShardService() shardService.ShardService {
	wire.Build(
		database.NewDB,
		shardService.NewShardService,
		shardDao.NewShardDao,
		configDao.NewTransactionDao,
	)

	return nil
}
