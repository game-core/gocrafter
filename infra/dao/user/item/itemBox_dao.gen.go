package item

import (
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/user/item"
	itemRepository "github.com/game-core/gocrafter/domain/repository/user/item"
)

type itemBoxDao struct {
	ShardConn *database.ShardConn
}

func NewItemBoxDao(conn *database.SqlHandler) itemRepository.ItemBoxRepository {
	return &itemBoxDao{
		ShardConn: conn.User,
	}
}

func (d *itemBoxDao) Create(entity *item.ItemBox, shardKey int, tx *gorm.DB) (*item.ItemBox, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&item.ItemBox{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) Delete(entity *item.ItemBox, shardKey int, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&item.ItemBox{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *itemBoxDao) FindByAccountID(AccountID int64, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindByAccountIDAndItemName(AccountID int64, ItemName string, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Where("item_name = ?", ItemName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindByID(ID int64, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindByItemName(ItemName string, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("item_name = ?", ItemName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindOrNilByAccountID(AccountID int64, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindOrNilByAccountIDAndItemName(AccountID int64, ItemName string, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Where("item_name = ?", ItemName).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindOrNilByID(ID int64, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindOrNilByItemName(ItemName string, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("item_name = ?", ItemName).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) List(limit int, shardKey int) (*item.ItemBoxs, error) {
	entity := &item.ItemBoxs{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) ListByAccountID(AccountID int64, shardKey int) (*item.ItemBoxs, error) {
	entity := &item.ItemBoxs{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) ListByAccountIDAndItemName(AccountID int64, ItemName string, shardKey int) (*item.ItemBoxs, error) {
	entity := &item.ItemBoxs{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("account_id = ?", AccountID).Where("item_name = ?", ItemName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) ListByItemName(ItemName string, shardKey int) (*item.ItemBoxs, error) {
	entity := &item.ItemBoxs{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("item_name = ?", ItemName).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) Save(entity *item.ItemBox, shardKey int, tx *gorm.DB) (*item.ItemBox, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&item.ItemBox{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
