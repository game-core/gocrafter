package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlHandlerInstance *MysqlHandler

type MysqlHandler struct {
	Common *MysqlConn
	Master *MysqlConn
	User   *ShardMysqlConn
}

type ShardMysqlConn struct {
	Shards map[string]*MysqlConn
}

type MysqlConn struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

// NewMysql インスタンスを作成する
func NewMysql() *MysqlHandler {
	return MysqlHandlerInstance
}

// InitMysql 初期化する
func InitMysql() (*MysqlHandler, error) {
	db := &MysqlHandler{}

	if err := db.commonDB(); err != nil {
		return nil, err
	}

	if err := db.masterDB(); err != nil {
		return nil, err
	}

	if err := db.shardUserDB(); err != nil {
		return nil, err
	}

	MysqlHandlerInstance = db
	return MysqlHandlerInstance, nil
}

// commonDB コネクションを作成する
func (s *MysqlHandler) commonDB() error {
	readMysqlConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("COMMON_MYSQL_READ_USER"),
		os.Getenv("COMMON_MYSQL_READ_PASSWORD"),
		os.Getenv("COMMON_MYSQL_READ_HOST"),
		os.Getenv("COMMON_MYSQL_DATABASE"),
	)

	writeMysqlConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("COMMON_MYSQL_WRITE_USER"),
		os.Getenv("COMMON_MYSQL_WRITE_PASSWORD"),
		os.Getenv("COMMON_MYSQL_WRITE_HOST"),
		os.Getenv("COMMON_MYSQL_DATABASE"),
	)

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readMysqlConn,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeMysqlConn,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	s.Common = &MysqlConn{
		ReadMysqlConn:  readDB,
		WriteMysqlConn: writeDB,
	}

	return nil
}

// masterDB コネクションを作成する
func (s *MysqlHandler) masterDB() error {
	readMysqlConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MASTER_MYSQL_READ_USER"),
		os.Getenv("MASTER_MYSQL_READ_PASSWORD"),
		os.Getenv("MASTER_MYSQL_READ_HOST"),
		os.Getenv("MASTER_MYSQL_DATABASE"),
	)

	writeMysqlConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MASTER_MYSQL_WRITE_USER"),
		os.Getenv("MASTER_MYSQL_WRITE_PASSWORD"),
		os.Getenv("MASTER_MYSQL_WRITE_HOST"),
		os.Getenv("MASTER_MYSQL_DATABASE"),
	)

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readMysqlConn,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeMysqlConn,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	s.Master = &MysqlConn{
		ReadMysqlConn:  readDB,
		WriteMysqlConn: writeDB,
	}

	return nil
}

// shardUserDB コネクションを作成する
func (s *MysqlHandler) shardUserDB() error {
	shardCountStr := os.Getenv("SHARD_COUNT")
	shardCount, err := strconv.Atoi(shardCountStr)
	if err != nil {
		return err
	}

	shards := make(map[string]*MysqlConn)
	for i := 0; i <= shardCount; i++ {
		userMysqlConn, err := s.userDB(fmt.Sprintf("_%v", i))
		if err != nil {
			return err
		}
		shards[os.Getenv(fmt.Sprintf("USER_MYSQL_SHARD_KEY_%v", i))] = userMysqlConn
	}

	s.User = &ShardMysqlConn{
		Shards: shards,
	}

	return nil
}

// userDB コネクションを作成する
func (s *MysqlHandler) userDB(shard string) (*MysqlConn, error) {
	readMysqlConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv(fmt.Sprintf("USER_MYSQL_READ_USER%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_READ_PASSWORD%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_READ_HOST%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_DATABASE%s", shard)),
	)

	writeMysqlConn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv(fmt.Sprintf("USER_MYSQL_WRITE_USER%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_WRITE_PASSWORD%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_WRITE_HOST%s", shard)),
		os.Getenv(fmt.Sprintf("USER_MYSQL_DATABASE%s", shard)),
	)

	readDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: readMysqlConn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	writeDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: writeMysqlConn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &MysqlConn{
		ReadMysqlConn:  readDB,
		WriteMysqlConn: writeDB,
	}, nil
}
