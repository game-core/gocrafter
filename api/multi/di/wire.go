//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/configs/database"

	authInterceptor "github.com/game-core/gocrafter/api/multi/presentation/interceptor/auth"
	accountService "github.com/game-core/gocrafter/pkg/domain/model/account"
	actionService "github.com/game-core/gocrafter/pkg/domain/model/action"
	friendService "github.com/game-core/gocrafter/pkg/domain/model/friend"
	idleBonusService "github.com/game-core/gocrafter/pkg/domain/model/idleBonus"
	itemService "github.com/game-core/gocrafter/pkg/domain/model/item"
	loginBonusService "github.com/game-core/gocrafter/pkg/domain/model/loginBonus"
	profileService "github.com/game-core/gocrafter/pkg/domain/model/profile"
	rarityService "github.com/game-core/gocrafter/pkg/domain/model/rarity"
	resourceService "github.com/game-core/gocrafter/pkg/domain/model/resource"
	shardService "github.com/game-core/gocrafter/pkg/domain/model/shard"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
	commonShardMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonShard"
	commonTransactionMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonTransaction"
	masterActionMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterAction"
	masterActionRunMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionRun"
	masterActionStepMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionStep"
	masterActionTriggerMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionTrigger"
	masterIdleBonusMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonus"
	masterIdleBonusEventMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusEvent"
	masterIdleBonusItemMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusItem"
	masterIdleBonusScheduleMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusSchedule"
	masterItemMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterItem"
	masterLoginBonusMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonus"
	masterLoginBonusEventMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusEvent"
	masterLoginBonusItemMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusItem"
	masterLoginBonusScheduleMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusSchedule"
	masterRarityMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterRarity"
	masterResourceMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterResource"
	masterTransactionMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterTransaction"
	userAccountMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userAccount"
	userActionMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userAction"
	userFriendMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userFriend"
	userIdleBonusMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userIdleBonus"
	userItemBoxMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userItemBox"
	userLoginBonusMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userLoginBonus"
	userProfileMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userProfile"
	userTransactionMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userTransaction"
	userAccountRedisDao "github.com/game-core/gocrafter/pkg/infrastructure/redis/user/userAccount"
	userTransactionRedisDao "github.com/game-core/gocrafter/pkg/infrastructure/redis/user/userTransaction"
)

func InitializeAuthInterceptor() authInterceptor.AuthInterceptor {
	wire.Build(
		authInterceptor.NewAuthInterceptor,
	)
	return nil
}

func InitializeAccountService() accountService.AccountService {
	wire.Build(
		database.NewMysql,
		database.NewRedis,
		accountService.NewAccountService,
		InitializeShardService,
		userAccountMysqlDao.NewUserAccountDao,
		userAccountRedisDao.NewUserAccountDao,
	)
	return nil
}

func InitializeActionService() actionService.ActionService {
	wire.Build(
		database.NewMysql,
		actionService.NewActionService,
		masterActionMysqlDao.NewMasterActionDao,
		masterActionRunMysqlDao.NewMasterActionRunDao,
		masterActionStepMysqlDao.NewMasterActionStepDao,
		masterActionTriggerMysqlDao.NewMasterActionTriggerDao,
		userActionMysqlDao.NewUserActionDao,
	)
	return nil
}

func InitializeFriendService() friendService.FriendService {
	wire.Build(
		database.NewMysql,
		friendService.NewFriendService,
		InitializeAccountService,
		userFriendMysqlDao.NewUserFriendDao,
	)
	return nil
}

func InitializeIdleBonusService() idleBonusService.IdleBonusService {
	wire.Build(
		database.NewMysql,
		idleBonusService.NewIdleBonusService,
		InitializeItemService,
		userIdleBonusMysqlDao.NewUserIdleBonusDao,
		masterIdleBonusMysqlDao.NewMasterIdleBonusDao,
		masterIdleBonusEventMysqlDao.NewMasterIdleBonusEventDao,
		masterIdleBonusItemMysqlDao.NewMasterIdleBonusItemDao,
		masterIdleBonusScheduleMysqlDao.NewMasterIdleBonusScheduleDao,
	)
	return nil
}

func InitializeItemService() itemService.ItemService {
	wire.Build(
		database.NewMysql,
		itemService.NewItemService,
		userItemBoxMysqlDao.NewUserItemBoxDao,
		masterItemMysqlDao.NewMasterItemDao,
	)
	return nil
}

func InitializeLoginBonusService() loginBonusService.LoginBonusService {
	wire.Build(
		database.NewMysql,
		loginBonusService.NewLoginBonusService,
		InitializeItemService,
		userLoginBonusMysqlDao.NewUserLoginBonusDao,
		masterLoginBonusMysqlDao.NewMasterLoginBonusDao,
		masterLoginBonusEventMysqlDao.NewMasterLoginBonusEventDao,
		masterLoginBonusItemMysqlDao.NewMasterLoginBonusItemDao,
		masterLoginBonusScheduleMysqlDao.NewMasterLoginBonusScheduleDao,
	)
	return nil
}

func InitializeProfileService() profileService.ProfileService {
	wire.Build(
		database.NewMysql,
		profileService.NewProfileService,
		userProfileMysqlDao.NewUserProfileDao,
	)
	return nil
}

func InitializeRarityService() rarityService.RarityService {
	wire.Build(
		database.NewMysql,
		rarityService.NewRarityService,
		masterRarityMysqlDao.NewMasterRarityDao,
	)
	return nil
}

func InitializeResourceService() resourceService.ResourceService {
	wire.Build(
		database.NewMysql,
		resourceService.NewResourceService,
		masterResourceMysqlDao.NewMasterResourceDao,
	)
	return nil
}

func InitializeShardService() shardService.ShardService {
	wire.Build(
		database.NewMysql,
		shardService.NewShardService,
		commonShardMysqlDao.NewCommonShardDao,
	)
	return nil
}

func InitializeTransactionService() transactionService.TransactionService {
	wire.Build(
		database.NewMysql,
		database.NewRedis,
		transactionService.NewTransactionService,
		commonTransactionMysqlDao.NewCommonTransactionDao,
		masterTransactionMysqlDao.NewMasterTransactionDao,
		userTransactionMysqlDao.NewUserTransactionDao,
		userTransactionRedisDao.NewUserTransactionDao,
	)
	return nil
}
