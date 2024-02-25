// Package item アイテム受け取りリクエスト
package item

type ItemReceiveRequests []*ItemReceiveRequest

type ItemReceiveRequest struct {
	UserId string
	Items  Items
}

func NewItemReceiveRequest() *ItemReceiveRequest {
	return &ItemReceiveRequest{}
}

func NewItemReceiveRequests() ItemReceiveRequests {
	return ItemReceiveRequests{}
}

func SetItemReceiveRequest(userId string, items Items) *ItemReceiveRequest {
	return &ItemReceiveRequest{
		UserId: userId,
		Items:  items,
	}
}
