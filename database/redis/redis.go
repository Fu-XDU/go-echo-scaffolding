package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/gommon/log"
)

var rdb *redis.Client

func Connect(IP string, port int, username string, password string, db int) (Rdb *redis.Client, err error) {
	dbAddr := fmt.Sprintf("%s:%d", IP, port)
	rdb = redis.NewClient(&redis.Options{
		Username: username,
		Addr:     dbAddr,
		Password: password,
		DB:       db,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		log.Errorf("Connect to redis '%s'@'%s:%d', DB:%d failed", username, IP, port, db)
		return
	}
	log.Infof("Successfully connected to redis '%s'@'%s:%d', DB:%d", username, IP, port, db)
	Rdb = rdb
	return
}

func Subscribe(rdb *redis.Client, channels ...string) (ch <-chan *redis.Message, err error) {
	if len(channels) == 0 {
		return
	}
	pubsub := rdb.Subscribe(channels...)
	_, err = pubsub.Receive()
	if err != nil {
		log.Fatal(err)
	}
	ch = pubsub.Channel()
	return
}

func Get(rdb *redis.Client, key string) (val string, err error) {
	res := rdb.Get(key)
	val, err = res.Val(), res.Err()
	return
}
func Set(rdb *redis.Client, key string, value string) (val string, err error) {
	res := rdb.Set(key, value, 0)
	val, err = res.Val(), res.Err()
	return
}
