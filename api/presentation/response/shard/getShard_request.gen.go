package shard

type GetShards []GetShard

type GetShard struct {
	Status int64 `json:"status"`

	NextShardKey int `json:"shard_key"`

	Shards *Shards `json:"shards"`
}
