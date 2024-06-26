//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/configs/database"

	accountHandler "github.com/game-core/gocrafter/api/game/presentation/handler/account"
	friendHandler "github.com/game-core/gocrafter/api/game/presentation/handler/friend"
	healthHandler "github.com/game-core/gocrafter/api/game/presentation/handler/health"
	idleBonusHandler "github.com/game-core/gocrafter/api/game/presentation/handler/idleBonus"
	loginBonusHandler "github.com/game-core/gocrafter/api/game/presentation/handler/loginBonus"
	profileHandler "github.com/game-core/gocrafter/api/game/presentation/handler/profile"
	rankingHandler "github.com/game-core/gocrafter/api/game/presentation/handler/ranking"
	roomHandler "github.com/game-core/gocrafter/api/game/presentation/handler/room"
	authInterceptor "github.com/game-core/gocrafter/api/game/presentation/interceptor/auth"
	accountUsecase "github.com/game-core/gocrafter/api/game/usecase/account"
	friendUsecase "github.com/game-core/gocrafter/api/game/usecase/friend"
	healthUsecase "github.com/game-core/gocrafter/api/game/usecase/health"
	idleBonusUsecase "github.com/game-core/gocrafter/api/game/usecase/idleBonus"
	loginBonusUsecase "github.com/game-core/gocrafter/api/game/usecase/loginBonus"
	profileUsecase "github.com/game-core/gocrafter/api/game/usecase/profile"
	rankingUsecase "github.com/game-core/gocrafter/api/game/usecase/ranking"
	roomUsecase "github.com/game-core/gocrafter/api/game/usecase/room"
	accountService "github.com/game-core/gocrafter/pkg/domain/model/account"
	actionService "github.com/game-core/gocrafter/pkg/domain/model/action"
	configService "github.com/game-core/gocrafter/pkg/domain/model/config"
	eventService "github.com/game-core/gocrafter/pkg/domain/model/event"
	friendService "github.com/game-core/gocrafter/pkg/domain/model/friend"
	healthService "github.com/game-core/gocrafter/pkg/domain/model/health"
	idleBonusService "github.com/game-core/gocrafter/pkg/domain/model/idleBonus"
	itemService "github.com/game-core/gocrafter/pkg/domain/model/item"
	loginBonusService "github.com/game-core/gocrafter/pkg/domain/model/loginBonus"
	profileService "github.com/game-core/gocrafter/pkg/domain/model/profile"
	rankingService "github.com/game-core/gocrafter/pkg/domain/model/ranking"
	rarityService "github.com/game-core/gocrafter/pkg/domain/model/rarity"
	resourceService "github.com/game-core/gocrafter/pkg/domain/model/resource"
	roomService "github.com/game-core/gocrafter/pkg/domain/model/room"
	shardService "github.com/game-core/gocrafter/pkg/domain/model/shard"
	transactionService "github.com/game-core/gocrafter/pkg/domain/model/transaction"
	commonRankingRoomMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRankingRoom"
	commonRankingWorldMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRankingWorld"
	commonRoomMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRoom"
	commonRoomUserMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRoomUser"
	commonShardMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonShard"
	commonTransactionMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonTransaction"
	masterActionMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterAction"
	masterActionRunMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionRun"
	masterActionStepMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionStep"
	masterActionTriggerMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionTrigger"
	masterConfigMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterConfig"
	masterEventMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterEvent"
	masterIdleBonusMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonus"
	masterIdleBonusEventMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusEvent"
	masterIdleBonusItemMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusItem"
	masterIdleBonusScheduleMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusSchedule"
	masterItemMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterItem"
	masterLoginBonusMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonus"
	masterLoginBonusEventMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusEvent"
	masterLoginBonusItemMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusItem"
	masterLoginBonusScheduleMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusSchedule"
	masterRankingMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterRanking"
	masterRankingEventMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterRankingEvent"
	masterRankingScopeMysqlDao "github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterRankingScope"
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
	userAccountTokenRedisDao "github.com/game-core/gocrafter/pkg/infrastructure/redis/user/userAccountToken"
	userTransactionRedisDao "github.com/game-core/gocrafter/pkg/infrastructure/redis/user/userTransaction"
)

func InitializeAuthInterceptor() authInterceptor.AuthInterceptor {
	wire.Build(
		authInterceptor.NewAuthInterceptor,
		InitializeAccountService,
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

func InitializeHealthHandler() healthHandler.HealthHandler {
	wire.Build(
		healthHandler.NewHealthHandler,
		InitializeHealthUsecase,
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

func InitializeRankingHandler() rankingHandler.RankingHandler {
	wire.Build(
		rankingHandler.NewRankingHandler,
		InitializeRankingUsecase,
	)
	return nil
}

func InitializeRoomHandler() roomHandler.RoomHandler {
	wire.Build(
		roomHandler.NewRoomHandler,
		InitializeRoomUsecase,
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

func InitializeHealthUsecase() healthUsecase.HealthUsecase {
	wire.Build(
		healthUsecase.NewHealthUsecase,
		InitializeHealthService,
	)
	return nil
}

func InitializeIdleBonusUsecase() idleBonusUsecase.IdleBonusUsecase {
	wire.Build(
		idleBonusUsecase.NewIdleBonusUsecase,
		InitializeIdleBonusService,
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

func InitializeRankingUsecase() rankingUsecase.RankingUsecase {
	wire.Build(
		rankingUsecase.NewRankingUsecase,
		InitializeRankingService,
		InitializeTransactionService,
	)
	return nil
}

func InitializeRoomUsecase() roomUsecase.RoomUsecase {
	wire.Build(
		roomUsecase.NewRoomUsecase,
		InitializeRoomService,
		InitializeTransactionService,
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
		userAccountTokenRedisDao.NewUserAccountTokenDao,
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

func InitializeConfigService() configService.ConfigService {
	wire.Build(
		database.NewMysql,
		configService.NewConfigService,
		masterConfigMysqlDao.NewMasterConfigDao,
	)
	return nil
}

func InitializeEventService() eventService.EventService {
	wire.Build(
		database.NewMysql,
		eventService.NewEventService,
		masterEventMysqlDao.NewMasterEventDao,
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

func InitializeHealthService() healthService.HealthService {
	wire.Build(
		healthService.NewHealthService,
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

func InitializeRankingService() rankingService.RankingService {
	wire.Build(
		database.NewMysql,
		rankingService.NewRankingService,
		InitializeRoomService,
		commonRankingRoomMysqlDao.NewCommonRankingRoomDao,
		commonRankingWorldMysqlDao.NewCommonRankingWorldDao,
		masterRankingMysqlDao.NewMasterRankingDao,
		masterRankingEventMysqlDao.NewMasterRankingEventDao,
		masterRankingScopeMysqlDao.NewMasterRankingScopeDao,
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

func InitializeRoomService() roomService.RoomService {
	wire.Build(
		database.NewMysql,
		roomService.NewRoomService,
		InitializeConfigService,
		InitializeFriendService,
		commonRoomMysqlDao.NewCommonRoomDao,
		commonRoomUserMysqlDao.NewCommonRoomUserDao,
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
