package dao

import "github.com/go-redis/redis"

type RedisConf struct {
	Address string
}

func GetRedis(rc *RedisConf) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: rc.Address,
	})

	return rdb
}
