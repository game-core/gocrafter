//go:generate mockgen -source=./resource_service.go -destination=./resource_service_mock.gen.go -package=resource
package resource

import (
	"context"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/resource/masterResource"
)

type ResourceService interface {
	GetAll(cxt context.Context) (masterResource.MasterResources, error)
	GetByResourceType(cxt context.Context, resourceType enum.ResourceType) (*masterResource.MasterResource, error)
}

type resourceService struct {
	masterResourceMysqlRepository masterResource.MasterResourceMysqlRepository
}

func NewResourceService(
	masterResourceMysqlRepository masterResource.MasterResourceMysqlRepository,
) ResourceService {
	return &resourceService{
		masterResourceMysqlRepository: masterResourceMysqlRepository,
	}
}

// GetAll リソース一覧を取得する
func (s *resourceService) GetAll(cxt context.Context) (masterResource.MasterResources, error) {
	results, err := s.masterResourceMysqlRepository.FindList(cxt)
	if err != nil {
		return nil, errors.NewMethodError("s.masterResourceMysqlRepository.FindList", err)
	}

	return results, nil
}

// GetByResourceType リソースを取得する
func (s *resourceService) GetByResourceType(cxt context.Context, resourceType enum.ResourceType) (*masterResource.MasterResource, error) {
	result, err := s.masterResourceMysqlRepository.FindByResourceType(cxt, resourceType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterResourceMysqlRepository.FindByResourceType", err)
	}

	return result, nil
}
