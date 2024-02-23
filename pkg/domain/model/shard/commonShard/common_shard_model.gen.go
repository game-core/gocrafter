// Package commonShard シャード管理
package commonShard

type CommonShards []*CommonShard

type CommonShard struct {
	Id       int64
	ShardKey string
	Name     string
	Count    int32
}

func NewCommonShard() *CommonShard {
	return &CommonShard{}
}

func NewCommonShards() CommonShards {
	return CommonShards{}
}

func SetCommonShard(id int64, shardKey string, name string, count int32) *CommonShard {
	return &CommonShard{
		Id:       id,
		ShardKey: shardKey,
		Name:     name,
		Count:    count,
	}
}
