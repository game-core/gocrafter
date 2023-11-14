package command

import (
	"fmt"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/architecture-template/echo-ddd/batch/service"
)

type ExampleCommand interface {
	ListExample() (err error)
}

type exampleCommand struct {
	exampleService service.ExampleService
}

func NewExampleCommand(exampleService service.ExampleService) ExampleCommand {
    return &exampleCommand{
        exampleService: exampleService,
    }
}

// ListExample exampleテーブル一覧を取得する
func (e *exampleCommand) ListExample() (err error) {
	results, err := e.exampleService.ListExample()
	if err != nil {
		return err
	}

	fmt.Println("List Examples:")
	for _, example := range *results {
		fmt.Printf("ID: %d, ExampleKey: %s, ExampleName: %s, CreatedAt: %s, UpdatedAt: %s\n",
			example.ID,
			example.ExampleKey,
			example.ExampleName,
			example.CreatedAt.Format(time.RFC3339),
			example.UpdatedAt.Format(time.RFC3339),
		)
	}

	return nil
}
