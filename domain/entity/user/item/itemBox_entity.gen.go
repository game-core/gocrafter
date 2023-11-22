package item

import (
	"time"
)

type ItemBoxs []ItemBox

type ItemBox struct {
	ID int64 `json:"id"`

	ShardKey int `json:"shard_key"`

	UserID int64 `json:"item_id"`

	ItemID int64 `json:"item_id"`

	Count int `json:"count"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
