//go:generate mockgen -source=./resource_service.go -destination=./resource_service_mock.gen.go -package=resource
package resource

import (
	"github.com/game-core/gocrafter/pkg/domain/model/resource/masterResource"
)

type ResourceService interface {
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
