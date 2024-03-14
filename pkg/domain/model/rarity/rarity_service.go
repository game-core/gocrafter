//go:generate mockgen -source=./rarity_service.go -destination=./rarity_service_mock.gen.go -package=rarity
package rarity

import (
	"github.com/game-core/gocrafter/pkg/domain/model/rarity/masterRarity"
)

type RarityService interface {
}

type rarityService struct {
	masterRarityRepository masterRarity.MasterRarityRepository
}

func NewRarityService(
	masterRarityRepository masterRarity.MasterRarityRepository,
) RarityService {
	return &rarityService{
		masterRarityRepository: masterRarityRepository,
	}
}
