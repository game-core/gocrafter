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
	"github.com/game-core/gocrafter/api/game/presentation/interceptor/auth"
	account2 "github.com/game-core/gocrafter/api/game/usecase/account"
	friend2 "github.com/game-core/gocrafter/api/game/usecase/friend"
	idleBonus2 "github.com/game-core/gocrafter/api/game/usecase/idleBonus"
	loginBonus2 "github.com/game-core/gocrafter/api/game/usecase/loginBonus"
	profile2 "github.com/game-core/gocrafter/api/game/usecase/profile"
	"github.com/game-core/gocrafter/configs/database"
	account3 "github.com/game-core/gocrafter/pkg/domain/model/account"
	"github.com/game-core/gocrafter/pkg/domain/model/action"
	friend3 "github.com/game-core/gocrafter/pkg/domain/model/friend"
	idleBonus3 "github.com/game-core/gocrafter/pkg/domain/model/idleBonus"
	"github.com/game-core/gocrafter/pkg/domain/model/item"
	loginBonus3 "github.com/game-core/gocrafter/pkg/domain/model/loginBonus"
	profile3 "github.com/game-core/gocrafter/pkg/domain/model/profile"
	"github.com/game-core/gocrafter/pkg/domain/model/rarity"
	"github.com/game-core/gocrafter/pkg/domain/model/resource"
	"github.com/game-core/gocrafter/pkg/domain/model/shard"
	"github.com/game-core/gocrafter/pkg/domain/model/transaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonShard"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/commonTransaction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterAction"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionRun"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionStep"
	"github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/masterActionTrigger"
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

func InitializeAccountService() account3.AccountService {
	shardService := InitializeShardService()
	sqlHandler := database.NewMysql()
	userAccountRepository := userAccount.NewUserAccountDao(sqlHandler)
	accountService := account3.NewAccountService(shardService, userAccountRepository)
	return accountService
}

func InitializeActionService() action.ActionService {
	sqlHandler := database.NewMysql()
	masterActionRepository := masterAction.NewMasterActionDao(sqlHandler)
	masterActionRunRepository := masterActionRun.NewMasterActionRunDao(sqlHandler)
	masterActionStepRepository := masterActionStep.NewMasterActionStepDao(sqlHandler)
	masterActionTriggerRepository := masterActionTrigger.NewMasterActionTriggerDao(sqlHandler)
	userActionRepository := userAction.NewUserActionDao(sqlHandler)
	actionService := action.NewActionService(masterActionRepository, masterActionRunRepository, masterActionStepRepository, masterActionTriggerRepository, userActionRepository)
	return actionService
}

func InitializeFriendService() friend3.FriendService {
	accountService := InitializeAccountService()
	sqlHandler := database.NewMysql()
	userFriendRepository := userFriend.NewUserFriendDao(sqlHandler)
	friendService := friend3.NewFriendService(accountService, userFriendRepository)
	return friendService
}

func InitializeIdleBonusService() idleBonus3.IdleBonusService {
	itemService := InitializeItemService()
	sqlHandler := database.NewMysql()
	userIdleBonusRepository := userIdleBonus.NewUserIdleBonusDao(sqlHandler)
	masterIdleBonusRepository := masterIdleBonus.NewMasterIdleBonusDao(sqlHandler)
	masterIdleBonusEventRepository := masterIdleBonusEvent.NewMasterIdleBonusEventDao(sqlHandler)
	masterIdleBonusItemRepository := masterIdleBonusItem.NewMasterIdleBonusItemDao(sqlHandler)
	masterIdleBonusScheduleRepository := masterIdleBonusSchedule.NewMasterIdleBonusScheduleDao(sqlHandler)
	idleBonusService := idleBonus3.NewIdleBonusService(itemService, userIdleBonusRepository, masterIdleBonusRepository, masterIdleBonusEventRepository, masterIdleBonusItemRepository, masterIdleBonusScheduleRepository)
	return idleBonusService
}

func InitializeItemService() item.ItemService {
	sqlHandler := database.NewMysql()
	userItemBoxRepository := userItemBox.NewUserItemBoxDao(sqlHandler)
	masterItemRepository := masterItem.NewMasterItemDao(sqlHandler)
	itemService := item.NewItemService(userItemBoxRepository, masterItemRepository)
	return itemService
}

func InitializeLoginBonusService() loginBonus3.LoginBonusService {
	itemService := InitializeItemService()
	sqlHandler := database.NewMysql()
	userLoginBonusRepository := userLoginBonus.NewUserLoginBonusDao(sqlHandler)
	masterLoginBonusRepository := masterLoginBonus.NewMasterLoginBonusDao(sqlHandler)
	masterLoginBonusEventRepository := masterLoginBonusEvent.NewMasterLoginBonusEventDao(sqlHandler)
	masterLoginBonusItemRepository := masterLoginBonusItem.NewMasterLoginBonusItemDao(sqlHandler)
	masterLoginBonusScheduleRepository := masterLoginBonusSchedule.NewMasterLoginBonusScheduleDao(sqlHandler)
	loginBonusService := loginBonus3.NewLoginBonusService(itemService, userLoginBonusRepository, masterLoginBonusRepository, masterLoginBonusEventRepository, masterLoginBonusItemRepository, masterLoginBonusScheduleRepository)
	return loginBonusService
}

func InitializeProfileService() profile3.ProfileService {
	sqlHandler := database.NewMysql()
	userProfileRepository := userProfile.NewUserProfileDao(sqlHandler)
	profileService := profile3.NewProfileService(userProfileRepository)
	return profileService
}

func InitializeRarityService() rarity.RarityService {
	sqlHandler := database.NewMysql()
	masterRarityRepository := masterRarity.NewMasterRarityDao(sqlHandler)
	rarityService := rarity.NewRarityService(masterRarityRepository)
	return rarityService
}

func InitializeResourceService() resource.ResourceService {
	sqlHandler := database.NewMysql()
	masterResourceRepository := masterResource.NewMasterResourceDao(sqlHandler)
	resourceService := resource.NewResourceService(masterResourceRepository)
	return resourceService
}

func InitializeShardService() shard.ShardService {
	sqlHandler := database.NewMysql()
	commonShardRepository := commonShard.NewCommonShardDao(sqlHandler)
	shardService := shard.NewShardService(commonShardRepository)
	return shardService
}

func InitializeTransactionService() transaction.TransactionService {
	sqlHandler := database.NewMysql()
	commonTransactionRepository := commonTransaction.NewCommonTransactionDao(sqlHandler)
	masterTransactionRepository := masterTransaction.NewMasterTransactionDao(sqlHandler)
	userTransactionRepository := masterTransaction2.NewUserTransactionDao(sqlHandler)
	transactionService := transaction.NewTransactionService(commonTransactionRepository, masterTransactionRepository, userTransactionRepository)
	return transactionService
}
