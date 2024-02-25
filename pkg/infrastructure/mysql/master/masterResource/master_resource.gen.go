// Package masterResource レアリティ
package masterResource

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterResources []*MasterResource

type MasterResource struct {
	Id           int64
	Name         string
	ResourceType enum.ResourceType
}

func NewMasterResource() *MasterResource {
	return &MasterResource{}
}

func NewMasterResources() MasterResources {
	return MasterResources{}
}

func SetMasterResource(id int64, name string, resourceType enum.ResourceType) *MasterResource {
	return &MasterResource{
		Id:           id,
		Name:         name,
		ResourceType: resourceType,
	}
}

func (t *MasterResource) TableName() string {
	return "master_resource"
}
