// Package commonRankingRoom ルームランキング
package commonRankingRoom

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/ranking/commonRankingRoom"
)

type commonRankingRoomDao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func NewCommonRankingRoomDao(conn *database.RedisHandler) commonRankingRoom.CommonRankingRoomRedisRepository {
	return &commonRankingRoomDao{
		ReadRedisConn:  conn.Common.ReadRedisConn,
		WriteRedisConn: conn.Common.WriteRedisConn,
	}
}

func (s *commonRankingRoomDao) Find(ctx context.Context, masterRankingId int64, roomId string, userId string) (*commonRankingRoom.CommonRankingRoom, error) {
	t := NewCommonRankingRoom()
	data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), fmt.Sprintf("%s:masterRankingId:%v,roomId:%v,userId:%v", t.TableName(), masterRankingId, roomId, userId)).Result()
	if err != nil {
		return nil, err
	}

	if err := t.JsonToTable(data); err != nil {
		return nil, err
	}

	return commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score), nil
}

func (s *commonRankingRoomDao) FindOrNil(ctx context.Context, masterRankingId int64, roomId string, userId string) (*commonRankingRoom.CommonRankingRoom, error) {
	t := NewCommonRankingRoom()
	data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), fmt.Sprintf("%s:masterRankingId:%v,roomId:%v,userId:%v", t.TableName(), masterRankingId, roomId, userId)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	if err := t.JsonToTable(data); err != nil {
		return nil, err
	}

	return commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score), nil
}

func (s *commonRankingRoomDao) Set(ctx context.Context, tx redis.Pipeliner, m *commonRankingRoom.CommonRankingRoom) (*commonRankingRoom.CommonRankingRoom, error) {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := &CommonRankingRoom{
		MasterRankingId: m.MasterRankingId,
		RoomId:          m.RoomId,
		UserId:          m.UserId,
		Score:           m.Score,
	}

	jt, err := t.TableToJson()
	if err != nil {
		return nil, err
	}

	if err := conn.HSet(ctx, t.TableName(), fmt.Sprintf("%s:masterRankingId:%v,roomId:%v,userId:%v", t.TableName(), m.MasterRankingId, m.RoomId, m.UserId), jt).Err(); err != nil {
		return nil, err
	}

	return commonRankingRoom.SetCommonRankingRoom(t.MasterRankingId, t.RoomId, t.UserId, t.Score), nil
}

func (s *commonRankingRoomDao) Delete(ctx context.Context, tx redis.Pipeliner, m *commonRankingRoom.CommonRankingRoom) error {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := NewCommonRankingRoom()
	if err := conn.HDel(ctx, t.TableName(), fmt.Sprintf("%s:masterRankingId:%v,roomId:%v,userId:%v", t.TableName(), m.MasterRankingId, m.RoomId, m.UserId)).Err(); err != nil {
		return err
	}

	return nil
}
