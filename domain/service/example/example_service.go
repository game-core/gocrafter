package example

import (
	response "github.com/game-core/gocrafter/api/presentation/response/example"
	exampleRepository "github.com/game-core/gocrafter/domain/repository/master/example"
)

type ExampleService interface {
	ListExample(limit int64) (*response.ListExample, error)
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
func (e *exampleService) ListExample(limit int64) (*response.ListExample, error) {
	ers, err := e.exampleRepository.List(limit)
	if err != nil {
		return nil, err
	}

	examples := make(response.Examples, len(*ers))
	for i, er := range *ers {
		example := &response.Example{
			ID:     er.ID,
			Name:   er.Name,
			Detail: er.Detail,
			Count:  er.Count,
		}
		examples[i] = *example
	}

	return &response.ListExample{
		Status: 200,
		Items:   &examples,
	}, nil
}
