package cache

import (
	"blog-gin/config"
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	Rab  *redis.Client
	Rctx context.Context
)

func init() {
	Rab = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword, // 没有密码，默认值
		DB:       config.RedisDb,       // 默认DB 0
	})

	Rctx = context.Background()
}

func ZScore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}
