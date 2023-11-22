package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

type SqlHandler struct {
	Config *Conn
	Master *Conn
	User   *ShardConn
}

type ShardConn struct {
	Shards map[int]*Conn
}

type Conn struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
}

func NewDB() *SqlHandler {
	return &SqlHandler{
		Config: configDB(),
		Master: masterDB(),
		User:   shardUserDB(),
	}
}

func masterDB() *Conn {
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

	readDB, err := gorm.Open("mysql", readConn)
	if err != nil {
		panic(err.Error())
	}

	writeDB, err := gorm.Open("mysql", writeConn)
	if err != nil {
		panic(err.Error())
	}

	readDB.SingularTable(true)
	writeDB.SingularTable(true)

	return &Conn{
		ReadConn:  readDB,
		WriteConn: writeDB,
	}
}

func configDB() *Conn {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("CONFIG_MYSQL_READ_USER"),
		os.Getenv("CONFIG_MYSQL_READ_PASSWORD"),
		os.Getenv("CONFIG_MYSQL_READ_HOST"),
		os.Getenv("CONFIG_MYSQL_DATABASE"),
	)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("CONFIG_MYSQL_WRITE_USER"),
		os.Getenv("CONFIG_MYSQL_WRITE_PASSWORD"),
		os.Getenv("CONFIG_MYSQL_WRITE_HOST"),
		os.Getenv("CONFIG_MYSQL_DATABASE"),
	)

	readDB, err := gorm.Open("mysql", readConn)
	if err != nil {
		panic(err.Error())
	}

	writeDB, err := gorm.Open("mysql", writeConn)
	if err != nil {
		panic(err.Error())
	}

	readDB.SingularTable(true)
	writeDB.SingularTable(true)

	return &Conn{
		ReadConn:  readDB,
		WriteConn: writeDB,
	}
}

func shardUserDB() *ShardConn {
	shardCountStr := os.Getenv("SHARD_COUNT")
	shardCount, err := strconv.Atoi(shardCountStr)
	if err != nil {
		panic(err.Error())
	}

	shards := make(map[int]*Conn)
	for i := 1; i <= shardCount; i++ {
		shards[i] = userDB(fmt.Sprintf("_%v", i))
	}

	return &ShardConn{
		Shards: shards,
	}
}

func userDB(shard string) *Conn {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("USER_MYSQL_READ_USER"+shard),
		os.Getenv("USER_MYSQL_READ_PASSWORD"+shard),
		os.Getenv("USER_MYSQL_READ_HOST"+shard),
		os.Getenv("USER_MYSQL_DATABASE"+shard),
	)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("USER_MYSQL_WRITE_USER"+shard),
		os.Getenv("USER_MYSQL_WRITE_PASSWORD"+shard),
		os.Getenv("USER_MYSQL_WRITE_HOST"+shard),
		os.Getenv("USER_MYSQL_DATABASE"+shard),
	)

	readDB, err := gorm.Open("mysql", readConn)
	if err != nil {
		panic(err.Error())
	}

	writeDB, err := gorm.Open("mysql", writeConn)
	if err != nil {
		panic(err.Error())
	}

	readDB.SingularTable(true)
	writeDB.SingularTable(true)

	return &Conn{
		ReadConn:  readDB,
		WriteConn: writeDB,
	}
}
