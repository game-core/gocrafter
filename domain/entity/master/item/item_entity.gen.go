package item

import (
	"time"
)

type Items []Item

type Item struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Detail string `json:"detail"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
