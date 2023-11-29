package item

type ReceiveItemInBoxs []ReceiveItemInBox

type ReceiveItemInBox struct {
	Status int64 `json:"status"`

	Items Items `json:"items"`
}

func ToReceiveItemInBox(Status int64, Items Items) *ReceiveItemInBox {
	return &ReceiveItemInBox{
		Status: Status,
		Items:  Items,
	}
}
