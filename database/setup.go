package database

import (
	"server/database/mysql"
	"server/database/redis"
	"server/utils/flags"
)

func Setup() (err error) {
	_, err = mysql.Connect(flags.MysqlHost, flags.MysqlPort, flags.MysqlUser, flags.MysqlPasswd, flags.MysqlDB)
	if err != nil {
		return
	}
	_, err = redis.Connect(flags.RedisHost, flags.RedisPort, flags.RedisUser, flags.RedisPasswd, flags.RedisDb)
	return
}
