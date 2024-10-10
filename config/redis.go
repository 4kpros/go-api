package config

import (
	"context"
	"fmt"
	"slices"

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

// String
func GetRedisStr(key string) (val string, err error) {
	val, err = RedisClient.Get(RedisContext, key).Result()
	return
}
func SetRedisStr(key string, val string) (err error) {
	err = RedisClient.Set(RedisContext, key, val, 0).Err()
	return
}
func DeleteRedisStr(key string) (val int64, err error) {
	val, err = RedisClient.Del(RedisContext, key).Result()
	return
}

// String array
func CheckRedisStrFromArrayStr(requiredToken string) func(string) (string, error) {
	return func(key string) (val string, err error) {
		var array, errArray = GetRedisArrayStr(key)
		if errArray != nil {
			err = errArray
			return
		}
		if slices.Contains(array, requiredToken) {
			val = requiredToken
		}
		return
	}
}
func GetRedisArrayStr(key string) (val []string, err error) {
	var len, errLen = RedisClient.LLen(RedisContext, key).Result()
	if errLen != nil {
		err = errLen
		return
	}
	val, err = RedisClient.LRange(RedisContext, key, 0, len-1).Result()
	return
}
func InsertRedisArrayStr(key string, val string) (err error) {
	err = RedisClient.LPush(RedisContext, key, val).Err()
	return
}
func DeleteRedisStrFromArrayStr(key string, index int64) (err error) {
	_, err = RedisClient.LTrim(RedisContext, key, index, index).Result()
	return
}
