package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Admin  *Conn
	Auth   *Conn
	Config *Conn
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
	return &SqlHandler{
		Admin:  adminDB(),
		Auth:   authDB(),
		Config: configDB(),
		Master: masterDB(),
		User:   shardUserDB(),
	}
}

func adminDB() *Conn {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("ADMIN_MYSQL_READ_USER"),
		os.Getenv("ADMIN_MYSQL_READ_PASSWORD"),
		os.Getenv("ADMIN_MYSQL_READ_HOST"),
		os.Getenv("ADMIN_MYSQL_DATABASE"),
	)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("ADMIN_MYSQL_WRITE_USER"),
		os.Getenv("ADMIN_MYSQL_WRITE_PASSWORD"),
		os.Getenv("ADMIN_MYSQL_WRITE_HOST"),
		os.Getenv("ADMIN_MYSQL_DATABASE"),
	)

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readConn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeConn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return &Conn{
		ReadConn:  readDB,
		WriteConn: writeDB,
	}
}

func authDB() *Conn {
	readConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("AUTH_MYSQL_READ_USER"),
		os.Getenv("AUTH_MYSQL_READ_PASSWORD"),
		os.Getenv("AUTH_MYSQL_READ_HOST"),
		os.Getenv("AUTH_MYSQL_DATABASE"),
	)

	writeConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("AUTH_MYSQL_WRITE_USER"),
		os.Getenv("AUTH_MYSQL_WRITE_PASSWORD"),
		os.Getenv("AUTH_MYSQL_WRITE_HOST"),
		os.Getenv("AUTH_MYSQL_DATABASE"),
	)

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readConn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeConn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

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

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readConn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeConn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return &Conn{
		ReadConn:  readDB,
		WriteConn: writeDB,
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

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readConn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeConn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

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

	shards := make(map[string]*Conn)
	for i := 1; i <= shardCount; i++ {
		shards[os.Getenv(fmt.Sprintf("USER_MYSQL_SHARD_KEY_%v", i))] = userDB(fmt.Sprintf("_%v", i))
	}

	return &ShardConn{
		Shards: shards,
	}
}

func userDB(shard string) *Conn {
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
		panic(err.Error())
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeConn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return &Conn{
		ReadConn:  readDB,
		WriteConn: writeDB,
	}
}
