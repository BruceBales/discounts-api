package dao

import (
	"fmt"

	"github.com/brucebales/discounts-api/src/internal/config"
	"github.com/go-redis/redis"
)

type RedisConf struct {
	Address string
}

func NewRedis() (*redis.Client, error) {
	conf := config.GetConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPass, // no password set
		DB:       0,              // use default DB
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
