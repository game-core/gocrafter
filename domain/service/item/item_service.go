//go:generate mockgen -source=./item_service.go -destination=./item_service_mock.gen.go -package=item
package item

import (
	"gorm.io/gorm"
	"log"

	request "github.com/game-core/gocrafter/api/presentation/request/item"
	response "github.com/game-core/gocrafter/api/presentation/response/item"
	masterItemEntity "github.com/game-core/gocrafter/domain/entity/master/item"
	userItemEntity "github.com/game-core/gocrafter/domain/entity/user/item"
	masterItemRepository "github.com/game-core/gocrafter/domain/repository/master/item"
	userRepository "github.com/game-core/gocrafter/domain/repository/user"
	userItemRepository "github.com/game-core/gocrafter/domain/repository/user/item"
)

type ItemService interface {
	GetItemToEntity(name string) (*masterItemEntity.Item, error)
	ReceiveItemInBox(req *request.ReceiveItemInBox) (*response.ReceiveItemInBox, error)
}

type itemService struct {
	transactionRepository userRepository.TransactionRepository
	itemRepository        masterItemRepository.ItemRepository
	itemBoxRepository     userItemRepository.ItemBoxRepository
}

func NewItemService(
	transactionRepository userRepository.TransactionRepository,
	itemRepository masterItemRepository.ItemRepository,
	itemBoxRepository userItemRepository.ItemBoxRepository,
) ItemService {
	return &itemService{
		transactionRepository: transactionRepository,
		itemBoxRepository:     itemBoxRepository,
		itemRepository:        itemRepository,
	}
}

// GetItemToEntity アイテムをEntityで取得する
func (s *itemService) GetItemToEntity(name string) (*masterItemEntity.Item, error) {
	e, err := s.itemRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// ReceiveItemInBox アイテムをボックスに受け取る
func (s *itemService) ReceiveItemInBox(req *request.ReceiveItemInBox) (*response.ReceiveItemInBox, error) {
	// transaction
	tx, err := s.transactionRepository.Begin(req.ShardKey)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := s.transactionRepository.Rollback(tx); err != nil {
				log.Panicln(err)
			}
		} else {
			if err := s.transactionRepository.Commit(tx); err != nil {
				log.Panicln(err)
			}
		}
	}()

	items, err := s.receiveItemInBox(&req.Items, req.AccountID, req.ShardKey, tx)
	if err != nil {
		return nil, err
	}

	return response.ToReceiveItemInBox(200, *items), nil
}

// receiveItemBox
func (s *itemService) receiveItemInBox(items *request.Items, accountID int64, shardKey string, tx *gorm.DB) (*response.Items, error) {
	var itemEntities response.Items
	for _, item := range *items {
		i, err := s.itemRepository.FindByName(item.Name)
		if err != nil {
			return nil, err
		}

		ib, err := s.itemBoxRepository.FindOrNilByAccountIDAndItemName(accountID, i.Name, shardKey)
		if err != nil {
			return nil, err
		}

		if ib == nil {
			ib = &userItemEntity.ItemBox{
				ShardKey:  shardKey,
				AccountID: accountID,
				ItemName:  item.Name,
				Count:     item.Count,
			}
		} else {
			ib.Count = ib.Count + item.Count
		}

		if _, err := s.itemBoxRepository.Save(ib, shardKey, tx); err != nil {
			return nil, err
		}

		itemEntities = append(itemEntities, response.Item{
			ID:     i.ID,
			Name:   i.Name,
			Detail: i.Detail,
			Count:  item.Count,
		})
	}

	return &itemEntities, nil
}
