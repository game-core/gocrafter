// Package commonShard シャード管理
package commonShard

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/shard/commonShard"
)

type commonShardDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewCommonShardDao(conn *database.MysqlHandler) commonShard.CommonShardMysqlRepository {
	return &commonShardDao{
		ReadMysqlConn:  conn.Common.ReadMysqlConn,
		WriteMysqlConn: conn.Common.WriteMysqlConn,
	}
}

func (s *commonShardDao) Find(ctx context.Context, id int64) (*commonShard.CommonShard, error) {
	t := NewCommonShard()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonShard.SetCommonShard(t.Id, t.ShardKey, t.Name, t.Count), nil
}

func (s *commonShardDao) FindOrNil(ctx context.Context, id int64) (*commonShard.CommonShard, error) {
	t := NewCommonShard()
	res := s.ReadMysqlConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonShard.SetCommonShard(t.Id, t.ShardKey, t.Name, t.Count), nil
}

func (s *commonShardDao) FindByShardKey(ctx context.Context, shardKey string) (*commonShard.CommonShard, error) {
	t := NewCommonShard()
	res := s.ReadMysqlConn.WithContext(ctx).Where("shard_key = ?", shardKey).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonShard.SetCommonShard(t.Id, t.ShardKey, t.Name, t.Count), nil
}

func (s *commonShardDao) FindOrNilByShardKey(ctx context.Context, shardKey string) (*commonShard.CommonShard, error) {
	t := NewCommonShard()
	res := s.ReadMysqlConn.WithContext(ctx).Where("shard_key = ?", shardKey).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonShard.SetCommonShard(t.Id, t.ShardKey, t.Name, t.Count), nil
}

func (s *commonShardDao) FindList(ctx context.Context) (commonShard.CommonShards, error) {
	ts := NewCommonShards()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonShard.NewCommonShards()
	for _, t := range ts {
		ms = append(ms, commonShard.SetCommonShard(t.Id, t.ShardKey, t.Name, t.Count))
	}

	return ms, nil
}

func (s *commonShardDao) FindListByShardKey(ctx context.Context, shardKey string) (commonShard.CommonShards, error) {
	ts := NewCommonShards()
	res := s.ReadMysqlConn.WithContext(ctx).Where("shard_key = ?", shardKey).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonShard.NewCommonShards()
	for _, t := range ts {
		ms = append(ms, commonShard.SetCommonShard(t.Id, t.ShardKey, t.Name, t.Count))
	}

	return ms, nil
}

func (s *commonShardDao) Create(ctx context.Context, tx *gorm.DB, m *commonShard.CommonShard) (*commonShard.CommonShard, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonShard{
		Id:       m.Id,
		ShardKey: m.ShardKey,
		Name:     m.Name,
		Count:    m.Count,
	}
	res := conn.Model(NewCommonShard()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonShard.SetCommonShard(t.Id, t.ShardKey, t.Name, t.Count), nil
}

func (s *commonShardDao) CreateList(ctx context.Context, tx *gorm.DB, ms commonShard.CommonShards) (commonShard.CommonShards, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewCommonShards()
	for _, m := range ms {
		t := &CommonShard{
			Id:       m.Id,
			ShardKey: m.ShardKey,
			Name:     m.Name,
			Count:    m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewCommonShard()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *commonShardDao) Update(ctx context.Context, tx *gorm.DB, m *commonShard.CommonShard) (*commonShard.CommonShard, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonShard{
		Id:       m.Id,
		ShardKey: m.ShardKey,
		Name:     m.Name,
		Count:    m.Count,
	}
	res := conn.Model(NewCommonShard()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonShard.SetCommonShard(t.Id, t.ShardKey, t.Name, t.Count), nil
}

func (s *commonShardDao) Delete(ctx context.Context, tx *gorm.DB, m *commonShard.CommonShard) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewCommonShard()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewCommonShard())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
