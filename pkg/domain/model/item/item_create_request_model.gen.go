// Package item アイテム作成リクエスト
package item

type ItemCreateRequests []*ItemCreateRequest

type ItemCreateRequest struct {
	UserId       string
	MasterItemId int64
	Count        int32
}

func NewItemCreateRequest() *ItemCreateRequest {
	return &ItemCreateRequest{}
}

func NewItemCreateRequests() ItemCreateRequests {
	return ItemCreateRequests{}
}

func SetItemCreateRequest(userId string, masterItemId int64, count int32) *ItemCreateRequest {
	return &ItemCreateRequest{
		UserId:       userId,
		MasterItemId: masterItemId,
		Count:        count,
	}
}
