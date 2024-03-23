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
	userItemBoxMysqlRepository userItemBox.UserItemBoxMysqlRepository
	masterItemMysqlRepository  masterItem.MasterItemMysqlRepository
}

func NewItemService(
	userItemBoxMysqlRepository userItemBox.UserItemBoxMysqlRepository,
	masterItemMysqlRepository masterItem.MasterItemMysqlRepository,
) ItemService {
	return &itemService{
		userItemBoxMysqlRepository: userItemBoxMysqlRepository,
		masterItemMysqlRepository:  masterItemMysqlRepository,
	}
}

// Create アイテムを作成する
func (s *itemService) Create(ctx context.Context, tx *gorm.DB, req *ItemCreateRequest) (*ItemCreateResponse, error) {
	masterItemModel, err := s.masterItemMysqlRepository.Find(ctx, req.MasterItemId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterItemMysqlRepository.Find", err)
	}

	userItemBoxModel, err := s.userItemBoxMysqlRepository.FindOrNil(ctx, req.UserId, req.MasterItemId)
	if err != nil {
		return nil, errors.NewMethodError("s.userItemBoxMysqlRepository.FindOrNil", err)
	}

	// 過去に取得したことがある場合
	if userItemBoxModel != nil {
		userItemBoxModel.Count = userItemBoxModel.Count + req.Count
		result, err := s.userItemBoxMysqlRepository.Update(ctx, tx, userItemBoxModel)
		if err != nil {
			return nil, errors.NewMethodError("s.userItemBoxMysqlRepository.Update", err)
		}

		return SetItemCreateResponse(result, masterItemModel), nil
	}

	// 新規に取得した場合
	result, err := s.userItemBoxMysqlRepository.Create(ctx, tx, userItemBox.SetUserItemBox(req.UserId, req.MasterItemId, req.Count))
	if err != nil {
		return nil, errors.NewMethodError("s.userItemBoxMysqlRepository.Create", err)
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
