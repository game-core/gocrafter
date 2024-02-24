//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/configs/database"

	accountHandler "github.com/game-core/gocrafter/api/game/presentation/handler/account"
	accountUsecase "github.com/game-core/gocrafter/api/game/usecase/account"
	accountService "github.com/game-core/gocrafter/pkg/domain/model/account"
	shardService "github.com/game-core/gocrafter/pkg/domain/model/shard"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
	commonShardDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonShard"
	commonTransactionDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonTransaction"
	masterTransactionDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterTransaction"
	userAccountDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userAccount"
	userTransactionDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userTransaction"
)

func InitializeAccountHandler() accountHandler.AccountHandler {
	wire.Build(
		accountHandler.NewAccountHandler,
		InitializeAccountUsecase,
	)
	return nil
}

func InitializeAccountUsecase() accountUsecase.AccountUsecase {
	wire.Build(
		accountUsecase.NewAccountUsecase,
		InitializeAccountService,
		InitializeTransactionService,
	)
	return nil
}

func InitializeAccountService() accountService.AccountService {
	wire.Build(
		accountService.NewAccountService,
		InitializeShardService,
		database.NewDB,
		userAccountDao.NewUserAccountDao,
	)
	return nil
}

func InitializeShardService() shardService.ShardService {
	wire.Build(
		shardService.NewShardService,
		database.NewDB,
		commonShardDao.NewCommonShardDao,
	)
	return nil
}

func InitializeTransactionService() transactionService.TransactionService {
	wire.Build(
		transactionService.NewTransactionService,
		database.NewDB,
		commonTransactionDao.NewCommonTransactionDao,
		masterTransactionDao.NewMasterTransactionDao,
		userTransactionDao.NewUserTransactionDao,
	)
	return nil
}
