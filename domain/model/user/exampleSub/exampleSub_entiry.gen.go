package exampleSub

import (
	"time"
)

type ExampleSubs []ExampleSub

type ExampleSub struct {
	ID int64 `json:"ID"`

	UserID int64 `json:"UserID"`

	ExampleID int64 `json:"ExampleID"`

	Name string `json:"Name"`

	Detail *string `json:"Detail"`

	Count int `json:"Count"`

	CreatedAt time.Time `json:"CreatedAt" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"UpdatedAt" gorm:"autoUpdateTime"`
}
