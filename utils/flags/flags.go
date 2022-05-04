package flags

import "gopkg.in/urfave/cli.v1"

var (
	PortFlag = cli.StringFlag{
		Name:   "port, p",
		Usage:  "Server port",
		Value:  "1423",
		EnvVar: "SERVER_PORT",
	}

	RedisHostFlag = cli.StringFlag{
		Name:        "RedisHost",
		Usage:       "RedisHost",
		Value:       "127.0.0.1",
		EnvVar:      "REDIS_HOST",
		Destination: &RedisHost,
	}
	RedisPortFlag = cli.IntFlag{
		Name:        "RedisPort",
		Usage:       "RedisPort",
		Value:       6379,
		EnvVar:      "REDIS_PORT",
		Destination: &RedisPort,
	}
	RedisDbFlag = cli.IntFlag{
		Name:        "RedisDb",
		Usage:       "RedisDb",
		Value:       0,
		EnvVar:      "REDIS_DB",
		Destination: &RedisDb,
	}
	RedisUserFlag = cli.StringFlag{
		Name:        "RedisUser",
		Usage:       "Database username for Redis",
		Value:       "",
		EnvVar:      "REDIS_USER",
		Destination: &RedisUser,
	}
	RedisPasswdFlag = cli.StringFlag{
		Name:        "RedisPasswd",
		Usage:       "Database password for Redis",
		Value:       "123456",
		EnvVar:      "REDIS_PASSWD",
		Destination: &RedisPasswd,
	}

	MysqlHostFlag = cli.StringFlag{
		Name:        "MysqlHost",
		Usage:       "MysqlHost",
		Value:       "127.0.0.1",
		EnvVar:      "MYSQL_HOST",
		Destination: &MysqlHost,
	}
	MysqlPortFlag = cli.IntFlag{
		Name:        "MysqlPort",
		Usage:       "MysqlPort",
		Value:       3306,
		EnvVar:      "MYSQL_PORT",
		Destination: &MysqlPort,
	}
	MysqlUserFlag = cli.StringFlag{
		Name:        "MysqlUser",
		Usage:       "MysqlUser",
		Value:       "root",
		EnvVar:      "MYSQL_USER",
		Destination: &MysqlUser,
	}
	MysqlPasswdFlag = cli.StringFlag{
		Name:        "MysqlPasswd",
		Usage:       "MysqlPasswd",
		Value:       "123456",
		EnvVar:      "MYSQL_PASSWD",
		Destination: &MysqlPasswd,
	}
	MysqlDBFlag = cli.StringFlag{
		Name:        "MysqlDB",
		Usage:       "mysqlDatabase",
		Value:       "emergency",
		EnvVar:      "MYSQL_DB",
		Destination: &MysqlDB,
	}
)
