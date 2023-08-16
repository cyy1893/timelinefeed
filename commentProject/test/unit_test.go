package test_test

import (
	"commentProject/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func createRedisClient(config *config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})
	return rdb
}

func TestRedisClient(t *testing.T) {
	// 加载Redis配置
	config := config.LoadRedisConfig()

	// 创建Redis客户端
	client := createRedisClient(config)

	value, err := client.Del(context.Background(), "key").Result()

	fmt.Println(value, err)
	fmt.Println()
}
