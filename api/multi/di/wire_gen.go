// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-core/gocrafter/api/multi/presentation/interceptor/auth"
	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/account"
	"github.com/game-core/gocrafter/pkg/domain/model/action"
	"github.com/game-core/gocrafter/pkg/domain/model/config"
	"github.com/game-core/gocrafter/pkg/domain/model/event"
	"github.com/game-core/gocrafter/pkg/domain/model/friend"
	"github.com/game-core/gocrafter/pkg/domain/model/health"
	"github.com/game-core/gocrafter/pkg/domain/model/idleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/item"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/profile"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking"
	"github.com/game-core/gocrafter/pkg/domain/model/rarity"
	"github.com/game-core/gocrafter/pkg/domain/model/resource"
	"github.com/game-core/gocrafter/pkg/domain/model/room"
	"github.com/game-core/gocrafter/pkg/domain/model/shard"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRankingRoom"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRankingWorld"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRoom"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRoomUser"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonShard"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonTransaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterAction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionRun"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionStep"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionTrigger"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterConfig"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterEvent"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonus"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusEvent"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusItem"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusSchedule"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterItem"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonus"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusEvent"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusItem"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusSchedule"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterRanking"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterRankingEvent"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterRankingScope"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterRarity"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterResource"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterTransaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userAccount"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userAction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userFriend"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userIdleBonus"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userItemBox"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userLoginBonus"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userProfile"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userTransaction"
	userAccount2 "github.com/game-core/gocrafter/pkg/infrastructure/redis/user/userAccount"
	"github.com/game-core/gocrafter/pkg/infrastructure/redis/user/userAccountToken"
	masterTransaction2 "github.com/game-core/gocrafter/pkg/infrastructure/redis/user/userTransaction"
)

// Injectors from wire.go:

func InitializeAuthInterceptor() auth.AuthInterceptor {
	authInterceptor := auth.NewAuthInterceptor()
	return authInterceptor
}

func InitializeAccountService() account.AccountService {
	shardService := InitializeShardService()
	mysqlHandler := database.NewMysql()
	userAccountMysqlRepository := userAccount.NewUserAccountDao(mysqlHandler)
	redisHandler := database.NewRedis()
	userAccountRedisRepository := userAccount2.NewUserAccountDao(redisHandler)
	userAccountTokenRedisRepository := userAccountToken.NewUserAccountTokenDao(redisHandler)
	accountService := account.NewAccountService(shardService, userAccountMysqlRepository, userAccountRedisRepository, userAccountTokenRedisRepository)
	return accountService
}

func InitializeActionService() action.ActionService {
	mysqlHandler := database.NewMysql()
	masterActionMysqlRepository := masterAction.NewMasterActionDao(mysqlHandler)
	masterActionRunMysqlRepository := masterActionRun.NewMasterActionRunDao(mysqlHandler)
	masterActionStepMysqlRepository := masterActionStep.NewMasterActionStepDao(mysqlHandler)
	masterActionTriggerMysqlRepository := masterActionTrigger.NewMasterActionTriggerDao(mysqlHandler)
	userActionMysqlRepository := userAction.NewUserActionDao(mysqlHandler)
	actionService := action.NewActionService(masterActionMysqlRepository, masterActionRunMysqlRepository, masterActionStepMysqlRepository, masterActionTriggerMysqlRepository, userActionMysqlRepository)
	return actionService
}

func InitializeConfigService() config.ConfigService {
	mysqlHandler := database.NewMysql()
	masterConfigMysqlRepository := masterConfig.NewMasterConfigDao(mysqlHandler)
	configService := config.NewConfigService(masterConfigMysqlRepository)
	return configService
}

func InitializeEventService() event.EventService {
	mysqlHandler := database.NewMysql()
	masterEventMysqlRepository := masterEvent.NewMasterEventDao(mysqlHandler)
	eventService := event.NewEventService(masterEventMysqlRepository)
	return eventService
}

func InitializeFriendService() friend.FriendService {
	accountService := InitializeAccountService()
	mysqlHandler := database.NewMysql()
	userFriendMysqlRepository := userFriend.NewUserFriendDao(mysqlHandler)
	friendService := friend.NewFriendService(accountService, userFriendMysqlRepository)
	return friendService
}

func InitializeHealthService() health.HealthService {
	healthService := health.NewHealthService()
	return healthService
}

func InitializeIdleBonusService() idleBonus.IdleBonusService {
	itemService := InitializeItemService()
	mysqlHandler := database.NewMysql()
	userIdleBonusMysqlRepository := userIdleBonus.NewUserIdleBonusDao(mysqlHandler)
	masterIdleBonusMysqlRepository := masterIdleBonus.NewMasterIdleBonusDao(mysqlHandler)
	masterIdleBonusEventMysqlRepository := masterIdleBonusEvent.NewMasterIdleBonusEventDao(mysqlHandler)
	masterIdleBonusItemMysqlRepository := masterIdleBonusItem.NewMasterIdleBonusItemDao(mysqlHandler)
	masterIdleBonusScheduleMysqlRepository := masterIdleBonusSchedule.NewMasterIdleBonusScheduleDao(mysqlHandler)
	idleBonusService := idleBonus.NewIdleBonusService(itemService, userIdleBonusMysqlRepository, masterIdleBonusMysqlRepository, masterIdleBonusEventMysqlRepository, masterIdleBonusItemMysqlRepository, masterIdleBonusScheduleMysqlRepository)
	return idleBonusService
}

func InitializeItemService() item.ItemService {
	mysqlHandler := database.NewMysql()
	userItemBoxMysqlRepository := userItemBox.NewUserItemBoxDao(mysqlHandler)
	masterItemMysqlRepository := masterItem.NewMasterItemDao(mysqlHandler)
	itemService := item.NewItemService(userItemBoxMysqlRepository, masterItemMysqlRepository)
	return itemService
}

func InitializeLoginBonusService() loginBonus.LoginBonusService {
	itemService := InitializeItemService()
	mysqlHandler := database.NewMysql()
	userLoginBonusMysqlRepository := userLoginBonus.NewUserLoginBonusDao(mysqlHandler)
	masterLoginBonusMysqlRepository := masterLoginBonus.NewMasterLoginBonusDao(mysqlHandler)
	masterLoginBonusEventMysqlRepository := masterLoginBonusEvent.NewMasterLoginBonusEventDao(mysqlHandler)
	masterLoginBonusItemMysqlRepository := masterLoginBonusItem.NewMasterLoginBonusItemDao(mysqlHandler)
	masterLoginBonusScheduleMysqlRepository := masterLoginBonusSchedule.NewMasterLoginBonusScheduleDao(mysqlHandler)
	loginBonusService := loginBonus.NewLoginBonusService(itemService, userLoginBonusMysqlRepository, masterLoginBonusMysqlRepository, masterLoginBonusEventMysqlRepository, masterLoginBonusItemMysqlRepository, masterLoginBonusScheduleMysqlRepository)
	return loginBonusService
}

func InitializeProfileService() profile.ProfileService {
	mysqlHandler := database.NewMysql()
	userProfileMysqlRepository := userProfile.NewUserProfileDao(mysqlHandler)
	profileService := profile.NewProfileService(userProfileMysqlRepository)
	return profileService
}

func InitializeRankingService() ranking.RankingService {
	roomService := InitializeRoomService()
	mysqlHandler := database.NewMysql()
	commonRankingRoomMysqlRepository := commonRankingRoom.NewCommonRankingRoomDao(mysqlHandler)
	commonRankingWorldMysqlRepository := commonRankingWorld.NewCommonRankingWorldDao(mysqlHandler)
	masterRankingMysqlRepository := masterRanking.NewMasterRankingDao(mysqlHandler)
	masterRankingEventMysqlRepository := masterRankingEvent.NewMasterRankingEventDao(mysqlHandler)
	masterRankingScopeMysqlRepository := masterRankingScope.NewMasterRankingScopeDao(mysqlHandler)
	rankingService := ranking.NewRankingService(roomService, commonRankingRoomMysqlRepository, commonRankingWorldMysqlRepository, masterRankingMysqlRepository, masterRankingEventMysqlRepository, masterRankingScopeMysqlRepository)
	return rankingService
}

func InitializeRarityService() rarity.RarityService {
	mysqlHandler := database.NewMysql()
	masterRarityMysqlRepository := masterRarity.NewMasterRarityDao(mysqlHandler)
	rarityService := rarity.NewRarityService(masterRarityMysqlRepository)
	return rarityService
}

func InitializeResourceService() resource.ResourceService {
	mysqlHandler := database.NewMysql()
	masterResourceMysqlRepository := masterResource.NewMasterResourceDao(mysqlHandler)
	resourceService := resource.NewResourceService(masterResourceMysqlRepository)
	return resourceService
}

func InitializeRoomService() room.RoomService {
	configService := InitializeConfigService()
	friendService := InitializeFriendService()
	mysqlHandler := database.NewMysql()
	commonRoomMysqlRepository := commonRoom.NewCommonRoomDao(mysqlHandler)
	commonRoomUserMysqlRepository := commonRoomUser.NewCommonRoomUserDao(mysqlHandler)
	roomService := room.NewRoomService(configService, friendService, commonRoomMysqlRepository, commonRoomUserMysqlRepository)
	return roomService
}

func InitializeShardService() shard.ShardService {
	mysqlHandler := database.NewMysql()
	commonShardMysqlRepository := commonShard.NewCommonShardDao(mysqlHandler)
	shardService := shard.NewShardService(commonShardMysqlRepository)
	return shardService
}

func InitializeTransactionService() transaction.TransactionService {
	mysqlHandler := database.NewMysql()
	commonTransactionMysqlRepository := commonTransaction.NewCommonTransactionDao(mysqlHandler)
	masterTransactionMysqlRepository := masterTransaction.NewMasterTransactionDao(mysqlHandler)
	userTransactionMysqlRepository := userTransaction.NewUserTransactionDao(mysqlHandler)
	redisHandler := database.NewRedis()
	userTransactionRedisRepository := masterTransaction2.NewUserTransactionDao(redisHandler)
	transactionService := transaction.NewTransactionService(commonTransactionMysqlRepository, masterTransactionMysqlRepository, userTransactionMysqlRepository, userTransactionRedisRepository)
	return transactionService
}
