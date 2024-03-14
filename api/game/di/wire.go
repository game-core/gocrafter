//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/configs/database"

	accountHandler "github.com/game-core/gocrafter/api/game/presentation/handler/account"
	friendHandler "github.com/game-core/gocrafter/api/game/presentation/handler/friend"
	idleBonusHandler "github.com/game-core/gocrafter/api/game/presentation/handler/idleBonus"
	loginBonusHandler "github.com/game-core/gocrafter/api/game/presentation/handler/loginBonus"
	profileHandler "github.com/game-core/gocrafter/api/game/presentation/handler/profile"
	authInterceptor "github.com/game-core/gocrafter/api/game/presentation/interceptor/auth"
	accountUsecase "github.com/game-core/gocrafter/api/game/usecase/account"
	friendUsecase "github.com/game-core/gocrafter/api/game/usecase/friend"
	loginBonusUsecase "github.com/game-core/gocrafter/api/game/usecase/loginBonus"
	profileUsecase "github.com/game-core/gocrafter/api/game/usecase/profile"
	accountService "github.com/game-core/gocrafter/pkg/domain/model/account"
	friendService "github.com/game-core/gocrafter/pkg/domain/model/friend"
	itemService "github.com/game-core/gocrafter/pkg/domain/model/item"
	loginBonusService "github.com/game-core/gocrafter/pkg/domain/model/loginBonus"
	profileService "github.com/game-core/gocrafter/pkg/domain/model/profile"
	shardService "github.com/game-core/gocrafter/pkg/domain/model/shard"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
	commonShardDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonShard"
	commonTransactionDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonTransaction"
	masterIdleBonusDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonus"
	masterIdleBonusEventDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusEvent"
	masterIdleBonusItemDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusItem"
	masterIdleBonusScheduleDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusSchedule"
	masterItemDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterItem"
	masterLoginBonusDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonus"
	masterLoginBonusEventDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusEvent"
	masterLoginBonusItemDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusItem"
	masterLoginBonusScheduleDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusSchedule"
	masterTransactionDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterTransaction"
	userAccountDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userAccount"
	userFriendDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userFriend"
	userIdleBonusDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userIdleBonus"
	userItemBoxDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userItemBox"
	userLoginBonusDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userLoginBonus"
	userProfileDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userProfile"
	userTransactionDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userTransaction"
)

func InitializeAuthInterceptor() authInterceptor.AuthInterceptor {
	wire.Build(
		authInterceptor.NewAuthInterceptor,
	)
	return nil
}

func InitializeAccountHandler() accountHandler.AccountHandler {
	wire.Build(
		accountHandler.NewAccountHandler,
		InitializeAccountUsecase,
	)
	return nil
}

func InitializeFriendHandler() friendHandler.FriendHandler {
	wire.Build(
		friendHandler.NewFriendHandler,
		InitializeFriendUsecase,
	)
	return nil
}

func InitializeIdleBonusHandler() idleBonusHandler.IdleBonusHandler {
	wire.Build(
		idleBonusHandler.NewIdleBonusHandler,
		InitializeIdleBonusUsecase,
	)
	return nil
}

func InitializeLoginBonusHandler() loginBonusHandler.LoginBonusHandler {
	wire.Build(
		loginBonusHandler.NewLoginBonusHandler,
		InitializeLoginBonusUsecase,
	)
	return nil
}

func InitializeProfileHandler() profileHandler.ProfileHandler {
	wire.Build(
		profileHandler.NewProfileHandler,
		InitializeProfileUsecase,
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

func InitializeFriendUsecase() friendUsecase.FriendUsecase {
	wire.Build(
		friendUsecase.NewFriendUsecase,
		InitializeFriendService,
		InitializeTransactionService,
	)
	return nil
}

func InitializeLoginBonusUsecase() loginBonusUsecase.LoginBonusUsecase {
	wire.Build(
		loginBonusUsecase.NewLoginBonusUsecase,
		InitializeLoginBonusService,
		InitializeTransactionService,
	)
	return nil
}

func InitializeProfileUsecase() profileUsecase.ProfileUsecase {
	wire.Build(
		profileUsecase.NewProfileUsecase,
		InitializeProfileService,
		InitializeTransactionService,
	)
	return nil
}

func InitializeAccountService() accountService.AccountService {
	wire.Build(
		database.NewDB,
		accountService.NewAccountService,
		InitializeShardService,
		userAccountDao.NewUserAccountDao,
	)
	return nil
}

func InitializeFriendService() friendService.FriendService {
	wire.Build(
		database.NewDB,
		friendService.NewFriendService,
		InitializeAccountService,
		userFriendDao.NewUserFriendDao,
	)
	return nil
}

func InitializeIdleBonusService() idleBonusService.IdleBonusService {
	wire.Build(
		database.NewDB,
		idleBonusService.NewIdleBonusService,
		InitializeItemService,
		userIdleBonusDao.NewUserIdleBonusDao,
		masterIdleBonusDao.NewMasterIdleBonusDao,
		masterIdleBonusEventDao.NewMasterIdleBonusEventDao,
		masterIdleBonusItemDao.NewMasterIdleBonusItemDao,
		masterIdleBonusScheduleDao.NewMasterIdleBonusScheduleDao,
	)
	return nil
}

func InitializeItemService() itemService.ItemService {
	wire.Build(
		database.NewDB,
		itemService.NewItemService,
		userItemBoxDao.NewUserItemBoxDao,
		masterItemDao.NewMasterItemDao,
	)
	return nil
}

func InitializeLoginBonusService() loginBonusService.LoginBonusService {
	wire.Build(
		database.NewDB,
		loginBonusService.NewLoginBonusService,
		InitializeItemService,
		userLoginBonusDao.NewUserLoginBonusDao,
		masterLoginBonusDao.NewMasterLoginBonusDao,
		masterLoginBonusEventDao.NewMasterLoginBonusEventDao,
		masterLoginBonusItemDao.NewMasterLoginBonusItemDao,
		masterLoginBonusScheduleDao.NewMasterLoginBonusScheduleDao,
	)
	return nil
}

func InitializeProfileService() profileService.ProfileService {
	wire.Build(
		database.NewDB,
		profileService.NewProfileService,
		userProfileDao.NewUserProfileDao,
	)
	return nil
}

func InitializeShardService() shardService.ShardService {
	wire.Build(
		database.NewDB,
		shardService.NewShardService,
		commonShardDao.NewCommonShardDao,
	)
	return nil
}

func InitializeTransactionService() transactionService.TransactionService {
	wire.Build(
		database.NewDB,
		transactionService.NewTransactionService,
		commonTransactionDao.NewCommonTransactionDao,
		masterTransactionDao.NewMasterTransactionDao,
		userTransactionDao.NewUserTransactionDao,
	)
	return nil
}
