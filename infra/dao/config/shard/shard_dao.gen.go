package shard

import (
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/config/shard"
	shardRepository "github.com/game-core/gocrafter/domain/repository/config/shard"
)

type shardDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewShardDao(conn *database.SqlHandler) shardRepository.ShardRepository {
	return &shardDao{
		Read:  conn.Config.ReadConn,
		Write: conn.Config.WriteConn,
	}
}

func (d *shardDao) Create(entity *shard.Shard, tx *gorm.DB) (*shard.Shard, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&shard.Shard{}).Create(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *shardDao) Delete(entity *shard.Shard, tx *gorm.DB) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&shard.Shard{}).Where("id = ?", entity.ID).Delete(entity)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (d *shardDao) FindByID(ID int64) (*shard.Shard, error) {
	entity := &shard.Shard{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *shardDao) FindByShardKey(ShardKey int) (*shard.Shard, error) {
	entity := &shard.Shard{}
	res := d.Read.Where("shard_key = ?", ShardKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *shardDao) FindOrNilByID(ID int64) (*shard.Shard, error) {
	entity := &shard.Shard{}
	res := d.Read.Where("id = ?", ID).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *shardDao) FindOrNilByShardKey(ShardKey int) (*shard.Shard, error) {
	entity := &shard.Shard{}
	res := d.Read.Where("shard_key = ?", ShardKey).Find(entity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *shardDao) List(limit int) (*shard.Shards, error) {
	entity := &shard.Shards{}
	res := d.Read.Limit(limit).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *shardDao) ListByShardKey(ShardKey int) (*shard.Shards, error) {
	entity := &shard.Shards{}
	res := d.Read.Where("shard_key = ?", ShardKey).Find(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (d *shardDao) Save(entity *shard.Shard, tx *gorm.DB) (*shard.Shard, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = d.Write
	}

	res := conn.Model(&shard.Shard{}).Where("id = ?", entity.ID).Save(entity)
	if err := res.Error; err != nil {
		return nil, err
	}

	return entity, nil
}
