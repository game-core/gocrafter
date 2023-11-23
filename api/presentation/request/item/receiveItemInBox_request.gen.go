package item

type ReceiveItemInBoxs []ReceiveItemInBox

type ReceiveItemInBox struct {
	ShardKey int `json:"shard_key"`

	AccountID int64 `json:"account_id"`

	ItemName string `json:"item_name"`
}
