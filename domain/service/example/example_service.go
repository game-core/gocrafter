package example

import (
	exampleEntity "github.com/game-core/gocrafter/domain/entity/master/example"
	exampleRepository "github.com/game-core/gocrafter/domain/repository/master/example"
)

type ExampleService interface {
	ListExampleBatch() (*exampleEntity.Examples, error)
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

func (e *exampleService) ListExampleBatch() (*exampleEntity.Examples, error) {
	result, err := e.exampleRepository.List(10)
	if err != nil {
		return nil, err
	}

	return result, nil
}
