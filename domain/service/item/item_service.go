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

	ib, err := s.itemBoxRepository.FindOrNilByAccountIDAndItemName(req.AccountID, req.ItemName, req.ShardKey)
	if err != nil {
		return nil, err
	}

	newIb := &userItemEntity.ItemBox{}
	if ib != nil {
		ib.Count = ib.Count + 1
		newIb, err = s.itemBoxRepository.Update(ib, req.ShardKey, tx)
		if err != nil {
			return nil, err
		}
	} else {
		newIb, err = s.itemBoxRepository.Create(
			&userItemEntity.ItemBox{
				ShardKey:  req.ShardKey,
				AccountID: req.AccountID,
				ItemName:  req.ItemName,
				Count:     1,
			}, req.ShardKey, tx,
		)
		if err != nil {
			return nil, err
		}
	}

	i, err := s.itemRepository.FindByName(newIb.ItemName)
	if err != nil {
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
