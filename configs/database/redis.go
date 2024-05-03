package database

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisHandlerInstance *RedisHandler

type RedisHandler struct {
	Common *RedisConn
	User   *RedisConn
}

type RedisConn struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

// NewRedis インスタンスを作成する
func NewRedis() *RedisHandler {
	return RedisHandlerInstance
}

// InitRedis 初期化する
func InitRedis() (*RedisHandler, error) {
	db := &RedisHandler{}

	if err := db.commonDB(); err != nil {
		return nil, fmt.Errorf("db.commonDB: %s", err)
	}

	if err := db.userDB(); err != nil {
		return nil, fmt.Errorf("db.userDB: %s", err)
	}

	RedisHandlerInstance = db

	return RedisHandlerInstance, nil
}

// commonDB コネクションを作成する
func (s *RedisHandler) commonDB() error {
	host := os.Getenv("COMMON_REDIS_WRITE_HOST")
	password := os.Getenv("COMMON_REDIS_PASSWORD")
	database, err := strconv.Atoi(os.Getenv("COMMON_REDIS_DATABASE"))
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %s", err)
	}

	if err := s.setRedis(database, host, password); err != nil {
		return fmt.Errorf("s.setRedis: %s", err)
	}

	readDB := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       database,
	})

	writeDB := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       database,
	})

	s.Common = &RedisConn{
		ReadRedisConn:  readDB,
		WriteRedisConn: writeDB,
	}

	return nil
}

// userDB コネクションを作成する
func (s *RedisHandler) userDB() error {
	host := os.Getenv("USER_REDIS_WRITE_HOST")
	password := os.Getenv("USER_REDIS_PASSWORD")
	database, err := strconv.Atoi(os.Getenv("USER_REDIS_DATABASE"))
	if err != nil {
		return err
	}

	if err := s.setRedis(database, host, password); err != nil {
		return err
	}

	readDB := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       database,
	})

	writeDB := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       database,
	})

	s.User = &RedisConn{
		ReadRedisConn:  readDB,
		WriteRedisConn: writeDB,
	}

	return nil
}

// setRedis コネクションをセットする
func (s *RedisHandler) setRedis(database int, host, password string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       database,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return err
	}

	return nil
}
