//go:generate mockgen -source=./example_service.go -destination=./example_service_mock.gen.go -package=example
package example

import (
	request "github.com/game-core/gocrafter/admin/presentation/request/example"
	response "github.com/game-core/gocrafter/admin/presentation/response/example"
	adminRepository "github.com/game-core/gocrafter/domain/repository/admin"
	exampleRepository "github.com/game-core/gocrafter/domain/repository/admin/example"
)

type ExampleService interface {
	GetExample(example *request.GetExample) (*response.GetExample, error)
}

type exampleService struct {
	transactionRepository adminRepository.TransactionRepository
	exampleRepository     exampleRepository.ExampleRepository
}

func NewExampleService(
	transactionRepository adminRepository.TransactionRepository,
	exampleRepository exampleRepository.ExampleRepository,
) ExampleService {
	return &exampleService{
		exampleRepository:     exampleRepository,
		transactionRepository: transactionRepository,
	}
}

// GetExample 取得する
func (e *exampleService) GetExample(example *request.GetExample) (*response.GetExample, error) {
	ex, err := e.exampleRepository.FindByID(example.ID)
	if err != nil {
		return nil, err
	}

	return response.ToGetExample(200, response.ToExample(ex.ID, ex.Name, ex.Detail, ex.Count)), nil
}
