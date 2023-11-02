package utils

import (
	"Innovation/model"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type RedisHelper struct {
	*redis.Client
}

var redisHelper *RedisHelper

var redisOnce sync.Once

var ctx *context.Context

func GetRedisHelper() *RedisHelper {
	return redisHelper
}

func NewRedisHelper() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "101.35.238.12:6379",
		Password:     "chenyi888",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	},
	)

	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client = rdb
		redisHelper = rdh
	})

	return rdb
}

func InitRedis() {
	rdb := NewRedisHelper()
	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(err)
		return
	}
}

func Marshal(data interface{}) []byte {
	marshal, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return marshal
}

func UnMarshalToUser(data string) *model.User {
	var res *model.User
	json.Unmarshal([]byte(data), &res)
	return res
}
