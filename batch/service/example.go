package service

import (
	"github.com/game-core/gocrafter/domain/model"
	"github.com/game-core/gocrafter/domain/repository"
)

type ExampleService interface {
	ListExample() (*model.Examples, error)
}

type exampleService struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleService(exampleRepository repository.ExampleRepository) ExampleService {
	return &exampleService{
		exampleRepository: exampleRepository,
	}
}

func (e *exampleService) ListExample() (*model.Examples, error) {
	result, err := e.exampleRepository.List(10)
	if err != nil {
		return nil, err
	}

	return result, nil
}
