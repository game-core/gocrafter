//go:generate mockgen -source=./example_service.go -destination=./example_service_mock.gen.go -package=example
package example

import (
	response "github.com/game-core/gocrafter/api/presentation/response/example"
	exampleRepository "github.com/game-core/gocrafter/domain/repository/master/example"
)

type ExampleService interface {
	ListExample(limit int) (*response.ListExample, error)
}

type exampleService struct {
	exampleRepository exampleRepository.ExampleRepository
}

func NewExampleService(
	exampleRepository exampleRepository.ExampleRepository,
) ExampleService {
	return &exampleService{
		exampleRepository: exampleRepository,
	}
}

// ListExample 一覧を取得する
func (e *exampleService) ListExample(limit int) (*response.ListExample, error) {
	ers, err := e.exampleRepository.List(limit)
	if err != nil {
		return nil, err
	}

	return &response.ListExample{
		Status: 200,
		Items:  response.ToExamples(ers),
	}, nil
}
