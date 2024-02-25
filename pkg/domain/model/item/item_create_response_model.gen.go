// Package item アイテム作成レスポンス
package item

import (
	"github.com/game-core/gocrafter/pkg/domain/model/item/masterItem"
	"github.com/game-core/gocrafter/pkg/domain/model/item/userItemBox"
)

type ItemCreateResponses []*ItemCreateResponse

type ItemCreateResponse struct {
	UserItemBox *userItemBox.UserItemBox
	MasterItem  *masterItem.MasterItem
}

func NewItemCreateResponse() *ItemCreateResponse {
	return &ItemCreateResponse{}
}

func NewItemCreateResponses() ItemCreateResponses {
	return ItemCreateResponses{}
}

func SetItemCreateResponse(userItemBox *userItemBox.UserItemBox, masterItem *masterItem.MasterItem) *ItemCreateResponse {
	return &ItemCreateResponse{
		UserItemBox: userItemBox,
		MasterItem:  masterItem,
	}
}
