//go:generate mockgen -source=./item_service.go -destination=./item_service_mock.gen.go -package=item
package item

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/item/masterItem"
	"github.com/game-core/gocrafter/pkg/domain/model/item/userItemBox"
)

type ItemService interface {
	Create(ctx context.Context, tx *gorm.DB, req *ItemCreateRequest) (*ItemCreateResponse, error)
	Receive(ctx context.Context, tx *gorm.DB, req *ItemReceiveRequest) (*ItemReceiveResponse, error)
}

type itemService struct {
	userItemBoxRepository userItemBox.UserItemBoxRepository
	masterItemRepository  masterItem.MasterItemRepository
}

func NewItemService(
	userItemBoxRepository userItemBox.UserItemBoxRepository,
	masterItemRepository masterItem.MasterItemRepository,
) ItemService {
	return &itemService{
		userItemBoxRepository: userItemBoxRepository,
		masterItemRepository:  masterItemRepository,
	}
}

// Create アイテムを作成する
func (s *itemService) Create(ctx context.Context, tx *gorm.DB, req *ItemCreateRequest) (*ItemCreateResponse, error) {
	masterItemModel, err := s.masterItemRepository.Find(ctx, req.MasterItemId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterItemRepository.Find", err)
	}

	userItemBoxModel, err := s.userItemBoxRepository.FindOrNil(ctx, req.UserId, req.MasterItemId)
	if err != nil {
		return nil, errors.NewMethodError("s.userItemBoxRepository.FindOrNil", err)
	}

	// 過去に取得したことがある場合
	if userItemBoxModel != nil {
		userItemBoxModel.Count = userItemBoxModel.Count + req.Count
		result, err := s.userItemBoxRepository.Update(ctx, tx, userItemBoxModel)
		if err != nil {
			return nil, errors.NewMethodError("s.userItemBoxRepository.Update", err)
		}

		return SetItemCreateResponse(result, masterItemModel), nil
	}

	// 新規に取得した場合
	result, err := s.userItemBoxRepository.Create(ctx, tx, userItemBox.SetUserItemBox(req.UserId, req.MasterItemId, req.Count))
	if err != nil {
		return nil, errors.NewMethodError("s.userItemBoxRepository.Create", err)
	}

	return SetItemCreateResponse(result, masterItemModel), nil
}

// Receive アイテムを受け取る
func (s *itemService) Receive(ctx context.Context, tx *gorm.DB, req *ItemReceiveRequest) (*ItemReceiveResponse, error) {
	var userItemBoxes userItemBox.UserItemBoxes
	var masterItems masterItem.MasterItems

	for _, item := range req.Items {
		itemCreateResponse, err := s.Create(ctx, tx, SetItemCreateRequest(req.UserId, item.MasterItemId, item.Count))
		if err != nil {
			return nil, errors.NewMethodError("s.Create", err)
		}

		userItemBoxes = append(userItemBoxes, itemCreateResponse.UserItemBox)
		masterItems = append(masterItems, itemCreateResponse.MasterItem)
	}

	return SetItemReceiveResponse(userItemBoxes, masterItems), nil
}
