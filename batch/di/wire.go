//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/config/database"

	exampleCommand "github.com/game-core/gocrafter/batch/command/example"
	exampleService "github.com/game-core/gocrafter/domain/service/api/example"
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
