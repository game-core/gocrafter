package example

import (
	"time"
)

type Examples []Example

type Example struct {
	ID int64 `json:"ID"`

	UserID int64 `json:"UserID"`

	Name string `json:"Name"`

	Detail *string `json:"Detail"`

	Count int `json:"Count"`

	CreatedAt time.Time `json:"CreatedAt" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"UpdatedAt" gorm:"autoUpdateTime"`
}
