// Package commonRankingWorld ワールドランキング
package commonRankingWorld

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingWorld"
)

type commonRankingWorldDao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func NewCommonRankingWorldDao(conn *database.RedisHandler) commonRankingWorld.CommonRankingWorldRedisRepository {
	return &commonRankingWorldDao{
		ReadRedisConn:  conn.Common.ReadRedisConn,
		WriteRedisConn: conn.Common.WriteRedisConn,
	}
}

func (s *commonRankingWorldDao) Find(ctx context.Context, masterRankingId int64, userId string) (*commonRankingWorld.CommonRankingWorld, error) {
	t := NewCommonRankingWorld()
	data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), fmt.Sprintf("%s:masterRankingId:%v,userId:%v", t.TableName(), masterRankingId, userId)).Result()
	if err != nil {
		return nil, err
	}

	if err := t.JsonToTable(data); err != nil {
		return nil, err
	}

	return commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score, t.RankedAt), nil
}

func (s *commonRankingWorldDao) FindOrNil(ctx context.Context, masterRankingId int64, userId string) (*commonRankingWorld.CommonRankingWorld, error) {
	t := NewCommonRankingWorld()
	data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), fmt.Sprintf("%s:masterRankingId:%v,userId:%v", t.TableName(), masterRankingId, userId)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	if err := t.JsonToTable(data); err != nil {
		return nil, err
	}

	return commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score, t.RankedAt), nil
}

func (s *commonRankingWorldDao) Set(ctx context.Context, tx redis.Pipeliner, m *commonRankingWorld.CommonRankingWorld) (*commonRankingWorld.CommonRankingWorld, error) {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := &CommonRankingWorld{
		MasterRankingId: m.MasterRankingId,
		UserId:          m.UserId,
		Score:           m.Score,
		RankedAt:        m.RankedAt,
	}

	jt, err := t.TableToJson()
	if err != nil {
		return nil, err
	}

	if err := conn.HSet(ctx, t.TableName(), fmt.Sprintf("%s:masterRankingId:%v,userId:%v", t.TableName(), m.MasterRankingId, m.UserId), jt).Err(); err != nil {
		return nil, err
	}

	return commonRankingWorld.SetCommonRankingWorld(t.MasterRankingId, t.UserId, t.Score, t.RankedAt), nil
}

func (s *commonRankingWorldDao) Delete(ctx context.Context, tx redis.Pipeliner, m *commonRankingWorld.CommonRankingWorld) error {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := NewCommonRankingWorld()
	if err := conn.HDel(ctx, t.TableName(), fmt.Sprintf("%s:masterRankingId:%v,userId:%v", t.TableName(), m.MasterRankingId, m.UserId)).Err(); err != nil {
		return err
	}

	return nil
}
