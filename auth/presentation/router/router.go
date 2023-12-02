package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "gorm.io/driver/mysql"

	"github.com/game-core/gocrafter/auth/di"
	_ "github.com/game-core/gocrafter/docs/swagger/auth"
	"github.com/game-core/gocrafter/log"
)

func Init() {
	// di: wire ./api/di/wire.go
	accountController := di.InitializeAccountController()
	accountMiddleware := di.InitializeAccountMiddleware()

	e := echo.New()

	// Swagger
	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	// Log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: log.GenerateAuthLog()}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// アカウント関連
	account := e.Group("/account")
	account.POST("/register_account", accountController.RegisterAccount())
	account.POST("/login_account", accountController.LoginAccount())

	accountWithToken := e.Group("/account")
	accountWithToken.Use(accountMiddleware.AccountMiddleware)
	accountWithToken.POST("/check_account", accountController.CheckAccount())

	e.Logger.Fatal(e.Start(":8000"))
}
