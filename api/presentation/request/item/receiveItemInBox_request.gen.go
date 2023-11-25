package item

type ReceiveItemInBoxs []ReceiveItemInBox

type ReceiveItemInBox struct {
	ShardKey string `json:"shard_key"`

	AccountID int64 `json:"account_id"`

	Items Items `json:"items"`
}
