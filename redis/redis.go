package redis

import (
	"context"
	"fmt"

	"github.com/luckyss6/go-kit/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password, // no password set
		DB:       cfg.Redis.DB,       // use default DB
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(fmt.Errorf("connect redis err: %s", err.Error()))
	}
	return rdb
}
