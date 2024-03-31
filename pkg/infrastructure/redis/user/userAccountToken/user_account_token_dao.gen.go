// Package userAccountToken ユーザーアカウントトークン
package userAccountToken

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccountToken"
)

type userAccountTokenDao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func NewUserAccountTokenDao(conn *database.RedisHandler) userAccountToken.UserAccountTokenRedisRepository {
	return &userAccountTokenDao{
		ReadRedisConn:  conn.User.ReadRedisConn,
		WriteRedisConn: conn.User.WriteRedisConn,
	}
}

func (s *userAccountTokenDao) Find(ctx context.Context, userId string) (*userAccountToken.UserAccountToken, error) {
	t := NewUserAccountToken()
	data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), userId)).Result()
	if err != nil {
		return nil, err
	}

	if err := t.JsonToTable(data); err != nil {
		return nil, err
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenDao) FindOrNil(ctx context.Context, userId string) (*userAccountToken.UserAccountToken, error) {
	t := NewUserAccountToken()
	data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), userId)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	if err := t.JsonToTable(data); err != nil {
		return nil, err
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenDao) Set(ctx context.Context, tx redis.Pipeliner, m *userAccountToken.UserAccountToken) (*userAccountToken.UserAccountToken, error) {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := &UserAccountToken{
		UserId: m.UserId,
		Token:  m.Token,
	}

	jt, err := t.TableToJson()
	if err != nil {
		return nil, err
	}

	if err := conn.HSet(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), m.UserId), jt).Err(); err != nil {
		return nil, err
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenDao) Delete(ctx context.Context, tx redis.Pipeliner, m *userAccountToken.UserAccountToken) error {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := NewUserAccountToken()
	if err := conn.HDel(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), m.UserId)).Err(); err != nil {
		return err
	}

	return nil
}
