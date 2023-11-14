package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/game-core/gocrafter/api/di"
	"github.com/game-core/gocrafter/log"
	_ "github.com/game-core/gocrafter/docs/swagger/api"
)

func Init() {
	// di: wire ./api/di/wire.go
	exampleController := di.InitializeExampleController()

	e := echo.New()

	// Swagger
	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	// Log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ Output: log.GenerateApiLog() }))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	example := e.Group("/example")
	example.GET("/:exampleKey/get_example", exampleController.GetExample())

	e.Logger.Fatal(e.Start(":8000"))
}
