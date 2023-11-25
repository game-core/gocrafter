//go:generate mockgen -source=./item_service.go -destination=./item_service_mock.gen.go -package=item
package item

import (
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

	i, err := s.itemRepository.FindByName(req.ItemName)
	if err != nil {
		return nil, err
	}

	ib, err := s.itemBoxRepository.FindOrNilByAccountIDAndItemName(req.AccountID, req.ItemName, req.ShardKey)
	if err != nil {
		return nil, err
	}

	if ib == nil {
		ib = &userItemEntity.ItemBox{
			ShardKey:  req.ShardKey,
			AccountID: req.AccountID,
			ItemName:  req.ItemName,
			Count:     req.Count,
		}
	} else {
		ib.Count = ib.Count + req.Count
	}

	if _, err := s.itemBoxRepository.Save(ib, req.ShardKey, tx); err != nil {
		return nil, err
	}

	return &response.ReceiveItemInBox{
		Status: 200,
		Item: response.Item{
			ID:     i.ID,
			Name:   i.Name,
			Detail: i.Detail,
		},
	}, nil
}
