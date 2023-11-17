//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/batch/command"
	"github.com/game-core/gocrafter/batch/service"
	"github.com/game-core/gocrafter/config/db"
	"github.com/game-core/gocrafter/infra/dao"
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
