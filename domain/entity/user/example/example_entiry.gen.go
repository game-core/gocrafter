package example

import (
	"time"
)

type Examples []Example

type Example struct {
	ID int64 `json:"id"`

	UserID int64 `json:"user_id"`

	Name string `json:"name"`

	Detail *string `json:"detail"`

	Count int `json:"count"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
