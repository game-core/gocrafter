// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-core/gocrafter/api/game/presentation/handler/account"
	"github.com/game-core/gocrafter/api/game/presentation/handler/friend"
	"github.com/game-core/gocrafter/api/game/presentation/handler/idleBonus"
	"github.com/game-core/gocrafter/api/game/presentation/handler/loginBonus"
	"github.com/game-core/gocrafter/api/game/presentation/handler/profile"
	"github.com/game-core/gocrafter/api/game/presentation/handler/room"
	"github.com/game-core/gocrafter/api/game/presentation/interceptor/auth"
	account2 "github.com/game-core/gocrafter/api/game/usecase/account"
	friend2 "github.com/game-core/gocrafter/api/game/usecase/friend"
	idleBonus2 "github.com/game-core/gocrafter/api/game/usecase/idleBonus"
	loginBonus2 "github.com/game-core/gocrafter/api/game/usecase/loginBonus"
	profile2 "github.com/game-core/gocrafter/api/game/usecase/profile"
	room2 "github.com/game-core/gocrafter/api/game/usecase/room"
	"github.com/game-core/gocrafter/configs/database"
	account3 "github.com/game-core/gocrafter/pkg/domain/model/account"
	"github.com/game-core/gocrafter/pkg/domain/model/action"
	"github.com/game-core/gocrafter/pkg/domain/model/config"
	friend3 "github.com/game-core/gocrafter/pkg/domain/model/friend"
	idleBonus3 "github.com/game-core/gocrafter/pkg/domain/model/idleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/item"
	loginBonus3 "github.com/game-core/gocrafter/pkg/domain/model/loginBonus"
	profile3 "github.com/game-core/gocrafter/pkg/domain/model/profile"
	"github.com/game-core/gocrafter/pkg/domain/model/rarity"
	"github.com/game-core/gocrafter/pkg/domain/model/resource"
	room3 "github.com/game-core/gocrafter/pkg/domain/model/room"
	"github.com/game-core/gocrafter/pkg/domain/model/shard"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRoom"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonRoomUser"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonShard"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonTransaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterAction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionRun"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionStep"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionTrigger"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterConfig"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonus"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusEvent"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusItem"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterIdleBonusSchedule"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterItem"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonus"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusEvent"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusItem"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterLoginBonusSchedule"
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
	masterTransaction2 "github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/userTransaction"
	userAccount2 "github.com/game-core/gocrafter/pkg/infrastructure/redis/user/userAccount"
	masterTransaction3 "github.com/game-core/gocrafter/pkg/infrastructure/redis/user/userTransaction"
)

// Injectors from wire.go:

func InitializeAuthInterceptor() auth.AuthInterceptor {
	authInterceptor := auth.NewAuthInterceptor()
	return authInterceptor
}

func InitializeAccountHandler() account.AccountHandler {
	accountUsecase := InitializeAccountUsecase()
	accountHandler := account.NewAccountHandler(accountUsecase)
	return accountHandler
}

func InitializeFriendHandler() friend.FriendHandler {
	friendUsecase := InitializeFriendUsecase()
	friendHandler := friend.NewFriendHandler(friendUsecase)
	return friendHandler
}

func InitializeIdleBonusHandler() idleBonus.IdleBonusHandler {
	idleBonusUsecase := InitializeIdleBonusUsecase()
	idleBonusHandler := idleBonus.NewIdleBonusHandler(idleBonusUsecase)
	return idleBonusHandler
}

func InitializeLoginBonusHandler() loginBonus.LoginBonusHandler {
	loginBonusUsecase := InitializeLoginBonusUsecase()
	loginBonusHandler := loginBonus.NewLoginBonusHandler(loginBonusUsecase)
	return loginBonusHandler
}

func InitializeProfileHandler() profile.ProfileHandler {
	profileUsecase := InitializeProfileUsecase()
	profileHandler := profile.NewProfileHandler(profileUsecase)
	return profileHandler
}

func InitializeRoomHandler() room.RoomHandler {
	roomUsecase := InitializeRoomUsecase()
	roomHandler := room.NewRoomHandler(roomUsecase)
	return roomHandler
}

func InitializeAccountUsecase() account2.AccountUsecase {
	accountService := InitializeAccountService()
	transactionService := InitializeTransactionService()
	accountUsecase := account2.NewAccountUsecase(accountService, transactionService)
	return accountUsecase
}

func InitializeFriendUsecase() friend2.FriendUsecase {
	friendService := InitializeFriendService()
	transactionService := InitializeTransactionService()
	friendUsecase := friend2.NewFriendUsecase(friendService, transactionService)
	return friendUsecase
}

func InitializeIdleBonusUsecase() idleBonus2.IdleBonusUsecase {
	idleBonusService := InitializeIdleBonusService()
	transactionService := InitializeTransactionService()
	idleBonusUsecase := idleBonus2.NewIdleBonusUsecase(idleBonusService, transactionService)
	return idleBonusUsecase
}

func InitializeLoginBonusUsecase() loginBonus2.LoginBonusUsecase {
	loginBonusService := InitializeLoginBonusService()
	transactionService := InitializeTransactionService()
	loginBonusUsecase := loginBonus2.NewLoginBonusUsecase(loginBonusService, transactionService)
	return loginBonusUsecase
}

func InitializeProfileUsecase() profile2.ProfileUsecase {
	profileService := InitializeProfileService()
	transactionService := InitializeTransactionService()
	profileUsecase := profile2.NewProfileUsecase(profileService, transactionService)
	return profileUsecase
}

func InitializeRoomUsecase() room2.RoomUsecase {
	roomService := InitializeRoomService()
	transactionService := InitializeTransactionService()
	roomUsecase := room2.NewRoomUsecase(roomService, transactionService)
	return roomUsecase
}

func InitializeAccountService() account3.AccountService {
	shardService := InitializeShardService()
	mysqlHandler := database.NewMysql()
	userAccountMysqlRepository := userAccount.NewUserAccountDao(mysqlHandler)
	redisHandler := database.NewRedis()
	userAccountRedisRepository := userAccount2.NewUserAccountDao(redisHandler)
	accountService := account3.NewAccountService(shardService, userAccountMysqlRepository, userAccountRedisRepository)
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

func InitializeFriendService() friend3.FriendService {
	accountService := InitializeAccountService()
	mysqlHandler := database.NewMysql()
	userFriendMysqlRepository := userFriend.NewUserFriendDao(mysqlHandler)
	friendService := friend3.NewFriendService(accountService, userFriendMysqlRepository)
	return friendService
}

func InitializeIdleBonusService() idleBonus3.IdleBonusService {
	itemService := InitializeItemService()
	mysqlHandler := database.NewMysql()
	userIdleBonusMysqlRepository := userIdleBonus.NewUserIdleBonusDao(mysqlHandler)
	masterIdleBonusMysqlRepository := masterIdleBonus.NewMasterIdleBonusDao(mysqlHandler)
	masterIdleBonusEventMysqlRepository := masterIdleBonusEvent.NewMasterIdleBonusEventDao(mysqlHandler)
	masterIdleBonusItemMysqlRepository := masterIdleBonusItem.NewMasterIdleBonusItemDao(mysqlHandler)
	masterIdleBonusScheduleMysqlRepository := masterIdleBonusSchedule.NewMasterIdleBonusScheduleDao(mysqlHandler)
	idleBonusService := idleBonus3.NewIdleBonusService(itemService, userIdleBonusMysqlRepository, masterIdleBonusMysqlRepository, masterIdleBonusEventMysqlRepository, masterIdleBonusItemMysqlRepository, masterIdleBonusScheduleMysqlRepository)
	return idleBonusService
}

func InitializeItemService() item.ItemService {
	mysqlHandler := database.NewMysql()
	userItemBoxMysqlRepository := userItemBox.NewUserItemBoxDao(mysqlHandler)
	masterItemMysqlRepository := masterItem.NewMasterItemDao(mysqlHandler)
	itemService := item.NewItemService(userItemBoxMysqlRepository, masterItemMysqlRepository)
	return itemService
}

func InitializeLoginBonusService() loginBonus3.LoginBonusService {
	itemService := InitializeItemService()
	mysqlHandler := database.NewMysql()
	userLoginBonusMysqlRepository := userLoginBonus.NewUserLoginBonusDao(mysqlHandler)
	masterLoginBonusMysqlRepository := masterLoginBonus.NewMasterLoginBonusDao(mysqlHandler)
	masterLoginBonusEventMysqlRepository := masterLoginBonusEvent.NewMasterLoginBonusEventDao(mysqlHandler)
	masterLoginBonusItemMysqlRepository := masterLoginBonusItem.NewMasterLoginBonusItemDao(mysqlHandler)
	masterLoginBonusScheduleMysqlRepository := masterLoginBonusSchedule.NewMasterLoginBonusScheduleDao(mysqlHandler)
	loginBonusService := loginBonus3.NewLoginBonusService(itemService, userLoginBonusMysqlRepository, masterLoginBonusMysqlRepository, masterLoginBonusEventMysqlRepository, masterLoginBonusItemMysqlRepository, masterLoginBonusScheduleMysqlRepository)
	return loginBonusService
}

func InitializeProfileService() profile3.ProfileService {
	mysqlHandler := database.NewMysql()
	userProfileMysqlRepository := userProfile.NewUserProfileDao(mysqlHandler)
	profileService := profile3.NewProfileService(userProfileMysqlRepository)
	return profileService
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

func InitializeRoomService() room3.RoomService {
	configService := InitializeConfigService()
	friendService := InitializeFriendService()
	mysqlHandler := database.NewMysql()
	commonRoomMysqlRepository := commonRoom.NewCommonRoomDao(mysqlHandler)
	commonRoomUserMysqlRepository := commonRoomUser.NewCommonRoomUserDao(mysqlHandler)
	roomService := room3.NewRoomService(configService, friendService, commonRoomMysqlRepository, commonRoomUserMysqlRepository)
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
	userTransactionMysqlRepository := masterTransaction2.NewUserTransactionDao(mysqlHandler)
	redisHandler := database.NewRedis()
	userTransactionRedisRepository := masterTransaction3.NewUserTransactionDao(redisHandler)
	transactionService := transaction.NewTransactionService(commonTransactionMysqlRepository, masterTransactionMysqlRepository, userTransactionMysqlRepository, userTransactionRedisRepository)
	return transactionService
}
