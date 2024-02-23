// Package commonShard シャード管理
package commonShard

import (
	"time"
)

type CommonShards []*CommonShard

type CommonShard struct {
	Id        int64
	ShardKey  string
	Name      string
	Count     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCommonShard() *CommonShard {
	return &CommonShard{}
}

func NewCommonShards() CommonShards {
	return CommonShards{}
}

func SetCommonShard(id int64, shardKey string, name string, count int32, createdAt time.Time, updatedAt time.Time) *CommonShard {
	return &CommonShard{
		Id:        id,
		ShardKey:  shardKey,
		Name:      name,
		Count:     count,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (t *CommonShard) TableName() string {
	return "common_shard"
}
