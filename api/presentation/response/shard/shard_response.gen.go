package shard

type Shards []Shard

type Shard struct {
	ID int64 `json:"id"`

	ShardKey string `json:"shard_key"`

	Name string `json:"name"`

	Count int `json:"count"`
}

func ToShard(ID int64, ShardKey string, Name string, Count int) *Shard {
	return &Shard{
		ID:       ID,
		ShardKey: ShardKey,
		Name:     Name,
		Count:    Count,
	}
}
