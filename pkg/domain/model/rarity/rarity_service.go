//go:generate mockgen -source=./rarity_service.go -destination=./rarity_service_mock.gen.go -package=rarity
package rarity

import (
	"context"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/enum"
	"github.com/game-core/gocrafter/pkg/domain/model/rarity/masterRarity"
)

type RarityService interface {
	GetAll(cxt context.Context) (masterRarity.MasterRarities, error)
	GetByRarityType(cxt context.Context, rarityType enum.RarityType) (*masterRarity.MasterRarity, error)
}

type rarityService struct {
	masterRarityMysqlRepository masterRarity.MasterRarityMysqlRepository
}

func NewRarityService(
	masterRarityMysqlRepository masterRarity.MasterRarityMysqlRepository,
) RarityService {
	return &rarityService{
		masterRarityMysqlRepository: masterRarityMysqlRepository,
	}
}

// GetAll レアリティ一覧を取得する
func (s *rarityService) GetAll(cxt context.Context) (masterRarity.MasterRarities, error) {
	results, err := s.masterRarityMysqlRepository.FindList(cxt)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRarityMysqlRepository.FindList", err)
	}

	return results, nil
}

// GetByRarityType レアリティを取得する
func (s *rarityService) GetByRarityType(cxt context.Context, rarityType enum.RarityType) (*masterRarity.MasterRarity, error) {
	result, err := s.masterRarityMysqlRepository.FindByRarityType(cxt, rarityType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterRarityMysqlRepository.FindByRarityType", err)
	}

	return result, nil
}
