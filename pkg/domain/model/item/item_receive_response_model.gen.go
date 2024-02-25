// Package item アイテム受け取りレスポンス
package item

import (
	"github.com/game-core/gocrafter/pkg/domain/model/item/masterItem"
	"github.com/game-core/gocrafter/pkg/domain/model/item/userItemBox"
)

type ItemReceiveResponses []*ItemReceiveResponse

type ItemReceiveResponse struct {
	UserItemBoxes userItemBox.UserItemBoxes
	MasterItems   masterItem.MasterItems
}

func NewItemReceiveResponse() *ItemReceiveResponse {
	return &ItemReceiveResponse{}
}

func NewItemReceiveResponses() ItemReceiveResponses {
	return ItemReceiveResponses{}
}

func SetItemReceiveResponse(userItemBoxes userItemBox.UserItemBoxes, masterItems masterItem.MasterItems) *ItemReceiveResponse {
	return &ItemReceiveResponse{
		UserItemBoxes: userItemBoxes,
		MasterItems:   masterItems,
	}
}
