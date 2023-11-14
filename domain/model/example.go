package model

import (
	"time"

	"gorm.io/gorm"
)

type Examples []Example

type Example struct {
	ID          int64          `json:"id"`
	ExampleKey  string         `json:"example_key"`
	ExampleName string         `json:"example_name"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyExample() *Example {
	return &Example{}
}

func (e *Example) IsEmpty() bool {
	return (e.ID == 0 &&
		e.ExampleKey == "" &&
		e.ExampleName == "")
}
