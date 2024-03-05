package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlHandlerInstance *SqlHandler

type SqlHandler struct {
	Common *Conn
	Master *Conn
	User   *ShardConn
}

type ShardConn struct {
	Shards map[string]*Conn
}

type Conn struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
}

func NewDB() *SqlHandler {
	return SqlHandlerInstance
}

func InitDB() (*SqlHandler, error) {
	common, err := commonDB()
	if err != nil {
		return nil, err
	}

	master, err := masterDB()
	if err != nil {
		return nil, err
	}

	user, err := shardUserDB()
	if err != nil {
		return nil, err
	}

	SqlHandlerInstance = &SqlHandler{
		Common: common,
		Master: master,
		User:   user,
	}

	return SqlHandlerInstance, err
}

func commonDB() (*Conn, error) {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("COMMON_MYSQL_READ_USER"),
		os.Getenv("COMMON_MYSQL_READ_PASSWORD"),
		os.Getenv("COMMON_MYSQL_READ_HOST"),
		os.Getenv("COMMON_MYSQL_DATABASE"),
	)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("COMMON_MYSQL_WRITE_USER"),
		os.Getenv("COMMON_MYSQL_WRITE_PASSWORD"),
		os.Getenv("COMMON_MYSQL_WRITE_HOST"),
		os.Getenv("COMMON_MYSQL_DATABASE"),
	)

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readConn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeConn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Conn{
		ReadConn:  readDB,
		WriteConn: writeDB,
	}, nil
}

func masterDB() (*Conn, error) {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MASTER_MYSQL_READ_USER"),
		os.Getenv("MASTER_MYSQL_READ_PASSWORD"),
		os.Getenv("MASTER_MYSQL_READ_HOST"),
		os.Getenv("MASTER_MYSQL_DATABASE"),
	)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MASTER_MYSQL_WRITE_USER"),
		os.Getenv("MASTER_MYSQL_WRITE_PASSWORD"),
		os.Getenv("MASTER_MYSQL_WRITE_HOST"),
		os.Getenv("MASTER_MYSQL_DATABASE"),
	)

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readConn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeConn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Conn{
		ReadConn:  readDB,
		WriteConn: writeDB,
	}, nil
}

func shardUserDB() (*ShardConn, error) {
	shardCountStr := os.Getenv("SHARD_COUNT")
	shardCount, err := strconv.Atoi(shardCountStr)
	if err != nil {
		return nil, err
	}

	shards := make(map[string]*Conn)
	for i := 0; i <= shardCount; i++ {
		userConn, err := userDB(fmt.Sprintf("_%v", i))
		if err != nil {
			return nil, err
		}
		shards[os.Getenv(fmt.Sprintf("USER_MYSQL_SHARD_KEY_%v", i))] = userConn
	}

	return &ShardConn{
		Shards: shards,
	}, nil
}

func userDB(shard string) (*Conn, error) {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv(fmt.Sprintf("USER_MYSQL_READ_USER%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_READ_PASSWORD%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_READ_HOST%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_DATABASE%s", shard)),
	)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv(fmt.Sprintf("USER_MYSQL_WRITE_USER%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_WRITE_PASSWORD%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_WRITE_HOST%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_DATABASE%s", shard)),
	)

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readConn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeConn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Conn{
		ReadConn:  readDB,
		WriteConn: writeDB,
	}, nil
}
