package dao

import (
	"github.com/go-redis/redis"
)

type RedisConf struct {
	Address string
}

func NewRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
