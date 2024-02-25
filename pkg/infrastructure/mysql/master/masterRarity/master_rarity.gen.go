// Package masterRarity レアリティ
package masterRarity

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterRarities []*MasterRarity

type MasterRarity struct {
	Id         int64
	Name       string
	RarityType enum.RarityType
}

func NewMasterRarity() *MasterRarity {
	return &MasterRarity{}
}

func NewMasterRarities() MasterRarities {
	return MasterRarities{}
}

func SetMasterRarity(id int64, name string, rarityType enum.RarityType) *MasterRarity {
	return &MasterRarity{
		Id:         id,
		Name:       name,
		RarityType: rarityType,
	}
}

func (t *MasterRarity) TableName() string {
	return "master_rarity"
}
