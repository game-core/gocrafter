// Package userItemBox ユーザーアイテムボックス
package userItemBox

import "time"

type UserItemBoxes []*UserItemBox

type UserItemBox struct {
	UserId       string
	MasterItemId int64
	Count        int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUserItemBox() *UserItemBox {
	return &UserItemBox{}
}

func NewUserItemBoxes() UserItemBoxes {
	return UserItemBoxes{}
}

func SetUserItemBox(userId string, masterItemId int64, count int32, createdAt time.Time, updatedAt time.Time) *UserItemBox {
	return &UserItemBox{
		UserId:       userId,
		MasterItemId: masterItemId,
		Count:        count,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func (t *UserItemBox) TableName() string {
	return "user_item_box"
}
