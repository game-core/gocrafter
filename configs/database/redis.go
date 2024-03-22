package database

var RedisHandlerInstance *RedisHandler

type RedisHandler struct {
	Common *MysqlConn
	Master *MysqlConn
	User   *ShardMysqlConn
}
