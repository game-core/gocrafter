package item

type ReceiveItemInBoxs []ReceiveItemInBox

type ReceiveItemInBox struct {
	Status int64 `json:"status"`

	Items Items `json:"items"`
}
