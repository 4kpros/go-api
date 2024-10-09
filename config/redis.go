package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisContext = context.Background()

func ConnectRedis() (err error) {
	addr := fmt.Sprintf("%s:%d", Env.RedisHost, Env.RedisPort)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: Env.RedisUserName,
		Password: Env.RedisPassword,
		DB:       Env.RedisDatabase,
	})

	return
}

func SetRedisVal(key string, val string) (err error) {
	err = RedisClient.Set(RedisContext, key, val, 0).Err()
	return
}

func GetRedisVal(key string) (val string, err error) {
	val, err = RedisClient.Get(RedisContext, key).Result()
	return
}

func DeleteRedisVal(key string) (val int64, err error) {
	val, err = RedisClient.Del(RedisContext, key).Result()
	return
}
