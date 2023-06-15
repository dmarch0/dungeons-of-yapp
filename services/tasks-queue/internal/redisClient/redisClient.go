package redisClient

import (
	"github.com/redis/go-redis/v9"
	"context"
	"fmt"
)

var conn *redis.Client

func ConnectToRedis() {
	ctx := context.Background()

	conn = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

	cmd, err := conn.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd)
}