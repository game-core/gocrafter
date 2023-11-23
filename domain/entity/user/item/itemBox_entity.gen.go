package item

import (
	"time"
)

type ItemBoxs []ItemBox

type ItemBox struct {
	ID int64 `json:"id"`

	ShardKey int `json:"shard_key"`

	AccountID int64 `json:"account_id"`

	ItemName string `json:"item_name"`

	Count int `json:"count"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (e *ItemBox) TableName() string {
	return "item_box"
}
