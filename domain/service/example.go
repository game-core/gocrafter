package service

import (
	"github.com/game-core/gocrafter/api/presentation/parameter"
	"github.com/game-core/gocrafter/domain/model"
	"github.com/game-core/gocrafter/domain/repository"
)

type ExampleService interface {
	FindByExampleKey(exampleKey *parameter.ExampleKey) (*model.Example, error)
}

type exampleService struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleService(
	exampleRepository repository.ExampleRepository,
	) ExampleService {
	return &exampleService{
		exampleRepository: exampleRepository,
	}
}

// FindByKey キーから取得する
func (e *exampleService) FindByExampleKey(exampleKey *parameter.ExampleKey) (*model.Example, error) {
	result, err := e.exampleRepository.FindByExampleKey(exampleKey.ExampleKey)
	if err != nil {
		return nil, err
	}

	return result, nil
}
