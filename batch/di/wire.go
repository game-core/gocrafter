// +build wireinject

package di

import (
    "github.com/google/wire"

    "github.com/architecture-template/echo-ddd/config/db"
    "github.com/architecture-template/echo-ddd/infra/dao"
    "github.com/architecture-template/echo-ddd/batch/service"	
    "github.com/architecture-template/echo-ddd/batch/command"
)

// example
func InitializeExampleCommand() command.ExampleCommand {
	wire.Build(
		db.NewDB,
		dao.NewExampleDao,
		service.NewExampleService,
		command.NewExampleCommand,
	)

    return nil
}
