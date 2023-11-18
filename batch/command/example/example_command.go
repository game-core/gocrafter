package example

import (
	"fmt"
	"time"

	exampleService "github.com/game-core/gocrafter/domain/service/example"
)

type ExampleCommand interface {
	ListExample() (err error)
}

type exampleCommand struct {
	exampleService exampleService.ExampleService
}

func NewExampleCommand(
	exampleService exampleService.ExampleService,
) ExampleCommand {
	return &exampleCommand{
		exampleService: exampleService,
	}
}

// ListExample exampleテーブル一覧を取得する
func (e *exampleCommand) ListExample() (err error) {
	results, err := e.exampleService.ListExampleBatch()
	if err != nil {
		return err
	}

	fmt.Println("List Examples:")
	for _, example := range *results {
		fmt.Printf("ID: %d, Name: %s, Detail: %s, Count: %d, CreatedAt: %s, UpdatedAt: %s\n",
			example.ID,
			example.Name,
			*example.Detail,
			example.Count,
			example.CreatedAt.Format(time.RFC3339),
			example.UpdatedAt.Format(time.RFC3339),
		)
	}

	return nil
}
