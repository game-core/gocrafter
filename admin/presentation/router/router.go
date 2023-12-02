package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "gorm.io/driver/mysql"

	"github.com/game-core/gocrafter/admin/di"
	_ "github.com/game-core/gocrafter/docs/swagger/admin"
	"github.com/game-core/gocrafter/log"
)

func Init() {
	// di: wire ./api/di/wire.go
	exampleController := di.InitializeExampleController()

	e := echo.New()

	// Swagger
	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	// Log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: log.GenerateAdminLog()}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Example関連
	account := e.Group("/example")
	account.POST("/get_example", exampleController.GetExample())

	e.Logger.Fatal(e.Start(":8000"))
}
