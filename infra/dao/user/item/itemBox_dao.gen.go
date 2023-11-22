package item

import (
	"github.com/jinzhu/gorm"

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

func (d *itemBoxDao) FindByID(ID int64, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindByItemID(ItemID int64, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("item_id = ?", ItemID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindByUserID(UserID int64, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("item_id = ?", UserID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) FindByUserIDAndItemID(UserID int64, ItemID int64, shardKey int) (*item.ItemBox, error) {
	entity := &item.ItemBox{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("item_id = ?", UserID).Where("item_id = ?", ItemID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) List(limit int64, shardKey int) (*item.ItemBoxs, error) {
	entity := &item.ItemBoxs{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) ListByItemID(ItemID int64, shardKey int) (*item.ItemBoxs, error) {
	entity := &item.ItemBoxs{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("item_id = ?", ItemID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) ListByUserID(UserID int64, shardKey int) (*item.ItemBoxs, error) {
	entity := &item.ItemBoxs{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("item_id = ?", UserID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) ListByUserIDAndItemID(UserID int64, ItemID int64, shardKey int) (*item.ItemBoxs, error) {
	entity := &item.ItemBoxs{}
	res := d.ShardConn.Shards[shardKey].ReadConn.Where("item_id = ?", UserID).Where("item_id = ?", ItemID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *itemBoxDao) Update(entity *item.ItemBox, shardKey int, tx *gorm.DB) (*item.ItemBox, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.ShardConn.Shards[shardKey].WriteConn
	}

	res := conn.Model(&item.ItemBox{}).Where("id = ?", entity.ID).Update(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
