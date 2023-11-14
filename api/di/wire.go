// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/game-core/gocrafter/config/database"
    "github.com/game-core/gocrafter/infra/dao"
    "github.com/game-core/gocrafter/api/service"	
    "github.com/game-core/gocrafter/api/presentation/controller"
	"github.com/game-core/gocrafter/api/presentation/middleware"
)

// example
func InitializeExampleController() controller.ExampleController {
	wire.Build(
		db.NewDB,
		dao.NewExampleDao,
		service.NewExampleService,
		controller.NewExampleController,
	)

    return nil
}

// user
func InitializeUserMiddleware() middleware.UserMiddleware {
    wire.Build(
		middleware.NewUserMiddleware,
    )
    return nil
}
