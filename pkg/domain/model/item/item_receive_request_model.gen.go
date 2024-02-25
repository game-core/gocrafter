// Package item アイテム受け取りリクエスト
package item

type ItemReceiveRequests []*ItemReceiveRequest

type ItemReceiveRequest struct {
	UserId        string
	MasterItemIds int64
}

func NewItemReceiveRequest() *ItemReceiveRequest {
	return &ItemReceiveRequest{}
}

func NewItemReceiveRequests() ItemReceiveRequests {
	return ItemReceiveRequests{}
}

func SetItemReceiveRequest(userId string, masterItemIds int64) *ItemReceiveRequest {
	return &ItemReceiveRequest{
		UserId:        userId,
		MasterItemIds: masterItemIds,
	}
}
