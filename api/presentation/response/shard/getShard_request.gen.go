package shard

type GetShards []GetShard

type GetShard struct {
	Status int64 `json:"status"`

	NextShardKey string `json:"next_shard_key"`

	Shards *Shards `json:"shards"`
}
