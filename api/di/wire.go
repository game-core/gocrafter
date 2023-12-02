//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/config/database"
	configDao "github.com/game-core/gocrafter/infra/dao/config"
	userDao "github.com/game-core/gocrafter/infra/dao/user"

	accountController "github.com/game-core/gocrafter/api/presentation/controller/account"
	loginRewardController "github.com/game-core/gocrafter/api/presentation/controller/loginReward"
	accountMiddleware "github.com/game-core/gocrafter/api/presentation/middleware/account"
	accountService "github.com/game-core/gocrafter/domain/service/api/account"
	eventService "github.com/game-core/gocrafter/domain/service/api/event"
	itemService "github.com/game-core/gocrafter/domain/service/api/item"
	loginRewardService "github.com/game-core/gocrafter/domain/service/api/loginReward"
	shardService "github.com/game-core/gocrafter/domain/service/api/shard"
	shardDao "github.com/game-core/gocrafter/infra/dao/config/shard"
	masterEventDao "github.com/game-core/gocrafter/infra/dao/master/event"
	masterItemDao "github.com/game-core/gocrafter/infra/dao/master/item"
	masterLoginRewardDao "github.com/game-core/gocrafter/infra/dao/master/loginReward"
	accountDao "github.com/game-core/gocrafter/infra/dao/user/account"
	userItemDao "github.com/game-core/gocrafter/infra/dao/user/item"
	userLoginRewardDao "github.com/game-core/gocrafter/infra/dao/user/loginReward"
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

func InitializeLoginRewardController() loginRewardController.LoginRewardController {
	wire.Build(
		loginRewardController.NewLoginRewardController,
		InitializeLoginRewardService,
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

func InitializeLoginRewardService() loginRewardService.LoginRewardService {
	wire.Build(
		database.NewDB,
		loginRewardService.NewLoginRewardService,
		userLoginRewardDao.NewLoginRewardStatusDao,
		masterLoginRewardDao.NewLoginRewardRewardDao,
		masterLoginRewardDao.NewLoginRewardModelDao,
		userDao.NewTransactionDao,
		InitializeEventService,
		InitializeItemService,
	)

	return nil
}

func InitializeEventService() eventService.EventService {
	wire.Build(
		database.NewDB,
		eventService.NewEventService,
		masterEventDao.NewEventDao,
	)

	return nil
}

func InitializeItemService() itemService.ItemService {
	wire.Build(
		database.NewDB,
		itemService.NewItemService,
		masterItemDao.NewItemDao,
		userItemDao.NewItemBoxDao,
		userDao.NewTransactionDao,
	)

	return nil
}
