// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/architecture-template/echo-ddd/config/database"
    "github.com/architecture-template/echo-ddd/infra/dao"
    "github.com/architecture-template/echo-ddd/api/service"	
    "github.com/architecture-template/echo-ddd/api/presentation/controller"
	"github.com/architecture-template/echo-ddd/api/presentation/middleware"
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
