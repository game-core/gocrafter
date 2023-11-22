package shard

import (
	"time"
)

type Shards []Shard

type Shard struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Count int `json:"count"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
