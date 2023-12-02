package example

import (
	"fmt"
	exampleService "github.com/game-core/gocrafter/domain/service/api/example"
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
	results, err := e.exampleService.ListExample(10)
	if err != nil {
		return err
	}

	fmt.Println("List Examples:")
	for _, example := range *results.Examples {
		fmt.Printf("ID: %d, Name: %s, Detail: %s, Count: %d\n",
			example.ID,
			example.Name,
			*example.Detail,
			example.Count,
		)
	}

	return nil
}
