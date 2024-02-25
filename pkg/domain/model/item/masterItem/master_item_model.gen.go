// Package masterItem アイテム
package masterItem

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterItems []*MasterItem

type MasterItem struct {
	Id           int64
	Name         string
	ResourceType enum.ResourceType
	RarityType   enum.RarityType
	Content      string
}

func NewMasterItem() *MasterItem {
	return &MasterItem{}
}

func NewMasterItems() MasterItems {
	return MasterItems{}
}

func SetMasterItem(id int64, name string, resourceType enum.ResourceType, rarityType enum.RarityType, content string) *MasterItem {
	return &MasterItem{
		Id:           id,
		Name:         name,
		ResourceType: resourceType,
		RarityType:   rarityType,
		Content:      content,
	}
}
