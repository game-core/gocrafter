package shard

import (
	"time"
)

type Shards []Shard

type Shard struct {
	ID int64 `json:"id"`

	ShardKey string `json:"shard_key"`

	Name string `json:"name"`

	Count int `json:"count"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *Shard) TableName() string {
	return "shard"
}
