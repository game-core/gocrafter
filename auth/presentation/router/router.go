package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "gorm.io/driver/mysql"

	_ "github.com/game-core/gocrafter/docs/swagger/auth"
	"github.com/game-core/gocrafter/log"
)

func Init() {
	// di: wire ./api/di/wire.go

	e := echo.New()

	// Swagger
	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	// Log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: log.GenerateApiLog()}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8000"))
}
