// Package masterResource レアリティ
//
//go:generate mockgen -source=./master_resource_repository.gen.go -destination=./master_resource_repository_mock.gen.go -package=masterResource
package masterResource

import (
	context "context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterResourceRepository interface {
	Find(ctx context.Context, id int64) (*MasterResource, error)
	FindOrNil(ctx context.Context, id int64) (*MasterResource, error)
	FindByResourceType(ctx context.Context, resourceType enum.ResourceType) (*MasterResource, error)
	FinOrNilByResourceType(ctx context.Context, resourceType enum.ResourceType) (*MasterResource, error)
	FindList(ctx context.Context) (MasterResources, error)
	FindListByResourceType(ctx context.Context, resourceType enum.ResourceType) (MasterResources, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterResource) (*MasterResource, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterResources) (MasterResources, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterResource) (*MasterResource, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterResource) error
}
