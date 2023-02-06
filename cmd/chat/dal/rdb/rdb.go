package rdb

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/wen-flower/easy-douyin/cmd/chat/cfg"
)

var (
	RDB  *redis.Client
	once sync.Once
)

// Init 初始化 Redis 连接
func Init() {
	once.Do(func() {
		RDB = redis.NewClient(&redis.Options{
			Addr: cfg.RedisAddr,
			DB:   0, // use default DB
		})

		_, err := RDB.Ping(context.Background()).Result()
		if err != nil {
			panic(err)
		}
	})
}
