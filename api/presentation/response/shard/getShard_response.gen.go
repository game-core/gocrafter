package shard

type GetShards []GetShard

type GetShard struct {
	Status int64 `json:"status"`

	NextShardKey string `json:"next_shard_key"`

	Shards *Shards `json:"shards"`
}

func ToGetShard(Status int64, NextShardKey string, Shards *Shards) *GetShard {
	return &GetShard{
		Status:       Status,
		NextShardKey: NextShardKey,
		Shards:       Shards,
	}
}
