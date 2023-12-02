//go:build wireinject
// +build wireinject

package di

import (
	exampleService "github.com/game-core/gocrafter/domain/service/api/example"
	"github.com/google/wire"

	exampleCommand "github.com/game-core/gocrafter/batch/command/example"
	"github.com/game-core/gocrafter/config/database"
	exampleDao "github.com/game-core/gocrafter/infra/dao/master/example"
)

func InitializeExampleCommand() exampleCommand.ExampleCommand {
	wire.Build(
		database.NewDB,
		exampleDao.NewExampleDao,
		exampleService.NewExampleService,
		exampleCommand.NewExampleCommand,
	)

	return nil
}
