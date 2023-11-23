package router

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/game-core/gocrafter/api/di"
	_ "github.com/game-core/gocrafter/docs/swagger/api"
	"github.com/game-core/gocrafter/log"
)

func Init() {
	// di: wire ./api/di/wire.go
	accountController := di.InitializeAccountController()
	loginRewardController := di.InitializeLoginRewardController()
	accountMiddleware := di.InitializeAccountMiddleware()

	e := echo.New()

	// Swagger
	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	// Log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: log.GenerateApiLog()}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// アカウント関連
	account := e.Group("/account")
	account.POST("/register_account", accountController.RegisterAccount())
	account.POST("/login_account", accountController.LoginAccount())

	accountWithToken := e.Group("/account")
	accountWithToken.Use(accountMiddleware.AccountMiddleware)
	accountWithToken.POST("/check_account", accountController.CheckAccount())

	// ログイン報酬関係
	loginReward := e.Group("/login_reward")
	loginReward.Use(accountMiddleware.AccountMiddleware)
	loginReward.POST("/receive_login_reward", loginRewardController.ReceiveLoginReward())

	e.Logger.Fatal(e.Start(":8000"))
}
