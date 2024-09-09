package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectToRedis() (err error) {
	addr := fmt.Sprintf("%s:%d", AppEnv.RedisHost, AppEnv.RedisPort)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: AppEnv.RedisUserName,
		Password: AppEnv.RedisPassword,
		DB:       AppEnv.RedisDatabase,
	})

	return
}

func SetRedisVal(key string, val string) (err error) {
	err = RedisClient.Set(ctx, key, val, 0).Err()
	return
}

func GetRedisVal(key string) (val string, err error) {
	val, err = RedisClient.Get(ctx, key).Result()
	return
}

func DeleteRedisVal(key string) (val int64, err error) {
	val, err = RedisClient.Del(ctx, key).Result()
	return
}
