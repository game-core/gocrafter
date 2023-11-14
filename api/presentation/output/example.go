package output

import (
	"github.com/game-core/gocrafter/domain/model"
)

type Example struct {
	ExampleKey  string `json:"example_key"`
	ExampleName string `json:"example_name"`
	Message     string `json:"message"`
}

func ToExample(u *model.Example) *Example {
	return &Example{
		ExampleKey:  u.ExampleKey,
		ExampleName: u.ExampleName,
		Message:     "get example completed",
	}
}
