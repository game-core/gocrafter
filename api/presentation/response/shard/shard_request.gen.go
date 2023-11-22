package shard

type Shards []Shard

type Shard struct {
	ID int64 `json:"id"`

	ShardKey int `json:"shard_key"`

	Name string `json:"name"`

	Count int `json:"count"`
}
