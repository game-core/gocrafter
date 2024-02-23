// Package commonShard シャード管理
package commonShard

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/shard/commonShard"
)

type commonShardDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
}

func NewCommonShardDao(conn *database.SqlHandler) commonShard.CommonShardRepository {
	return &commonShardDao{
		ReadConn:  conn.Common.ReadConn,
		WriteConn: conn.Common.WriteConn,
	}
}

func (s *commonShardDao) FindList(ctx context.Context) (commonShard.CommonShards, error) {
	ts := NewCommonShards()
	res := s.ReadConn.WithContext(ctx).Find(ts)
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
		conn = s.WriteConn
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
		conn = s.WriteConn
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
