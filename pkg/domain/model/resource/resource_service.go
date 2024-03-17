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
	masterResourceRepository masterResource.MasterResourceRepository
}

func NewResourceService(
	masterResourceRepository masterResource.MasterResourceRepository,
) ResourceService {
	return &resourceService{
		masterResourceRepository: masterResourceRepository,
	}
}

// GetAll リソース一覧を取得する
func (s *resourceService) GetAll(cxt context.Context) (masterResource.MasterResources, error) {
	results, err := s.masterResourceRepository.FindList(cxt)
	if err != nil {
		return nil, errors.NewMethodError("s.masterResourceRepository.FindList", err)
	}

	return results, nil
}

// GetByResourceType リソースを取得する
func (s *resourceService) GetByResourceType(cxt context.Context, resourceType enum.ResourceType) (*masterResource.MasterResource, error) {
	result, err := s.masterResourceRepository.FindByResourceType(cxt, resourceType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterResourceRepository.FindByResourceType", err)
	}

	return result, nil
}
